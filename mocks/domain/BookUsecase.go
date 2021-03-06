// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/golang-common-packages/template/domain"
import mock "github.com/stretchr/testify/mock"
import reflect "reflect"

// BookUsecase is an autogenerated mock type for the BookUsecase type
type BookUsecase struct {
	mock.Mock
}

// DeleteBook provides a mock function with given fields: bookID
func (_m *BookUsecase) DeleteBook(bookID string) (interface{}, error) {
	ret := _m.Called(bookID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBooks provides a mock function with given fields: books
func (_m *BookUsecase) InsertBooks(books *[]domain.Book) (interface{}, error) {
	ret := _m.Called(books)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*[]domain.Book) interface{}); ok {
		r0 = rf(books)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*[]domain.Book) error); ok {
		r1 = rf(books)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBooks provides a mock function with given fields: limit, dataModel
func (_m *BookUsecase) ListBooks(limit int64, dataModel reflect.Type) (interface{}, error) {
	ret := _m.Called(limit, dataModel)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(int64, reflect.Type) interface{}); ok {
		r0 = rf(limit, dataModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, reflect.Type) error); ok {
		r1 = rf(limit, dataModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: update
func (_m *BookUsecase) UpdateBook(update domain.Book) (interface{}, error) {
	ret := _m.Called(update)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(domain.Book) interface{}); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Book) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
