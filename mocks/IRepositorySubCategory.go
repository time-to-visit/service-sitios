// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "service-sites/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepositorySubCategory is an autogenerated mock type for the IRepositorySubCategory type
type IRepositorySubCategory struct {
	mock.Mock
}

// DeleteSubcategory provides a mock function with given fields: idCategory, idSubcategory
func (_m *IRepositorySubCategory) DeleteSubcategory(idCategory int, idSubcategory int) error {
	ret := _m.Called(idCategory, idSubcategory)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(idCategory, idSubcategory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOneSubcategory provides a mock function with given fields: idSubCategory
func (_m *IRepositorySubCategory) FindOneSubcategory(idSubCategory int) (*entity.Subcategory, error) {
	ret := _m.Called(idSubCategory)

	var r0 *entity.Subcategory
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*entity.Subcategory, error)); ok {
		return rf(idSubCategory)
	}
	if rf, ok := ret.Get(0).(func(int) *entity.Subcategory); ok {
		r0 = rf(idSubCategory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Subcategory)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idSubCategory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSubcategory provides a mock function with given fields: idCategory
func (_m *IRepositorySubCategory) FindSubcategory(idCategory int) (*[]entity.Subcategory, error) {
	ret := _m.Called(idCategory)

	var r0 *[]entity.Subcategory
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]entity.Subcategory, error)); ok {
		return rf(idCategory)
	}
	if rf, ok := ret.Get(0).(func(int) *[]entity.Subcategory); ok {
		r0 = rf(idCategory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Subcategory)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idCategory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertSubcategory provides a mock function with given fields: review
func (_m *IRepositorySubCategory) InsertSubcategory(review entity.Subcategory) (*entity.Subcategory, error) {
	ret := _m.Called(review)

	var r0 *entity.Subcategory
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Subcategory) (*entity.Subcategory, error)); ok {
		return rf(review)
	}
	if rf, ok := ret.Get(0).(func(entity.Subcategory) *entity.Subcategory); ok {
		r0 = rf(review)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Subcategory)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Subcategory) error); ok {
		r1 = rf(review)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepositorySubCategory interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepositorySubCategory creates a new instance of IRepositorySubCategory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepositorySubCategory(t mockConstructorTestingTNewIRepositorySubCategory) *IRepositorySubCategory {
	mock := &IRepositorySubCategory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
