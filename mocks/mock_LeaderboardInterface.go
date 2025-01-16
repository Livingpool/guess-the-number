// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/Livingpool/service"
	mock "github.com/stretchr/testify/mock"
)

// MockLeaderboardInterface is an autogenerated mock type for the LeaderboardInterface type
type MockLeaderboardInterface struct {
	mock.Mock
}

type MockLeaderboardInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLeaderboardInterface) EXPECT() *MockLeaderboardInterface_Expecter {
	return &MockLeaderboardInterface_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *MockLeaderboardInterface) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLeaderboardInterface_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockLeaderboardInterface_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockLeaderboardInterface_Expecter) Close() *MockLeaderboardInterface_Close_Call {
	return &MockLeaderboardInterface_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockLeaderboardInterface_Close_Call) Run(run func()) *MockLeaderboardInterface_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLeaderboardInterface_Close_Call) Return(_a0 error) *MockLeaderboardInterface_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLeaderboardInterface_Close_Call) RunAndReturn(run func() error) *MockLeaderboardInterface_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, boardId, name
func (_m *MockLeaderboardInterface) Get(ctx context.Context, boardId int, name string) ([]service.Record, error) {
	ret := _m.Called(ctx, boardId, name)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []service.Record
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) ([]service.Record, error)); ok {
		return rf(ctx, boardId, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string) []service.Record); ok {
		r0 = rf(ctx, boardId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]service.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, boardId, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLeaderboardInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockLeaderboardInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - boardId int
//   - name string
func (_e *MockLeaderboardInterface_Expecter) Get(ctx interface{}, boardId interface{}, name interface{}) *MockLeaderboardInterface_Get_Call {
	return &MockLeaderboardInterface_Get_Call{Call: _e.mock.On("Get", ctx, boardId, name)}
}

func (_c *MockLeaderboardInterface_Get_Call) Run(run func(ctx context.Context, boardId int, name string)) *MockLeaderboardInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string))
	})
	return _c
}

func (_c *MockLeaderboardInterface_Get_Call) Return(_a0 []service.Record, _a1 error) *MockLeaderboardInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLeaderboardInterface_Get_Call) RunAndReturn(run func(context.Context, int, string) ([]service.Record, error)) *MockLeaderboardInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with no fields
func (_m *MockLeaderboardInterface) Init() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLeaderboardInterface_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockLeaderboardInterface_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *MockLeaderboardInterface_Expecter) Init() *MockLeaderboardInterface_Init_Call {
	return &MockLeaderboardInterface_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *MockLeaderboardInterface_Init_Call) Run(run func()) *MockLeaderboardInterface_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLeaderboardInterface_Init_Call) Return(_a0 error) *MockLeaderboardInterface_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLeaderboardInterface_Init_Call) RunAndReturn(run func() error) *MockLeaderboardInterface_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: ctx, data
func (_m *MockLeaderboardInterface) Insert(ctx context.Context, data service.Record) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, service.Record) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLeaderboardInterface_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockLeaderboardInterface_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - ctx context.Context
//   - data service.Record
func (_e *MockLeaderboardInterface_Expecter) Insert(ctx interface{}, data interface{}) *MockLeaderboardInterface_Insert_Call {
	return &MockLeaderboardInterface_Insert_Call{Call: _e.mock.On("Insert", ctx, data)}
}

func (_c *MockLeaderboardInterface_Insert_Call) Run(run func(ctx context.Context, data service.Record)) *MockLeaderboardInterface_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(service.Record))
	})
	return _c
}

func (_c *MockLeaderboardInterface_Insert_Call) Return(_a0 error) *MockLeaderboardInterface_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLeaderboardInterface_Insert_Call) RunAndReturn(run func(context.Context, service.Record) error) *MockLeaderboardInterface_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLeaderboardInterface creates a new instance of MockLeaderboardInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLeaderboardInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLeaderboardInterface {
	mock := &MockLeaderboardInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}