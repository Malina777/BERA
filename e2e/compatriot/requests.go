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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RPCRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int64  `json:"id"`
}

type ResponseErr struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int64       `json:"id"`
	Result  any         `json:"result"`
	Err     ResponseErr `json:"error"`
}

type RPCOutput struct {
	Request  RPCRequest  `json:"request"`
	Response RPCResponse `json:"response"`
}

// TODO: make this a buffer stream to not store all tests in memory?
var requests []RPCRequest

// Query loads prexisting JSON-RPC calls from a file and queries the chain
func query(outputFile string) error {
	calls := requests

	var output []RPCOutput
	for i := 0; i < len(calls); i++ {
		result, err := call(calls[i])
		if err != nil {
			return fmt.Errorf("Query: An error occurred %v when calling\n", err)
		}
		output = append(output, result)
	}

	// TODO: bring back when results make sense
	// if err := sanityCheck(output); err != nil {
	// 	return fmt.Errorf("Query: An error occurred %v when sanity checking results\n", err)
	// }

	// add the results to a file and format
	content, err := Marshal(output)
	if err != nil {
		return fmt.Errorf("Query: An error occurred %v when marshalling output\n", err)
	}

	if err = os.WriteFile("./"+outputFile, content, 0644); err != nil {
		return fmt.Errorf("call: An error occurred %v when writing output\n", err)
	}

	fmt.Println("finished querying")
	return nil
}

// call makes a JSON-RPC call to the chain and saves the results
func call(postRequest RPCRequest) (RPCOutput, error) {
	postBody, _ := json.Marshal(postRequest)
	buffer := bytes.NewBuffer(postBody)

	body, err := makeRequest(POLARIS_RPC, buffer)
	if err != nil {
		return RPCOutput{}, fmt.Errorf("call: An error occurred %v when making the request\n", err)
	}
	var response RPCResponse
	json.Unmarshal([]byte(body), &response)

	return RPCOutput{Request: postRequest, Response: response}, nil
}

// makeRequest makes the actual HTTP request to the chain
func makeRequest(rpc string, postBuffer *bytes.Buffer) (string, error) {
	response, err := http.Post(rpc, "application/json", postBuffer)
	if err != nil {
		return "", fmt.Errorf("makeRequest: An Error Occured %v when posting\n", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("makeRequest: An Error Occured %v when reading response\n", err)
	}
	return string(body), nil
}

// Marshal marshals the output slice to JSON
func Marshal(output []RPCOutput) ([]byte, error) {
	jsonOutput, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Marshal: An error occurred %v trying to marshal data\n", err)
	}
	return jsonOutput, nil
}