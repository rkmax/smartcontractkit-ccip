// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	ccipdata "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"

	context "context"

	internal "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

// OnRampReader is an autogenerated mock type for the OnRampReader type
type OnRampReader struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *OnRampReader) Address() (common.Address, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func() (common.Address, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields: qopts
func (_m *OnRampReader) Close(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDynamicConfig provides a mock function with given fields:
func (_m *OnRampReader) GetDynamicConfig() (ccipdata.OnRampDynamicConfig, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDynamicConfig")
	}

	var r0 ccipdata.OnRampDynamicConfig
	var r1 error
	if rf, ok := ret.Get(0).(func() (ccipdata.OnRampDynamicConfig, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ccipdata.OnRampDynamicConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ccipdata.OnRampDynamicConfig)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSendRequestsBetweenSeqNums provides a mock function with given fields: ctx, seqNumMin, seqNumMax, finalized
func (_m *OnRampReader) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin uint64, seqNumMax uint64, finalized bool) ([]ccipdata.Event[internal.EVM2EVMMessage], error) {
	ret := _m.Called(ctx, seqNumMin, seqNumMax, finalized)

	if len(ret) == 0 {
		panic("no return value specified for GetSendRequestsBetweenSeqNums")
	}

	var r0 []ccipdata.Event[internal.EVM2EVMMessage]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, bool) ([]ccipdata.Event[internal.EVM2EVMMessage], error)); ok {
		return rf(ctx, seqNumMin, seqNumMax, finalized)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, bool) []ccipdata.Event[internal.EVM2EVMMessage]); ok {
		r0 = rf(ctx, seqNumMin, seqNumMax, finalized)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipdata.Event[internal.EVM2EVMMessage])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, bool) error); ok {
		r1 = rf(ctx, seqNumMin, seqNumMax, finalized)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterFilters provides a mock function with given fields: qopts
func (_m *OnRampReader) RegisterFilters(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RouterAddress provides a mock function with given fields:
func (_m *OnRampReader) RouterAddress() (common.Address, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RouterAddress")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func() (common.Address, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOnRampReader creates a new instance of OnRampReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOnRampReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *OnRampReader {
	mock := &OnRampReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
