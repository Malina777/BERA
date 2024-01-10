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

package historical

import (
	"context"

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/x/evm/plugins"
	"github.com/berachain/polaris/eth/core"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/params"
)

// Plugin is the interface that must be implemented by the plugin.
type Plugin interface {
	core.HistoricalPlugin
	plugins.HasGenesis
}

// plugin keeps track of polaris blocks via headers.
type plugin struct {
	// ctx is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// chainConfig stores the chain configuration for the evm chain.make
	chainConfig *params.ChainConfig
	// bp represents the block plugin, used for accessing historical block headers.
	bp core.BlockPlugin
	// storekey is the store key for the header store.
	storeKey storetypes.StoreKey
}

// NewPlugin creates a new instance of the block plugin from the given context.
func NewPlugin(
	chainConfig *params.ChainConfig, bp core.BlockPlugin,
	_ storetypes.StoreKey, storekey storetypes.StoreKey,
) Plugin {
	return &plugin{
		chainConfig: chainConfig,
		bp:          bp,
		storeKey:    storekey,
	}
}

// Prepare implements core.HistoricalPlugin.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}
