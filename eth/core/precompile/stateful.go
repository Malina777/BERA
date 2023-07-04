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

package precompile

import (
	"context"
	"math/big"
	"reflect"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"
)

// NumBytesMethodID is the number of bytes used to represent a ABI method's ID.
const NumBytesMethodID = 4

// stateful is a container for running stateful and dynamic precompiled contracts.
type stateful struct {
	// Registrable is the base precompile implementation.
	Registrable
	// idsToMethods is a mapping of method IDs (string of first 4 bytes of the keccak256 hash of
	// method signatures) to native precompile functions. The signature key is provided by the
	// precompile creator and must exactly match the signature in the geth abi.Method.Sig field
	// (geth abi format). Please check core/precompile/container/method.go for more information.
	idsToMethods map[string]*Method
	// receive      *Method // TODO: implement
	// fallback     *Method // TODO: implement

}

// NewStateful creates and returns a new `stateful` with the given method ids precompile functions map.
func NewStateful(
	rp Registrable, idsToMethods map[string]*Method,
) vm.PrecompileContainer {
	return &stateful{
		Registrable:  rp,
		idsToMethods: idsToMethods,
	}
}

// Run loads the corresponding precompile method for given input, executes it, and handles
// output.
//
// Run implements `PrecompileContainer`.
func (sc *stateful) Run(
	ctx context.Context,
	evm EVM,
	input []byte,
	caller common.Address,
	value *big.Int,
	readonly bool,
) ([]byte, error) {
	if sc.idsToMethods == nil {
		return nil, ErrContainerHasNoMethods
	}
	if len(input) < NumBytesMethodID {
		return nil, ErrInvalidInputToPrecompile
	}

	// Extract the method ID from the input and load the method.
	method, found := sc.idsToMethods[utils.UnsafeBytesToStr(input[:NumBytesMethodID])]
	if !found {
		return nil, ErrMethodNotFound
	}

	// Get args ready for precompile call.
	// TODO, remove most of these args.
	// In the future , we should only need the arguments from the method according to the ABI
	// and a context rather than all of these.
	return method.Call(
		[]reflect.Value{
			reflect.ValueOf(sc.Registrable),
			reflect.ValueOf(ctx),
			reflect.ValueOf(evm),
			reflect.ValueOf(caller),
			reflect.ValueOf(value),
			reflect.ValueOf(readonly),
		}, input)
}

// RequiredGas checks the Method corresponding to input for the required gas amount.
//
// RequiredGas implements PrecompileContainer.
func (sc *stateful) RequiredGas(input []byte) uint64 {
	if sc.idsToMethods == nil || len(input) < NumBytesMethodID {
		return 0
	}

	// Extract the method ID from the input and load the method.
	method, found := sc.idsToMethods[utils.UnsafeBytesToStr(input[:NumBytesMethodID])]
	if !found {
		return 0
	}

	return method.RequiredGas
}
