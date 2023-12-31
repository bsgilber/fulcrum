// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// IMessage is an autogenerated mock type for the IMessage type
type IMessage struct {
	mock.Mock
}

type IMessage_Expecter struct {
	mock *mock.Mock
}

func (_m *IMessage) EXPECT() *IMessage_Expecter {
	return &IMessage_Expecter{mock: &_m.Mock}
}

// GeMessageId provides a mock function with given fields:
func (_m *IMessage) GeMessageId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IMessage_GeMessageId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GeMessageId'
type IMessage_GeMessageId_Call struct {
	*mock.Call
}

// GeMessageId is a helper method to define mock.On call
func (_e *IMessage_Expecter) GeMessageId() *IMessage_GeMessageId_Call {
	return &IMessage_GeMessageId_Call{Call: _e.mock.On("GeMessageId")}
}

func (_c *IMessage_GeMessageId_Call) Run(run func()) *IMessage_GeMessageId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessage_GeMessageId_Call) Return(_a0 string) *IMessage_GeMessageId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessage_GeMessageId_Call) RunAndReturn(run func() string) *IMessage_GeMessageId_Call {
	_c.Call.Return(run)
	return _c
}

// GetCreated provides a mock function with given fields:
func (_m *IMessage) GetCreated() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// IMessage_GetCreated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCreated'
type IMessage_GetCreated_Call struct {
	*mock.Call
}

// GetCreated is a helper method to define mock.On call
func (_e *IMessage_Expecter) GetCreated() *IMessage_GetCreated_Call {
	return &IMessage_GetCreated_Call{Call: _e.mock.On("GetCreated")}
}

func (_c *IMessage_GetCreated_Call) Run(run func()) *IMessage_GetCreated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessage_GetCreated_Call) Return(_a0 time.Time) *IMessage_GetCreated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessage_GetCreated_Call) RunAndReturn(run func() time.Time) *IMessage_GetCreated_Call {
	_c.Call.Return(run)
	return _c
}

// GetEventTypeName provides a mock function with given fields:
func (_m *IMessage) GetEventTypeName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IMessage_GetEventTypeName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventTypeName'
type IMessage_GetEventTypeName_Call struct {
	*mock.Call
}

// GetEventTypeName is a helper method to define mock.On call
func (_e *IMessage_Expecter) GetEventTypeName() *IMessage_GetEventTypeName_Call {
	return &IMessage_GetEventTypeName_Call{Call: _e.mock.On("GetEventTypeName")}
}

func (_c *IMessage_GetEventTypeName_Call) Run(run func()) *IMessage_GetEventTypeName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessage_GetEventTypeName_Call) Return(_a0 string) *IMessage_GetEventTypeName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessage_GetEventTypeName_Call) RunAndReturn(run func() string) *IMessage_GetEventTypeName_Call {
	_c.Call.Return(run)
	return _c
}

// IsMessage provides a mock function with given fields:
func (_m *IMessage) IsMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IMessage_IsMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsMessage'
type IMessage_IsMessage_Call struct {
	*mock.Call
}

// IsMessage is a helper method to define mock.On call
func (_e *IMessage_Expecter) IsMessage() *IMessage_IsMessage_Call {
	return &IMessage_IsMessage_Call{Call: _e.mock.On("IsMessage")}
}

func (_c *IMessage_IsMessage_Call) Run(run func()) *IMessage_IsMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessage_IsMessage_Call) Return(_a0 bool) *IMessage_IsMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessage_IsMessage_Call) RunAndReturn(run func() bool) *IMessage_IsMessage_Call {
	_c.Call.Return(run)
	return _c
}

// SetEventTypeName provides a mock function with given fields: _a0
func (_m *IMessage) SetEventTypeName(_a0 string) {
	_m.Called(_a0)
}

// IMessage_SetEventTypeName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEventTypeName'
type IMessage_SetEventTypeName_Call struct {
	*mock.Call
}

// SetEventTypeName is a helper method to define mock.On call
//   - _a0 string
func (_e *IMessage_Expecter) SetEventTypeName(_a0 interface{}) *IMessage_SetEventTypeName_Call {
	return &IMessage_SetEventTypeName_Call{Call: _e.mock.On("SetEventTypeName", _a0)}
}

func (_c *IMessage_SetEventTypeName_Call) Run(run func(_a0 string)) *IMessage_SetEventTypeName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IMessage_SetEventTypeName_Call) Return() *IMessage_SetEventTypeName_Call {
	_c.Call.Return()
	return _c
}

func (_c *IMessage_SetEventTypeName_Call) RunAndReturn(run func(string)) *IMessage_SetEventTypeName_Call {
	_c.Call.Return(run)
	return _c
}

// NewIMessage creates a new instance of IMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMessage(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMessage {
	mock := &IMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
