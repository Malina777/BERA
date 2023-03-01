// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethereumcoretypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"pkg.berachain.dev/stargazer/eth/core"
	ethcoretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"sync"
)

// Ensure, that BlockPluginMock does implement core.BlockPlugin.
// If this is not the case, regenerate this file with moq.
var _ core.BlockPlugin = &BlockPluginMock{}

// BlockPluginMock is a mock implementation of core.BlockPlugin.
//
//	func TestSomethingThatUsesBlockPlugin(t *testing.T) {
//
//		// make and configure a mocked core.BlockPlugin
//		mockedBlockPlugin := &BlockPluginMock{
//			BaseFeeFunc: func() uint64 {
//				panic("mock out the BaseFee method")
//			},
//			GetStargazerBlockByHashFunc: func(hash common.Hash) *ethcoretypes.StargazerBlock {
//				panic("mock out the GetStargazerBlockByHash method")
//			},
//			GetStargazerBlockByNumberFunc: func(n int64) *ethcoretypes.StargazerBlock {
//				panic("mock out the GetStargazerBlockByNumber method")
//			},
//			GetStargazerHeaderByNumberFunc: func(number int64) (*ethcoretypes.StargazerHeader, error) {
//				panic("mock out the GetStargazerHeaderByNumber method")
//			},
//			GetTransactionBlockNumberFunc: func(hash common.Hash) *big.Int {
//				panic("mock out the GetTransactionBlockNumber method")
//			},
//			GetTransactionByHashFunc: func(hash common.Hash) *ethereumcoretypes.Transaction {
//				panic("mock out the GetTransactionByHash method")
//			},
//			PrepareFunc: func(contextMoqParam context.Context)  {
//				panic("mock out the Prepare method")
//			},
//			PrepareHeaderFunc: func(ctx sdk.Context, header *ethcoretypes.StargazerHeader) *ethcoretypes.StargazerHeader {
//				panic("mock out the PrepareHeader method")
//			},
//			ProcessHeaderFunc: func(ctx sdk.Context, header *ethcoretypes.StargazerHeader) error {
//				panic("mock out the ProcessHeader method")
//			},
//		}
//
//		// use mockedBlockPlugin in code that requires core.BlockPlugin
//		// and then make assertions.
//
//	}
type BlockPluginMock struct {
	// BaseFeeFunc mocks the BaseFee method.
	BaseFeeFunc func() uint64

	// GetStargazerBlockByHashFunc mocks the GetStargazerBlockByHash method.
	GetStargazerBlockByHashFunc func(hash common.Hash) *ethcoretypes.StargazerBlock

	// GetStargazerBlockByNumberFunc mocks the GetStargazerBlockByNumber method.
	GetStargazerBlockByNumberFunc func(n int64) *ethcoretypes.StargazerBlock

	// GetStargazerHeaderByNumberFunc mocks the GetStargazerHeaderByNumber method.
	GetStargazerHeaderByNumberFunc func(number int64) (*ethcoretypes.StargazerHeader, error)

	// GetTransactionBlockNumberFunc mocks the GetTransactionBlockNumber method.
	GetTransactionBlockNumberFunc func(hash common.Hash) *big.Int

	// GetTransactionByHashFunc mocks the GetTransactionByHash method.
	GetTransactionByHashFunc func(hash common.Hash) *ethereumcoretypes.Transaction

	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(contextMoqParam context.Context)

	// PrepareHeaderFunc mocks the PrepareHeader method.
	PrepareHeaderFunc func(ctx sdk.Context, header *ethcoretypes.StargazerHeader) *ethcoretypes.StargazerHeader

	// ProcessHeaderFunc mocks the ProcessHeader method.
	ProcessHeaderFunc func(ctx sdk.Context, header *ethcoretypes.StargazerHeader) error

	// calls tracks calls to the methods.
	calls struct {
		// BaseFee holds details about calls to the BaseFee method.
		BaseFee []struct {
		}
		// GetStargazerBlockByHash holds details about calls to the GetStargazerBlockByHash method.
		GetStargazerBlockByHash []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// GetStargazerBlockByNumber holds details about calls to the GetStargazerBlockByNumber method.
		GetStargazerBlockByNumber []struct {
			// N is the n argument value.
			N int64
		}
		// GetStargazerHeaderByNumber holds details about calls to the GetStargazerHeaderByNumber method.
		GetStargazerHeaderByNumber []struct {
			// Number is the number argument value.
			Number int64
		}
		// GetTransactionBlockNumber holds details about calls to the GetTransactionBlockNumber method.
		GetTransactionBlockNumber []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// GetTransactionByHash holds details about calls to the GetTransactionByHash method.
		GetTransactionByHash []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// PrepareHeader holds details about calls to the PrepareHeader method.
		PrepareHeader []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Header is the header argument value.
			Header *ethcoretypes.StargazerHeader
		}
		// ProcessHeader holds details about calls to the ProcessHeader method.
		ProcessHeader []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Header is the header argument value.
			Header *ethcoretypes.StargazerHeader
		}
	}
	lockBaseFee                    sync.RWMutex
	lockGetStargazerBlockByHash    sync.RWMutex
	lockGetStargazerBlockByNumber  sync.RWMutex
	lockGetStargazerHeaderByNumber sync.RWMutex
	lockGetTransactionBlockNumber  sync.RWMutex
	lockGetTransactionByHash       sync.RWMutex
	lockPrepare                    sync.RWMutex
	lockPrepareHeader              sync.RWMutex
	lockProcessHeader              sync.RWMutex
}

