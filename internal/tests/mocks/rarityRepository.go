// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/williamwinkler/hs-card-service/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// RarityRepository is an autogenerated mock type for the RarityRepository type
type RarityRepository struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields:
func (_m *RarityRepository) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMany provides a mock function with given fields: raritys
func (_m *RarityRepository) InsertMany(raritys []domain.Rarity) error {
	ret := _m.Called(raritys)

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.Rarity) error); ok {
		r0 = rf(raritys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRarityRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRarityRepository creates a new instance of RarityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRarityRepository(t mockConstructorTestingTNewRarityRepository) *RarityRepository {
	mock := &RarityRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}