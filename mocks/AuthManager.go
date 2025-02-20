// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	utils "github.com/fridrock/avito-shop/utils"
	mock "github.com/stretchr/testify/mock"
)

// AuthManager is an autogenerated mock type for the AuthManager type
type AuthManager struct {
	mock.Mock
}

// AuthMiddleware provides a mock function with given fields: h
func (_m *AuthManager) AuthMiddleware(h utils.HandlerWithError) utils.HandlerWithError {
	ret := _m.Called(h)

	if len(ret) == 0 {
		panic("no return value specified for AuthMiddleware")
	}

	var r0 utils.HandlerWithError
	if rf, ok := ret.Get(0).(func(utils.HandlerWithError) utils.HandlerWithError); ok {
		r0 = rf(h)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.HandlerWithError)
		}
	}

	return r0
}

// NewAuthManager creates a new instance of AuthManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthManager {
	mock := &AuthManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
