//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/gce_vm_image_import/domain (interfaces: ScratchBucketCreatorInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockScratchBucketCreatorInterface is a mock of ScratchBucketCreatorInterface interface
type MockScratchBucketCreatorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockScratchBucketCreatorInterfaceMockRecorder
}

// MockScratchBucketCreatorInterfaceMockRecorder is the mock recorder for MockScratchBucketCreatorInterface
type MockScratchBucketCreatorInterfaceMockRecorder struct {
	mock *MockScratchBucketCreatorInterface
}

// NewMockScratchBucketCreatorInterface creates a new mock instance
func NewMockScratchBucketCreatorInterface(ctrl *gomock.Controller) *MockScratchBucketCreatorInterface {
	mock := &MockScratchBucketCreatorInterface{ctrl: ctrl}
	mock.recorder = &MockScratchBucketCreatorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScratchBucketCreatorInterface) EXPECT() *MockScratchBucketCreatorInterfaceMockRecorder {
	return m.recorder
}

// CreateScratchBucket mocks base method
func (m *MockScratchBucketCreatorInterface) CreateScratchBucket(arg0, arg1 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScratchBucket", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateScratchBucket indicates an expected call of CreateScratchBucket
func (mr *MockScratchBucketCreatorInterfaceMockRecorder) CreateScratchBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScratchBucket", reflect.TypeOf((*MockScratchBucketCreatorInterface)(nil).CreateScratchBucket), arg0, arg1)
}
