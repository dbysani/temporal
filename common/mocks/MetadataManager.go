// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mocks

import mock "github.com/stretchr/testify/mock"
import persistence "github.com/uber/cadence/common/persistence"

// MetadataManager is an autogenerated mock type for the MetadataManager type
type MetadataManager struct {
	mock.Mock
}

// CreateDomain provides a mock function with given fields: request
func (_m *MetadataManager) CreateDomain(request *persistence.CreateDomainRequest) (*persistence.CreateDomainResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.CreateDomainResponse
	if rf, ok := ret.Get(0).(func(*persistence.CreateDomainRequest) *persistence.CreateDomainResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CreateDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.CreateDomainRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDomain provides a mock function with given fields: request
func (_m *MetadataManager) DeleteDomain(request *persistence.DeleteDomainRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.DeleteDomainRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDomainByName provides a mock function with given fields: request
func (_m *MetadataManager) DeleteDomainByName(request *persistence.DeleteDomainByNameRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.DeleteDomainByNameRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDomain provides a mock function with given fields: request
func (_m *MetadataManager) GetDomain(request *persistence.GetDomainRequest) (*persistence.GetDomainResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetDomainResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetDomainRequest) *persistence.GetDomainResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetDomainRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDomain provides a mock function with given fields: request
func (_m *MetadataManager) UpdateDomain(request *persistence.UpdateDomainRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.UpdateDomainRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}