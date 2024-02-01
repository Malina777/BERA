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

package txpool

import (
	"errors"
	"time"

	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/lib/utils"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// AnteHandle implements sdk.AnteHandler.
// It is used to determine whether transactions should be ejected
// from the comet mempool.
func (m *Mempool) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	// The transaction put into this function by CheckTx
	// is a transaction from CometBFT mempool.
	telemetry.IncrCounter(float32(1), MetricKeyCometPoolTxs)
	msgs := tx.GetMsgs()

	ctx.Logger().Info("AnteHandle Polaris Mempool", "msgs", len(msgs), "simulate", simulate)

	// TODO: Record the time it takes to build a payload.

	// We only want to eject transactions from comet on recheck.
	if ctx.ExecMode() == sdk.ExecModeCheck || ctx.ExecMode() == sdk.ExecModeReCheck {
		if wet, ok := utils.GetAs[*types.WrappedEthereumTransaction](msgs[0]); ok {
			ethTx := wet.Unwrap()
			ctx.Logger().Info(
				"AnteHandle for eth tx", "tx", ethTx.Hash(), "mode", ctx.ExecMode(),
				"isRemote", m.crc.IsRemoteTx(ethTx.Hash()),
			)
			if shouldEject := m.shouldEjectFromCometMempool(
				ctx, ethTx,
			); shouldEject {
				ctx.Logger().Info("AnteHandle dropping tx from comet mempool", "tx", ethTx.Hash())
				telemetry.IncrCounter(float32(1), MetricKeyAnteEjectedTxs)
				return ctx, errors.New("eject from comet mempool")
			} else if ctx.ExecMode() == sdk.ExecModeReCheck && m.forceBroadcastOnRecheck {
				ctx.Logger().Info("AnteHandle forwarding to validator", "tx", ethTx.Hash())
				// We optionally force a re-broadcast.
				m.ForwardToValidator(ethTx)
			}
			ctx.Logger().Info("AnteHandle NOT dropping comet mempool", "tx", ethTx.Hash())
		}
	}
	return next(ctx, tx, simulate)
}

// shouldEject returns true if the transaction should be ejected from the CometBFT mempool.
func (m *Mempool) shouldEjectFromCometMempool(
	ctx sdk.Context, tx *ethtypes.Transaction,
) bool {
	defer telemetry.MeasureSince(time.Now(), MetricKeyTimeShouldEject)
	if tx == nil {
		ctx.Logger().Info("shouldEjectFromCometMempool: tx is nil")
		return false
	}

	// First check things that are stateless.
	if m.validateStateless(ctx, tx) {
		ctx.Logger().Info("shouldEjectFromCometMempool: stateless failed", "tx", tx.Hash())
		return true
	}

	// Then check for things that are stateful.
	return m.validateStateful(ctx, tx)
}

// validateStateless returns whether the tx of the given hash is stateless.
func (m *Mempool) validateStateless(ctx sdk.Context, tx *ethtypes.Transaction) bool {
	txHash := tx.Hash()
	currentTime := ctx.BlockTime()

	// 1. If the transaction has been in the mempool for longer than the configured timeout.
	expired := currentTime.Sub(m.crc.TimeFirstSeen(txHash)) > m.lifetime
	if expired {
		telemetry.IncrCounter(float32(1), MetricKeyAnteShouldEjectExpiredTx)
	}

	// 2. If the transaction's gas params are less than or equal to the configured limit.
	priceLeLimit := tx.GasPrice().Cmp(m.priceLimit) <= 0
	if priceLeLimit {
		telemetry.IncrCounter(float32(1), MetricKeyAnteShouldEjectPriceLimit)
	}

	ctx.Logger().Info(
		"validateStateless",
		"timeFirstSeen", m.crc.TimeFirstSeen(txHash),
		"timeInMempool", currentTime.Sub(m.crc.TimeFirstSeen(txHash)).Seconds(),
		"txHash", txHash, "currentTime", currentTime,
		"expired", expired, "priceLeLimit", priceLeLimit,
	)

	return expired || priceLeLimit
}

// includedCanonicalChain returns whether the tx of the given hash is included in the canonical
// Eth chain.
func (m *Mempool) validateStateful(ctx sdk.Context, tx *ethtypes.Transaction) bool {
	// // 1. If the transaction has been included in a block.
	// signer := ethtypes.LatestSignerForChainID(m.chainConfig.ChainID)
	// if _, err := ethtypes.Sender(signer, tx); err != nil {
	// 	return true
	// }

	included := m.chain.GetTransactionLookup(tx.Hash()) != nil
	telemetry.IncrCounter(float32(1), MetricKeyAnteShouldEjectInclusion)
	ctx.Logger().Info("validateStateful", "included", included, "txHash", tx.Hash())
	return included
}
