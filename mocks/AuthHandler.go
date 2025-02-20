// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AuthHandler is an autogenerated mock type for the AuthHandler type
type AuthHandler struct {
	mock.Mock
}

// Auth provides a mock function with given fields: w, r
func (_m *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) (int, error) {
	ret := _m.Called(w, r)

	if len(ret) == 0 {
		panic("no return value specified for Auth")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) (int, error)); ok {
		return rf(w, r)
	}
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) int); ok {
		r0 = rf(w, r)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(http.ResponseWriter, *http.Request) error); ok {
		r1 = rf(w, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthHandler creates a new instance of AuthHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthHandler {
	mock := &AuthHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
