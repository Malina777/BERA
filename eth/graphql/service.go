// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package graphql

//
// import (
//
//	"context"
//	"encoding/json"
//	"net/http"
//	"strconv"
//	"sync"
//	"time"
//
//	"github.com/ethereum/go-ethereum/eth/filters"
//	"github.com/ethereum/go-ethereum/ethapi"
//	"github.com/ethereum/go-ethereum/node"
//	"github.com/ethereum/go-ethereum/rpc"
//	"github.com/graph-gophers/graphql-go"
//	gqlErrors "github.com/graph-gophers/graphql-go/errors"
//
// )
//
//	type handler struct {
//		Schema *graphql.Schema
//	}
//
//	func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		var params struct {
//			Query         string                 `json:"query"`
//			OperationName string                 `json:"operationName"`
//			Variables     map[string]interface{} `json:"variables"`
//		}
//		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		var (
//			ctx       = r.Context()
//			responded sync.Once
//			timer     *time.Timer
//			cancel    context.CancelFunc
//		)
//		ctx, cancel = context.WithCancel(ctx)
//		defer cancel()
//
//		if timeout, ok := rpc.ContextRequestTimeout(ctx); ok {
//			timer = time.AfterFunc(timeout, func() {
//				responded.Do(func() {
//					// Cancel request handling.
//					cancel()
//
//					// Create the timeout response.
//					response := &graphql.Response{
//						Errors: []*gqlErrors.QueryError{{Message: "request timed out"}},
//					}
//					responseJSON, err := json.Marshal(response)
//					if err != nil {
//						http.Error(w, err.Error(), http.StatusInternalServerError)
//						return
//					}
//
//					// Setting this disables gzip compression in package node.
//					w.Header().Set("transfer-encoding", "identity")
//
//					// Flush the response. Since we are writing close to the response timeout,
//					// chunked transfer encoding must be disabled by setting content-length.
//					w.Header().Set("content-type", "application/json")
//					w.Header().Set("content-length", strconv.Itoa(len(responseJSON)))
//					w.Write(responseJSON)
//					if flush, ok := w.(http.Flusher); ok {
//						flush.Flush()
//					}
//				})
//			})
//		}
//
//		response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
//		timer.Stop()
//		responded.Do(func() {
//			responseJSON, err := json.Marshal(response)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//			if len(response.Errors) > 0 {
//				w.WriteHeader(http.StatusBadRequest)
//			}
//			w.Header().Set("Content-Type", "application/json")
//			w.Write(responseJSON)
//		})
//	}
//
// // New constructs a new GraphQL service instance.
//
//	func New(stack *node.Node, backend ethapi.Backend, filterSystem *filters.FilterSystem, cors, vhosts []string) error {
//		_, err := newHandler(stack, backend, filterSystem, cors, vhosts)
//		return err
//	}
//
// // newHandler returns a new `http.Handler` that will answer GraphQL queries.
// // It additionally exports an interactive query browser on the / endpoint.
//
//	func newHandler(stack *node.Node, backend ethapi.Backend, filterSystem *filters.FilterSystem, cors, vhosts []string) (*handler, error) {
//		q := Resolver{backend, filterSystem}
//
//		s, err := graphql.ParseSchema(schema, &q)
//		if err != nil {
//			return nil, err
//		}
//		h := handler{Schema: s}
//		handler := node.NewHTTPHandlerStack(h, cors, vhosts, nil)
//
//		stack.RegisterHandler("GraphQL UI", "/graphql/ui", GraphiQL{})
//		stack.RegisterHandler("GraphQL", "/graphql", handler)
//		stack.RegisterHandler("GraphQL", "/graphql/", handler)
//
//		return &h, nil
//	}
