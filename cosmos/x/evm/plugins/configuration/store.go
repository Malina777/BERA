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

package configuration

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethparams "pkg.berachain.dev/polaris/eth/params"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	enclib "pkg.berachain.dev/polaris/lib/encoding"
)

// GetParams is used to get the params for the evm module.
func (p *plugin) GetParams(ctx sdk.Context) *types.Params {
	bz := ctx.KVStore(p.storeKey).Get([]byte{types.ParamsKey})
	if bz == nil {
		return &types.Params{}
	}

	var params types.Params
	if err := params.Unmarshal(bz); err != nil {
		panic(err)
	}

	// update cached values
	if p.evmDenom != params.EvmDenom {
		p.evmDenom = params.EvmDenom
	}
	if newChainConfig := enclib.MustUnmarshalJSON[ethparams.ChainConfig]([]byte(params.ChainConfig)); p.chainConfig != newChainConfig { //nolint:lll // ok.
		p.chainConfig = newChainConfig
	}

	return &params
}

// SetParams is used to set the params for the evm module.
func (p *plugin) SetParams(ctx sdk.Context, params *types.Params) {
	bz, err := params.Marshal()
	if err != nil {
		panic(err)
	}
	ctx.KVStore(p.storeKey).Set([]byte{types.ParamsKey}, bz)

	// update cached values
	p.evmDenom = params.EvmDenom
	p.chainConfig = enclib.MustUnmarshalJSON[ethparams.ChainConfig]([]byte(params.ChainConfig))
}
