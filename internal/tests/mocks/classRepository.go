// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/williamwinkler/hs-card-service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ClassRepository is an autogenerated mock type for the ClassRepository type
type ClassRepository struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields:
func (_m *ClassRepository) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMany provides a mock function with given fields: classes
func (_m *ClassRepository) InsertMany(classes []domain.Class) error {
	ret := _m.Called(classes)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.Class) error); ok {
		r0 = rf(classes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClassRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassRepository creates a new instance of ClassRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassRepository(t mockConstructorTestingTNewClassRepository) *ClassRepository {
	mock := &ClassRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
