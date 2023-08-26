// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	domain "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/core/domain"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/satori/go.uuid"
)

// IAggregateRoot is an autogenerated mock type for the IAggregateRoot type
type IAggregateRoot struct {
	mock.Mock
}

type IAggregateRoot_Expecter struct {
	mock *mock.Mock
}

func (_m *IAggregateRoot) EXPECT() *IAggregateRoot_Expecter {
	return &IAggregateRoot_Expecter{mock: &_m.Mock}
}

// AddDomainEvents provides a mock function with given fields: event
func (_m *IAggregateRoot) AddDomainEvents(event domain.IDomainEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IDomainEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IAggregateRoot_AddDomainEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddDomainEvents'
type IAggregateRoot_AddDomainEvents_Call struct {
	*mock.Call
}

// AddDomainEvents is a helper method to define mock.On call
//   - event domain.IDomainEvent
func (_e *IAggregateRoot_Expecter) AddDomainEvents(event interface{}) *IAggregateRoot_AddDomainEvents_Call {
	return &IAggregateRoot_AddDomainEvents_Call{Call: _e.mock.On("AddDomainEvents", event)}
}

func (_c *IAggregateRoot_AddDomainEvents_Call) Run(run func(event domain.IDomainEvent)) *IAggregateRoot_AddDomainEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.IDomainEvent))
	})
	return _c
}

func (_c *IAggregateRoot_AddDomainEvents_Call) Return(_a0 error) *IAggregateRoot_AddDomainEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_AddDomainEvents_Call) RunAndReturn(run func(domain.IDomainEvent) error) *IAggregateRoot_AddDomainEvents_Call {
	_c.Call.Return(run)
	return _c
}

// CreatedAt provides a mock function with given fields:
func (_m *IAggregateRoot) CreatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// IAggregateRoot_CreatedAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatedAt'
type IAggregateRoot_CreatedAt_Call struct {
	*mock.Call
}

// CreatedAt is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) CreatedAt() *IAggregateRoot_CreatedAt_Call {
	return &IAggregateRoot_CreatedAt_Call{Call: _e.mock.On("CreatedAt")}
}

func (_c *IAggregateRoot_CreatedAt_Call) Run(run func()) *IAggregateRoot_CreatedAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_CreatedAt_Call) Return(_a0 time.Time) *IAggregateRoot_CreatedAt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_CreatedAt_Call) RunAndReturn(run func() time.Time) *IAggregateRoot_CreatedAt_Call {
	_c.Call.Return(run)
	return _c
}

// GetUncommittedEvents provides a mock function with given fields:
func (_m *IAggregateRoot) GetUncommittedEvents() []domain.IDomainEvent {
	ret := _m.Called()

	var r0 []domain.IDomainEvent
	if rf, ok := ret.Get(0).(func() []domain.IDomainEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IDomainEvent)
		}
	}

	return r0
}

// IAggregateRoot_GetUncommittedEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUncommittedEvents'
type IAggregateRoot_GetUncommittedEvents_Call struct {
	*mock.Call
}

// GetUncommittedEvents is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) GetUncommittedEvents() *IAggregateRoot_GetUncommittedEvents_Call {
	return &IAggregateRoot_GetUncommittedEvents_Call{Call: _e.mock.On("GetUncommittedEvents")}
}

func (_c *IAggregateRoot_GetUncommittedEvents_Call) Run(run func()) *IAggregateRoot_GetUncommittedEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_GetUncommittedEvents_Call) Return(_a0 []domain.IDomainEvent) *IAggregateRoot_GetUncommittedEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_GetUncommittedEvents_Call) RunAndReturn(run func() []domain.IDomainEvent) *IAggregateRoot_GetUncommittedEvents_Call {
	_c.Call.Return(run)
	return _c
}

// HasUncommittedEvents provides a mock function with given fields:
func (_m *IAggregateRoot) HasUncommittedEvents() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IAggregateRoot_HasUncommittedEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasUncommittedEvents'
type IAggregateRoot_HasUncommittedEvents_Call struct {
	*mock.Call
}

