// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	driver "database/sql/driver"

	mock "github.com/stretchr/testify/mock"

	types "github.com/plutov/formulosity/pkg/types"
)

// Answer is an autogenerated mock type for the Answer type
type Answer struct {
	mock.Mock
}

type Answer_Expecter struct {
	mock *mock.Mock
}

func (_m *Answer) EXPECT() *Answer_Expecter {
	return &Answer_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: q
func (_m *Answer) Validate(q types.Question) error {
	ret := _m.Called(q)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Question) error); ok {
		r0 = rf(q)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Answer_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type Answer_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - q types.Question
func (_e *Answer_Expecter) Validate(q interface{}) *Answer_Validate_Call {
	return &Answer_Validate_Call{Call: _e.mock.On("Validate", q)}
}

func (_c *Answer_Validate_Call) Run(run func(q types.Question)) *Answer_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Question))
	})
	return _c
}

func (_c *Answer_Validate_Call) Return(_a0 error) *Answer_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Answer_Validate_Call) RunAndReturn(run func(types.Question) error) *Answer_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// Value provides a mock function with given fields:
func (_m *Answer) Value() (driver.Value, error) {
	ret := _m.Called()

	var r0 driver.Value
	var r1 error
	if rf, ok := ret.Get(0).(func() (driver.Value, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() driver.Value); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Value)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Answer_Value_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Value'
type Answer_Value_Call struct {
	*mock.Call
}

// Value is a helper method to define mock.On call
func (_e *Answer_Expecter) Value() *Answer_Value_Call {
	return &Answer_Value_Call{Call: _e.mock.On("Value")}
}

func (_c *Answer_Value_Call) Run(run func()) *Answer_Value_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Answer_Value_Call) Return(_a0 driver.Value, _a1 error) *Answer_Value_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Answer_Value_Call) RunAndReturn(run func() (driver.Value, error)) *Answer_Value_Call {
	_c.Call.Return(run)
	return _c
}

// NewAnswer creates a new instance of Answer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAnswer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Answer {
	mock := &Answer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
