// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-sites/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositoryContact is an autogenerated mock type for the IRepositoryContact type
type IRepositoryContact struct {
	mock.Mock
}

// DeleteContact provides a mock function with given fields: idSites, idContact
func (_m *IRepositoryContact) DeleteContact(idSites int, idContact int) error {
	ret := _m.Called(idSites, idContact)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(idSites, idContact)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertContact provides a mock function with given fields: contact
func (_m *IRepositoryContact) InsertContact(contact entity.Contact) (*entity.Contact, error) {
	ret := _m.Called(contact)

	var r0 *entity.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Contact) (*entity.Contact, error)); ok {
		return rf(contact)
	}
	if rf, ok := ret.Get(0).(func(entity.Contact) *entity.Contact); ok {
		r0 = rf(contact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Contact) error); ok {
		r1 = rf(contact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositoryContact interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositoryContact creates a new instance of IRepositoryContact. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositoryContact(t mockConstructorTestingTNewIRepositoryContact) *IRepositoryContact {
	mock := &IRepositoryContact{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}