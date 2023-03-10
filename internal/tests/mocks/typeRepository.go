// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/williamwinkler/hs-card-service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// TypeRepository is an autogenerated mock type for the TypeRepository type
type TypeRepository struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields:
func (_m *TypeRepository) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMany provides a mock function with given fields: types
func (_m *TypeRepository) InsertMany(types []domain.Type) error {
	ret := _m.Called(types)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.Type) error); ok {
		r0 = rf(types)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTypeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypeRepository creates a new instance of TypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypeRepository(t mockConstructorTestingTNewTypeRepository) *TypeRepository {
	mock := &TypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
