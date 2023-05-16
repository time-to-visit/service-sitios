// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-sites/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryResource is an autogenerated mock type for the IRepositoryResource type
type IRepositoryResource struct {
	mock.Mock
}

// DeleteResource provides a mock function with given fields: idSites, idResource
func (_m *IRepositoryResource) DeleteResource(idSites int, idResource int) error {
	ret := _m.Called(idSites, idResource)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(idSites, idResource)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOneResource provides a mock function with given fields: idSites, idResource
func (_m *IRepositoryResource) FindOneResource(idSites int, idResource int) (*entity.Resource, error) {
	ret := _m.Called(idSites, idResource)

	var r0 *entity.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (*entity.Resource, error)); ok {
		return rf(idSites, idResource)
	}
	if rf, ok := ret.Get(0).(func(int, int) *entity.Resource); ok {
		r0 = rf(idSites, idResource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(idSites, idResource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindResource provides a mock function with given fields: idSites
func (_m *IRepositoryResource) FindResource(idSites int) (*[]entity.Resource, error) {
	ret := _m.Called(idSites)

	var r0 *[]entity.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]entity.Resource, error)); ok {
		return rf(idSites)
	}
	if rf, ok := ret.Get(0).(func(int) *[]entity.Resource); ok {
		r0 = rf(idSites)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idSites)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertResource provides a mock function with given fields: resource
func (_m *IRepositoryResource) InsertResource(resource entity.Resource) (*entity.Resource, error) {
	ret := _m.Called(resource)

	var r0 *entity.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Resource) (*entity.Resource, error)); ok {
		return rf(resource)
	}
	if rf, ok := ret.Get(0).(func(entity.Resource) *entity.Resource); ok {
		r0 = rf(resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Resource) error); ok {
		r1 = rf(resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryResource interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryResource creates a new instance of IRepositoryResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryResource(t mockConstructorTestingTNewIRepositoryResource) *IRepositoryResource {
	mock := &IRepositoryResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}