// Code generated by mockery v2.34.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	journal "pkg.berachain.dev/polaris/eth/core/state/journal"

	mock "github.com/stretchr/testify/mock"
)

// TransientStorage is an autogenerated mock type for the TransientStorage type
type TransientStorage struct {
	mock.Mock
}

type TransientStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *TransientStorage) EXPECT() *TransientStorage_Expecter {
	return &TransientStorage_Expecter{mock: &_m.Mock}
}

// Clone provides a mock function with given fields:
func (_m *TransientStorage) Clone() journal.TransientStorage {
	ret := _m.Called()

	var r0 journal.TransientStorage
	if rf, ok := ret.Get(0).(func() journal.TransientStorage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(journal.TransientStorage)
		}
	}

	return r0
}

// TransientStorage_Clone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clone'
type TransientStorage_Clone_Call struct {
	*mock.Call
}

// Clone is a helper method to define mock.On call
func (_e *TransientStorage_Expecter) Clone() *TransientStorage_Clone_Call {
	return &TransientStorage_Clone_Call{Call: _e.mock.On("Clone")}
}

func (_c *TransientStorage_Clone_Call) Run(run func()) *TransientStorage_Clone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TransientStorage_Clone_Call) Return(_a0 journal.TransientStorage) *TransientStorage_Clone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransientStorage_Clone_Call) RunAndReturn(run func() journal.TransientStorage) *TransientStorage_Clone_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields:
func (_m *TransientStorage) Finalize() {
	_m.Called()
}

// TransientStorage_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type TransientStorage_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
func (_e *TransientStorage_Expecter) Finalize() *TransientStorage_Finalize_Call {
	return &TransientStorage_Finalize_Call{Call: _e.mock.On("Finalize")}
}

func (_c *TransientStorage_Finalize_Call) Run(run func()) *TransientStorage_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TransientStorage_Finalize_Call) Return() *TransientStorage_Finalize_Call {
	_c.Call.Return()
	return _c
}

func (_c *TransientStorage_Finalize_Call) RunAndReturn(run func()) *TransientStorage_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransientState provides a mock function with given fields: addr, key
func (_m *TransientStorage) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	ret := _m.Called(addr, key)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) common.Hash); ok {
		r0 = rf(addr, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// TransientStorage_GetTransientState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransientState'
type TransientStorage_GetTransientState_Call struct {
	*mock.Call
}

// GetTransientState is a helper method to define mock.On call
//   - addr common.Address
//   - key common.Hash
func (_e *TransientStorage_Expecter) GetTransientState(addr interface{}, key interface{}) *TransientStorage_GetTransientState_Call {
	return &TransientStorage_GetTransientState_Call{Call: _e.mock.On("GetTransientState", addr, key)}
}

func (_c *TransientStorage_GetTransientState_Call) Run(run func(addr common.Address, key common.Hash)) *TransientStorage_GetTransientState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *TransientStorage_GetTransientState_Call) Return(_a0 common.Hash) *TransientStorage_GetTransientState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransientStorage_GetTransientState_Call) RunAndReturn(run func(common.Address, common.Hash) common.Hash) *TransientStorage_GetTransientState_Call {
	_c.Call.Return(run)
	return _c
}

// RegistryKey provides a mock function with given fields:
func (_m *TransientStorage) RegistryKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TransientStorage_RegistryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegistryKey'
type TransientStorage_RegistryKey_Call struct {
	*mock.Call
}

// RegistryKey is a helper method to define mock.On call
func (_e *TransientStorage_Expecter) RegistryKey() *TransientStorage_RegistryKey_Call {
	return &TransientStorage_RegistryKey_Call{Call: _e.mock.On("RegistryKey")}
}

func (_c *TransientStorage_RegistryKey_Call) Run(run func()) *TransientStorage_RegistryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TransientStorage_RegistryKey_Call) Return(_a0 string) *TransientStorage_RegistryKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransientStorage_RegistryKey_Call) RunAndReturn(run func() string) *TransientStorage_RegistryKey_Call {
	_c.Call.Return(run)
	return _c
}

// RevertToSnapshot provides a mock function with given fields: _a0
func (_m *TransientStorage) RevertToSnapshot(_a0 int) {
	_m.Called(_a0)
}

// TransientStorage_RevertToSnapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertToSnapshot'
type TransientStorage_RevertToSnapshot_Call struct {
	*mock.Call
}

// RevertToSnapshot is a helper method to define mock.On call
//   - _a0 int
func (_e *TransientStorage_Expecter) RevertToSnapshot(_a0 interface{}) *TransientStorage_RevertToSnapshot_Call {
	return &TransientStorage_RevertToSnapshot_Call{Call: _e.mock.On("RevertToSnapshot", _a0)}
}

func (_c *TransientStorage_RevertToSnapshot_Call) Run(run func(_a0 int)) *TransientStorage_RevertToSnapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *TransientStorage_RevertToSnapshot_Call) Return() *TransientStorage_RevertToSnapshot_Call {
	_c.Call.Return()
	return _c
}

func (_c *TransientStorage_RevertToSnapshot_Call) RunAndReturn(run func(int)) *TransientStorage_RevertToSnapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SetTransientState provides a mock function with given fields: addr, key, value
func (_m *TransientStorage) SetTransientState(addr common.Address, key common.Hash, value common.Hash) {
	_m.Called(addr, key, value)
}

// TransientStorage_SetTransientState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTransientState'
type TransientStorage_SetTransientState_Call struct {
	*mock.Call
}

// SetTransientState is a helper method to define mock.On call
//   - addr common.Address
//   - key common.Hash
//   - value common.Hash
func (_e *TransientStorage_Expecter) SetTransientState(addr interface{}, key interface{}, value interface{}) *TransientStorage_SetTransientState_Call {
	return &TransientStorage_SetTransientState_Call{Call: _e.mock.On("SetTransientState", addr, key, value)}
}

func (_c *TransientStorage_SetTransientState_Call) Run(run func(addr common.Address, key common.Hash, value common.Hash)) *TransientStorage_SetTransientState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash), args[2].(common.Hash))
	})
	return _c
}

func (_c *TransientStorage_SetTransientState_Call) Return() *TransientStorage_SetTransientState_Call {
	_c.Call.Return()
	return _c
}

func (_c *TransientStorage_SetTransientState_Call) RunAndReturn(run func(common.Address, common.Hash, common.Hash)) *TransientStorage_SetTransientState_Call {
	_c.Call.Return(run)
	return _c
}

// Snapshot provides a mock function with given fields:
func (_m *TransientStorage) Snapshot() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// TransientStorage_Snapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Snapshot'
type TransientStorage_Snapshot_Call struct {
	*mock.Call
}

// Snapshot is a helper method to define mock.On call
func (_e *TransientStorage_Expecter) Snapshot() *TransientStorage_Snapshot_Call {
	return &TransientStorage_Snapshot_Call{Call: _e.mock.On("Snapshot")}
}

func (_c *TransientStorage_Snapshot_Call) Run(run func()) *TransientStorage_Snapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TransientStorage_Snapshot_Call) Return(_a0 int) *TransientStorage_Snapshot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransientStorage_Snapshot_Call) RunAndReturn(run func() int) *TransientStorage_Snapshot_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransientStorage creates a new instance of TransientStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransientStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransientStorage {
	mock := &TransientStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
