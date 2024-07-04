// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	storage_v1 "github.com/jaegertracing/jaeger/proto-gen/storage_v1"
	mock "github.com/stretchr/testify/mock"
)

// DependenciesReaderPluginServer is an autogenerated mock type for the DependenciesReaderPluginServer type
type DependenciesReaderPluginServer struct {
	mock.Mock
}

// GetDependencies provides a mock function with given fields: _a0, _a1
func (_m *DependenciesReaderPluginServer) GetDependencies(_a0 context.Context, _a1 *storage_v1.GetDependenciesRequest) (*storage_v1.GetDependenciesResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDependencies")
	}

	var r0 *storage_v1.GetDependenciesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetDependenciesRequest) (*storage_v1.GetDependenciesResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetDependenciesRequest) *storage_v1.GetDependenciesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage_v1.GetDependenciesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage_v1.GetDependenciesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDependenciesReaderPluginServer creates a new instance of DependenciesReaderPluginServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDependenciesReaderPluginServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *DependenciesReaderPluginServer {
	mock := &DependenciesReaderPluginServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
