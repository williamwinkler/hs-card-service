// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/williamwinkler/hs-card-service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// SetRepository is an autogenerated mock type for the SetRepository type
type SetRepository struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields:
func (_m *SetRepository) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMany provides a mock function with given fields: sets
func (_m *SetRepository) InsertMany(sets []domain.Set) error {
	ret := _m.Called(sets)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.Set) error); ok {
		r0 = rf(sets)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSetRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewSetRepository creates a new instance of SetRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSetRepository(t mockConstructorTestingTNewSetRepository) *SetRepository {
	mock := &SetRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
