// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	consumer "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/consumer"

	mock "github.com/stretchr/testify/mock"

	types "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/types"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

type Consumer_Expecter struct {
	mock *mock.Mock
}

func (_m *Consumer) EXPECT() *Consumer_Expecter {
	return &Consumer_Expecter{mock: &_m.Mock}
}

// AddMessageConsumedHandler provides a mock function with given fields: _a0
func (_m *Consumer) IsConsumed(_a0 func(types.IMessage)) {
	_m.Called(_a0)
}

// Consumer_AddMessageConsumedHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConsumed'
type Consumer_AddMessageConsumedHandler_Call struct {
	*mock.Call
}

// AddMessageConsumedHandler is a helper method to define mock.On call
//   - _a0 func(types.IMessage)
func (_e *Consumer_Expecter) AddMessageConsumedHandler(_a0 interface{}) *Consumer_AddMessageConsumedHandler_Call {
	return &Consumer_AddMessageConsumedHandler_Call{Call: _e.mock.On("IsConsumed", _a0)}
}

func (_c *Consumer_AddMessageConsumedHandler_Call) Run(run func(_a0 func(types.IMessage))) *Consumer_AddMessageConsumedHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(types.IMessage)))
	})
	return _c
}

func (_c *Consumer_AddMessageConsumedHandler_Call) Return() *Consumer_AddMessageConsumedHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *Consumer_AddMessageConsumedHandler_Call) RunAndReturn(run func(func(types.IMessage))) *Consumer_AddMessageConsumedHandler_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectHandler provides a mock function with given fields: handler
func (_m *Consumer) ConnectHandler(handler consumer.ConsumerHandler) {
	_m.Called(handler)
}

// Consumer_ConnectHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectHandler'
type Consumer_ConnectHandler_Call struct {
	*mock.Call
}

// ConnectHandler is a helper method to define mock.On call
//   - handler consumer.ConsumerHandler
func (_e *Consumer_Expecter) ConnectHandler(handler interface{}) *Consumer_ConnectHandler_Call {
	return &Consumer_ConnectHandler_Call{Call: _e.mock.On("ConnectHandler", handler)}
}

func (_c *Consumer_ConnectHandler_Call) Run(run func(handler consumer.ConsumerHandler)) *Consumer_ConnectHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(consumer.ConsumerHandler))
	})
	return _c
}

func (_c *Consumer_ConnectHandler_Call) Return() *Consumer_ConnectHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *Consumer_ConnectHandler_Call) RunAndReturn(run func(consumer.ConsumerHandler)) *Consumer_ConnectHandler_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *Consumer) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consumer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Consumer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Consumer_Expecter) Start(ctx interface{}) *Consumer_Start_Call {
	return &Consumer_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Consumer_Start_Call) Run(run func(ctx context.Context)) *Consumer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Consumer_Start_Call) Return(_a0 error) *Consumer_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Consumer_Start_Call) RunAndReturn(run func(context.Context) error) *Consumer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *Consumer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consumer_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Consumer_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Consumer_Expecter) Stop() *Consumer_Stop_Call {
	return &Consumer_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Consumer_Stop_Call) Run(run func()) *Consumer_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Consumer_Stop_Call) Return(_a0 error) *Consumer_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Consumer_Stop_Call) RunAndReturn(run func() error) *Consumer_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewConsumer creates a new instance of Consumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
