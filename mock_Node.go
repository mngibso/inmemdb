// Code generated by mockery v1.0.0. DO NOT EDIT.

package main

import (
	btree "github.com/google/btree"
	mock "github.com/stretchr/testify/mock"
)

// MockNode is an autogenerated mock type for the Node type
type MockNode struct {
	mock.Mock
}

// GetKey provides a mock function with given fields:
func (_m *MockNode) GetKey() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *MockNode) GetValue() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Less provides a mock function with given fields: than
func (_m *MockNode) Less(than btree.Item) bool {
	ret := _m.Called(than)

	var r0 bool
	if rf, ok := ret.Get(0).(func(btree.Item) bool); ok {
		r0 = rf(than)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}