// Code generated by mockery v2.42.1. DO NOT EDIT.

//go:build !compile

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	websocket "github.com/coder/websocket"
)

// MockConnection is an autogenerated mock type for the Connection type
type MockConnection struct {
	mock.Mock
}

type MockConnection_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnection) EXPECT() *MockConnection_Expecter {
	return &MockConnection_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: status, reason, closingCtx
func (_m *MockConnection) Close(status websocket.StatusCode, reason string, closingCtx ...context.Context) error {
	_va := make([]interface{}, len(closingCtx))
	for _i := range closingCtx {
		_va[_i] = closingCtx[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, status, reason)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(websocket.StatusCode, string, ...context.Context) error); ok {
		r0 = rf(status, reason, closingCtx...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnection_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockConnection_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - status websocket.StatusCode
//   - reason string
//   - closingCtx ...context.Context
func (_e *MockConnection_Expecter) Close(status interface{}, reason interface{}, closingCtx ...interface{}) *MockConnection_Close_Call {
	return &MockConnection_Close_Call{Call: _e.mock.On("Close",
		append([]interface{}{status, reason}, closingCtx...)...)}
}

func (_c *MockConnection_Close_Call) Run(run func(status websocket.StatusCode, reason string, closingCtx ...context.Context)) *MockConnection_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]context.Context, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(context.Context)
			}
		}
		run(args[0].(websocket.StatusCode), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockConnection_Close_Call) Return(_a0 error) *MockConnection_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnection_Close_Call) RunAndReturn(run func(websocket.StatusCode, string, ...context.Context) error) *MockConnection_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *MockConnection) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockConnection_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockConnection_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockConnection_Expecter) Context() *MockConnection_Context_Call {
	return &MockConnection_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockConnection_Context_Call) Run(run func()) *MockConnection_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConnection_Context_Call) Return(_a0 context.Context) *MockConnection_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnection_Context_Call) RunAndReturn(run func() context.Context) *MockConnection_Context_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *MockConnection) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockConnection_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type MockConnection_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *MockConnection_Expecter) ID() *MockConnection_ID_Call {
	return &MockConnection_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *MockConnection_ID_Call) Run(run func()) *MockConnection_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConnection_ID_Call) Return(_a0 string) *MockConnection_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnection_ID_Call) RunAndReturn(run func() string) *MockConnection_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: msgType, msg
func (_m *MockConnection) Send(msgType websocket.MessageType, msg []byte) error {
	ret := _m.Called(msgType, msg)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(websocket.MessageType, []byte) error); ok {
		r0 = rf(msgType, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnection_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockConnection_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - msgType websocket.MessageType
//   - msg []byte
func (_e *MockConnection_Expecter) Send(msgType interface{}, msg interface{}) *MockConnection_Send_Call {
	return &MockConnection_Send_Call{Call: _e.mock.On("Send", msgType, msg)}
}

func (_c *MockConnection_Send_Call) Run(run func(msgType websocket.MessageType, msg []byte)) *MockConnection_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(websocket.MessageType), args[1].([]byte))
	})
	return _c
}

func (_c *MockConnection_Send_Call) Return(_a0 error) *MockConnection_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnection_Send_Call) RunAndReturn(run func(websocket.MessageType, []byte) error) *MockConnection_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnection creates a new instance of MockConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnection {
	mock := &MockConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
