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

package keeper

import (
	"fmt"
	"sync"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/erc20/store"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
)

// Keeper of this module maintains collections of erc20.
type Keeper struct {
	storeKey   storetypes.StoreKey
	bankKeeper BankKeeper
	authority  sdk.AccAddress

	deployLock sync.Mutex
}

// NewKeeper creates new instances of the erc20 Keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
	bk BankKeeper,
	authority sdk.AccAddress,
) *Keeper {
	return &Keeper{
		storeKey:   storeKey,
		bankKeeper: bk,
		authority:  authority,
	}
}

// DenomKVStore returns a KVStore for the given denom.
func (k *Keeper) DenomKVStore(ctx sdk.Context) store.DenomKVStore {
	return store.NewDenomKVStore(ctx.KVStore(k.storeKey))
}

// RegisterDenomTokenPair registers a new token pair.
func (k *Keeper) RegisterDenomTokenPair(ctx sdk.Context, denom string, token common.Address) {
	k.DenomKVStore(ctx).SetAddressForDenom(denom, token)
	k.DenomKVStore(ctx).SetDenomForAddress(token, denom)
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