// HasUncommittedEvents is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) HasUncommittedEvents() *IAggregateRoot_HasUncommittedEvents_Call {
	return &IAggregateRoot_HasUncommittedEvents_Call{Call: _e.mock.On("HasUncommittedEvents")}
}

func (_c *IAggregateRoot_HasUncommittedEvents_Call) Run(run func()) *IAggregateRoot_HasUncommittedEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_HasUncommittedEvents_Call) Return(_a0 bool) *IAggregateRoot_HasUncommittedEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_HasUncommittedEvents_Call) RunAndReturn(run func() bool) *IAggregateRoot_HasUncommittedEvents_Call {
	_c.Call.Return(run)
	return _c
}

// Id provides a mock function with given fields:
func (_m *IAggregateRoot) Id() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// IAggregateRoot_Id_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Id'
type IAggregateRoot_Id_Call struct {
	*mock.Call
}

// Id is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) Id() *IAggregateRoot_Id_Call {
	return &IAggregateRoot_Id_Call{Call: _e.mock.On("Id")}
}

func (_c *IAggregateRoot_Id_Call) Run(run func()) *IAggregateRoot_Id_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_Id_Call) Return(_a0 uuid.UUID) *IAggregateRoot_Id_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_Id_Call) RunAndReturn(run func() uuid.UUID) *IAggregateRoot_Id_Call {
	_c.Call.Return(run)
	return _c
}

// MarkUncommittedEventAsCommitted provides a mock function with given fields:
func (_m *IAggregateRoot) MarkUncommittedEventAsCommitted() {
	_m.Called()
}

// IAggregateRoot_MarkUncommittedEventAsCommitted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkUncommittedEventAsCommitted'
type IAggregateRoot_MarkUncommittedEventAsCommitted_Call struct {
	*mock.Call
}

// MarkUncommittedEventAsCommitted is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) MarkUncommittedEventAsCommitted() *IAggregateRoot_MarkUncommittedEventAsCommitted_Call {
	return &IAggregateRoot_MarkUncommittedEventAsCommitted_Call{Call: _e.mock.On("MarkUncommittedEventAsCommitted")}
}

func (_c *IAggregateRoot_MarkUncommittedEventAsCommitted_Call) Run(run func()) *IAggregateRoot_MarkUncommittedEventAsCommitted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_MarkUncommittedEventAsCommitted_Call) Return() *IAggregateRoot_MarkUncommittedEventAsCommitted_Call {
	_c.Call.Return()
	return _c
}

func (_c *IAggregateRoot_MarkUncommittedEventAsCommitted_Call) RunAndReturn(run func()) *IAggregateRoot_MarkUncommittedEventAsCommitted_Call {
	_c.Call.Return(run)
	return _c
}

// OriginalVersion provides a mock function with given fields:
func (_m *IAggregateRoot) OriginalVersion() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// IAggregateRoot_OriginalVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OriginalVersion'
type IAggregateRoot_OriginalVersion_Call struct {
	*mock.Call
}

// OriginalVersion is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) OriginalVersion() *IAggregateRoot_OriginalVersion_Call {
	return &IAggregateRoot_OriginalVersion_Call{Call: _e.mock.On("OriginalVersion")}
}

func (_c *IAggregateRoot_OriginalVersion_Call) Run(run func()) *IAggregateRoot_OriginalVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_OriginalVersion_Call) Return(_a0 int64) *IAggregateRoot_OriginalVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_OriginalVersion_Call) RunAndReturn(run func() int64) *IAggregateRoot_OriginalVersion_Call {
	_c.Call.Return(run)
	return _c
}

// SetEntityType provides a mock function with given fields: entityType
func (_m *IAggregateRoot) SetEntityType(entityType string) {
	_m.Called(entityType)
}

// IAggregateRoot_SetEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEntityType'
type IAggregateRoot_SetEntityType_Call struct {
	*mock.Call
}

// SetEntityType is a helper method to define mock.On call
//   - entityType string
func (_e *IAggregateRoot_Expecter) SetEntityType(entityType interface{}) *IAggregateRoot_SetEntityType_Call {
	return &IAggregateRoot_SetEntityType_Call{Call: _e.mock.On("SetEntityType", entityType)}
}

