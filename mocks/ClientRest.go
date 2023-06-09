// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	rest "github.com/sendgrid/rest"
	mock "github.com/stretchr/testify/mock"
)

// ClientRest is an autogenerated mock type for the ClientRest type
type ClientRest struct {
	mock.Mock
}

// Delete provides a mock function with given fields: endpoint, body
func (_m *ClientRest) Delete(endpoint string, body interface{}) (*rest.Response, error) {
	ret := _m.Called(endpoint, body)

	var r0 *rest.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (*rest.Response, error)); ok {
		return rf(endpoint, body)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) *rest.Response); ok {
		r0 = rf(endpoint, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(endpoint, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: endpoint
func (_m *ClientRest) Get(endpoint string) (*rest.Response, error) {
	ret := _m.Called(endpoint)

	var r0 *rest.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*rest.Response, error)); ok {
		return rf(endpoint)
	}
	if rf, ok := ret.Get(0).(func(string) *rest.Response); ok {
		r0 = rf(endpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(endpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: endpoint, body
func (_m *ClientRest) Post(endpoint string, body interface{}) (*rest.Response, error) {
	ret := _m.Called(endpoint, body)

	var r0 *rest.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (*rest.Response, error)); ok {
		return rf(endpoint, body)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) *rest.Response); ok {
		r0 = rf(endpoint, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(endpoint, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: endpoint, body
func (_m *ClientRest) Put(endpoint string, body interface{}) (*rest.Response, error) {
	ret := _m.Called(endpoint, body)

	var r0 *rest.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (*rest.Response, error)); ok {
		return rf(endpoint, body)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) *rest.Response); ok {
		r0 = rf(endpoint, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(endpoint, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientRest interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientRest creates a new instance of ClientRest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientRest(t mockConstructorTestingTNewClientRest) *ClientRest {
	mock := &ClientRest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
