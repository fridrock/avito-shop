// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	storage "github.com/fridrock/avito-shop/storage"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ProductStorage is an autogenerated mock type for the ProductStorage type
type ProductStorage struct {
	mock.Mock
}

// Buy provides a mock function with given fields: _a0, _a1
func (_m *ProductStorage) Buy(_a0 uuid.UUID, _a1 storage.Product) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Buy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, storage.Product) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindProductByName provides a mock function with given fields: _a0
func (_m *ProductStorage) FindProductByName(_a0 string) (storage.Product, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindProductByName")
	}

	var r0 storage.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (storage.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) storage.Product); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(storage.Product)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductStorage creates a new instance of ProductStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductStorage {
	mock := &ProductStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