// BaseFee calls BaseFeeFunc.
func (mock *BlockPluginMock) BaseFee() uint64 {
	if mock.BaseFeeFunc == nil {
		panic("BlockPluginMock.BaseFeeFunc: method is nil but BlockPlugin.BaseFee was just called")
	}
	callInfo := struct {
	}{}
	mock.lockBaseFee.Lock()
	mock.calls.BaseFee = append(mock.calls.BaseFee, callInfo)
	mock.lockBaseFee.Unlock()
	return mock.BaseFeeFunc()
}

// BaseFeeCalls gets all the calls that were made to BaseFee.
// Check the length with:
//
//	len(mockedBlockPlugin.BaseFeeCalls())
func (mock *BlockPluginMock) BaseFeeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockBaseFee.RLock()
	calls = mock.calls.BaseFee
	mock.lockBaseFee.RUnlock()
	return calls
}

// GetStargazerBlockByHash calls GetStargazerBlockByHashFunc.
func (mock *BlockPluginMock) GetStargazerBlockByHash(hash common.Hash) *ethcoretypes.StargazerBlock {
	if mock.GetStargazerBlockByHashFunc == nil {
		panic("BlockPluginMock.GetStargazerBlockByHashFunc: method is nil but BlockPlugin.GetStargazerBlockByHash was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetStargazerBlockByHash.Lock()
	mock.calls.GetStargazerBlockByHash = append(mock.calls.GetStargazerBlockByHash, callInfo)
	mock.lockGetStargazerBlockByHash.Unlock()
	return mock.GetStargazerBlockByHashFunc(hash)
}

// GetStargazerBlockByHashCalls gets all the calls that were made to GetStargazerBlockByHash.
// Check the length with:
//
//	len(mockedBlockPlugin.GetStargazerBlockByHashCalls())
func (mock *BlockPluginMock) GetStargazerBlockByHashCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetStargazerBlockByHash.RLock()
	calls = mock.calls.GetStargazerBlockByHash
	mock.lockGetStargazerBlockByHash.RUnlock()
	return calls
}

// GetStargazerBlockByNumber calls GetStargazerBlockByNumberFunc.
func (mock *BlockPluginMock) GetStargazerBlockByNumber(n int64) *ethcoretypes.StargazerBlock {
	if mock.GetStargazerBlockByNumberFunc == nil {
		panic("BlockPluginMock.GetStargazerBlockByNumberFunc: method is nil but BlockPlugin.GetStargazerBlockByNumber was just called")
	}
	callInfo := struct {
		N int64
	}{
		N: n,
	}
	mock.lockGetStargazerBlockByNumber.Lock()
	mock.calls.GetStargazerBlockByNumber = append(mock.calls.GetStargazerBlockByNumber, callInfo)
	mock.lockGetStargazerBlockByNumber.Unlock()
	return mock.GetStargazerBlockByNumberFunc(n)
}

// GetStargazerBlockByNumberCalls gets all the calls that were made to GetStargazerBlockByNumber.
// Check the length with:
//
//	len(mockedBlockPlugin.GetStargazerBlockByNumberCalls())
func (mock *BlockPluginMock) GetStargazerBlockByNumberCalls() []struct {
	N int64
} {
	var calls []struct {
		N int64
	}
	mock.lockGetStargazerBlockByNumber.RLock()
	calls = mock.calls.GetStargazerBlockByNumber
	mock.lockGetStargazerBlockByNumber.RUnlock()
	return calls
}

// GetStargazerHeaderByNumber calls GetStargazerHeaderByNumberFunc.
func (mock *BlockPluginMock) GetStargazerHeaderByNumber(number int64) (*ethcoretypes.StargazerHeader, error) {
	if mock.GetStargazerHeaderByNumberFunc == nil {
		panic("BlockPluginMock.GetStargazerHeaderByNumberFunc: method is nil but BlockPlugin.GetStargazerHeaderByNumber was just called")
	}
	callInfo := struct {
		Number int64
	}{
		Number: number,
	}
	mock.lockGetStargazerHeaderByNumber.Lock()
	mock.calls.GetStargazerHeaderByNumber = append(mock.calls.GetStargazerHeaderByNumber, callInfo)
	mock.lockGetStargazerHeaderByNumber.Unlock()
	return mock.GetStargazerHeaderByNumberFunc(number)
}

// GetStargazerHeaderByNumberCalls gets all the calls that were made to GetStargazerHeaderByNumber.
// Check the length with:
//
//	len(mockedBlockPlugin.GetStargazerHeaderByNumberCalls())
func (mock *BlockPluginMock) GetStargazerHeaderByNumberCalls() []struct {
	Number int64
} {
	var calls []struct {
		Number int64
	}
	mock.lockGetStargazerHeaderByNumber.RLock()
	calls = mock.calls.GetStargazerHeaderByNumber
	mock.lockGetStargazerHeaderByNumber.RUnlock()
	return calls
}

// GetTransactionBlockNumber calls GetTransactionBlockNumberFunc.
func (mock *BlockPluginMock) GetTransactionBlockNumber(hash common.Hash) *big.Int {
	if mock.GetTransactionBlockNumberFunc == nil {
		panic("BlockPluginMock.GetTransactionBlockNumberFunc: method is nil but BlockPlugin.GetTransactionBlockNumber was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetTransactionBlockNumber.Lock()
	mock.calls.GetTransactionBlockNumber = append(mock.calls.GetTransactionBlockNumber, callInfo)
	mock.lockGetTransactionBlockNumber.Unlock()
	return mock.GetTransactionBlockNumberFunc(hash)
}

// GetTransactionBlockNumberCalls gets all the calls that were made to GetTransactionBlockNumber.
// Check the length with:
//
//	len(mockedBlockPlugin.GetTransactionBlockNumberCalls())
func (mock *BlockPluginMock) GetTransactionBlockNumberCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetTransactionBlockNumber.RLock()
	calls = mock.calls.GetTransactionBlockNumber
	mock.lockGetTransactionBlockNumber.RUnlock()
	return calls
}

// GetTransactionByHash calls GetTransactionByHashFunc.
func (mock *BlockPluginMock) GetTransactionByHash(hash common.Hash) *ethereumcoretypes.Transaction {
	if mock.GetTransactionByHashFunc == nil {
		panic("BlockPluginMock.GetTransactionByHashFunc: method is nil but BlockPlugin.GetTransactionByHash was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetTransactionByHash.Lock()
	mock.calls.GetTransactionByHash = append(mock.calls.GetTransactionByHash, callInfo)
	mock.lockGetTransactionByHash.Unlock()
	return mock.GetTransactionByHashFunc(hash)
}

// GetTransactionByHashCalls gets all the calls that were made to GetTransactionByHash.
// Check the length with:
//
//	len(mockedBlockPlugin.GetTransactionByHashCalls())
func (mock *BlockPluginMock) GetTransactionByHashCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetTransactionByHash.RLock()
	calls = mock.calls.GetTransactionByHash
	mock.lockGetTransactionByHash.RUnlock()
	return calls
}

// Prepare calls PrepareFunc.
func (mock *BlockPluginMock) Prepare(contextMoqParam context.Context) {
	if mock.PrepareFunc == nil {
		panic("BlockPluginMock.PrepareFunc: method is nil but BlockPlugin.Prepare was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	mock.lockPrepare.Unlock()
	mock.PrepareFunc(contextMoqParam)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//
//	len(mockedBlockPlugin.PrepareCalls())
func (mock *BlockPluginMock) PrepareCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockPrepare.RLock()
	calls = mock.calls.Prepare
	mock.lockPrepare.RUnlock()
	return calls
}

// PrepareHeader calls PrepareHeaderFunc.
func (mock *BlockPluginMock) PrepareHeader(ctx sdk.Context, header *ethcoretypes.StargazerHeader) *ethcoretypes.StargazerHeader {
	if mock.PrepareHeaderFunc == nil {
		panic("BlockPluginMock.PrepareHeaderFunc: method is nil but BlockPlugin.PrepareHeader was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Header *ethcoretypes.StargazerHeader
	}{
		Ctx:    ctx,
		Header: header,
	}
	mock.lockPrepareHeader.Lock()
	mock.calls.PrepareHeader = append(mock.calls.PrepareHeader, callInfo)
	mock.lockPrepareHeader.Unlock()
	return mock.PrepareHeaderFunc(ctx, header)
}

// PrepareHeaderCalls gets all the calls that were made to PrepareHeader.
// Check the length with:
//
//	len(mockedBlockPlugin.PrepareHeaderCalls())
func (mock *BlockPluginMock) PrepareHeaderCalls() []struct {
	Ctx    sdk.Context
	Header *ethcoretypes.StargazerHeader
} {
	var calls []struct {
		Ctx    sdk.Context
		Header *ethcoretypes.StargazerHeader
	}
	mock.lockPrepareHeader.RLock()
	calls = mock.calls.PrepareHeader
	mock.lockPrepareHeader.RUnlock()
	return calls
}

// ProcessHeader calls ProcessHeaderFunc.
func (mock *BlockPluginMock) ProcessHeader(ctx sdk.Context, header *ethcoretypes.StargazerHeader) error {
	if mock.ProcessHeaderFunc == nil {
		panic("BlockPluginMock.ProcessHeaderFunc: method is nil but BlockPlugin.ProcessHeader was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Header *ethcoretypes.StargazerHeader
	}{
		Ctx:    ctx,
		Header: header,
	}
	mock.lockProcessHeader.Lock()
	mock.calls.ProcessHeader = append(mock.calls.ProcessHeader, callInfo)
	mock.lockProcessHeader.Unlock()
	return mock.ProcessHeaderFunc(ctx, header)
}

// ProcessHeaderCalls gets all the calls that were made to ProcessHeader.
// Check the length with:
//
//	len(mockedBlockPlugin.ProcessHeaderCalls())
func (mock *BlockPluginMock) ProcessHeaderCalls() []struct {
	Ctx    sdk.Context
	Header *ethcoretypes.StargazerHeader
} {
	var calls []struct {
		Ctx    sdk.Context
		Header *ethcoretypes.StargazerHeader
	}
	mock.lockProcessHeader.RLock()
	calls = mock.calls.ProcessHeader
	mock.lockProcessHeader.RUnlock()
	return calls
}