func (_c *IAggregateRoot_SetEntityType_Call) Run(run func(entityType string)) *IAggregateRoot_SetEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IAggregateRoot_SetEntityType_Call) Return() *IAggregateRoot_SetEntityType_Call {
	_c.Call.Return()
	return _c
}

func (_c *IAggregateRoot_SetEntityType_Call) RunAndReturn(run func(string)) *IAggregateRoot_SetEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// SetId provides a mock function with given fields: id
func (_m *IAggregateRoot) SetId(id uuid.UUID) {
	_m.Called(id)
}

// IAggregateRoot_SetId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetId'
type IAggregateRoot_SetId_Call struct {
	*mock.Call
}

// SetId is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *IAggregateRoot_Expecter) SetId(id interface{}) *IAggregateRoot_SetId_Call {
	return &IAggregateRoot_SetId_Call{Call: _e.mock.On("SetId", id)}
}

func (_c *IAggregateRoot_SetId_Call) Run(run func(id uuid.UUID)) *IAggregateRoot_SetId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *IAggregateRoot_SetId_Call) Return() *IAggregateRoot_SetId_Call {
	_c.Call.Return()
	return _c
}

func (_c *IAggregateRoot_SetId_Call) RunAndReturn(run func(uuid.UUID)) *IAggregateRoot_SetId_Call {
	_c.Call.Return(run)
	return _c
}

// SetUpdatedAt provides a mock function with given fields: updatedAt
func (_m *IAggregateRoot) SetUpdatedAt(updatedAt time.Time) {
	_m.Called(updatedAt)
}

// IAggregateRoot_SetUpdatedAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUpdatedAt'
type IAggregateRoot_SetUpdatedAt_Call struct {
	*mock.Call
}

// SetUpdatedAt is a helper method to define mock.On call
//   - updatedAt time.Time
func (_e *IAggregateRoot_Expecter) SetUpdatedAt(updatedAt interface{}) *IAggregateRoot_SetUpdatedAt_Call {
	return &IAggregateRoot_SetUpdatedAt_Call{Call: _e.mock.On("SetUpdatedAt", updatedAt)}
}

func (_c *IAggregateRoot_SetUpdatedAt_Call) Run(run func(updatedAt time.Time)) *IAggregateRoot_SetUpdatedAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time))
	})
	return _c
}

func (_c *IAggregateRoot_SetUpdatedAt_Call) Return() *IAggregateRoot_SetUpdatedAt_Call {
	_c.Call.Return()
	return _c
}

func (_c *IAggregateRoot_SetUpdatedAt_Call) RunAndReturn(run func(time.Time)) *IAggregateRoot_SetUpdatedAt_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *IAggregateRoot) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IAggregateRoot_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type IAggregateRoot_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) String() *IAggregateRoot_String_Call {
	return &IAggregateRoot_String_Call{Call: _e.mock.On("String")}
}

func (_c *IAggregateRoot_String_Call) Run(run func()) *IAggregateRoot_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_String_Call) Return(_a0 string) *IAggregateRoot_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_String_Call) RunAndReturn(run func() string) *IAggregateRoot_String_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatedAt provides a mock function with given fields:
func (_m *IAggregateRoot) UpdatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// IAggregateRoot_UpdatedAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatedAt'
type IAggregateRoot_UpdatedAt_Call struct {
	*mock.Call
}

// UpdatedAt is a helper method to define mock.On call
func (_e *IAggregateRoot_Expecter) UpdatedAt() *IAggregateRoot_UpdatedAt_Call {
	return &IAggregateRoot_UpdatedAt_Call{Call: _e.mock.On("UpdatedAt")}
}

func (_c *IAggregateRoot_UpdatedAt_Call) Run(run func()) *IAggregateRoot_UpdatedAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IAggregateRoot_UpdatedAt_Call) Return(_a0 time.Time) *IAggregateRoot_UpdatedAt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAggregateRoot_UpdatedAt_Call) RunAndReturn(run func() time.Time) *IAggregateRoot_UpdatedAt_Call {
	_c.Call.Return(run)
	return _c
}

// NewIAggregateRoot creates a new instance of IAggregateRoot. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAggregateRoot(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAggregateRoot {
	mock := &IAggregateRoot{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}