/*
*Copyright (c) 2019-2021, Alibaba Group Holding Limited;
*Licensed under the Apache License, Version 2.0 (the "License");
*you may not use this file except in compliance with the License.
*You may obtain a copy of the License at

*   http://www.apache.org/licenses/LICENSE-2.0

*Unless required by applicable law or agreed to in writing, software
*distributed under the License is distributed on an "AS IS" BASIS,
*WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*See the License for the specific language governing permissions and
*limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: ./wfengine/wf_runtime_memento.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWfRuntimeMementoStorage is a mock of WfRuntimeMementoStorage interface.
type MockWfRuntimeMementoStorage struct {
	ctrl     *gomock.Controller
	recorder *MockWfRuntimeMementoStorageMockRecorder
}

// MockWfRuntimeMementoStorageMockRecorder is the mock recorder for MockWfRuntimeMementoStorage.
type MockWfRuntimeMementoStorageMockRecorder struct {
	mock *MockWfRuntimeMementoStorage
}

// NewMockWfRuntimeMementoStorage creates a new mock instance.
func NewMockWfRuntimeMementoStorage(ctrl *gomock.Controller) *MockWfRuntimeMementoStorage {
	mock := &MockWfRuntimeMementoStorage{ctrl: ctrl}
	mock.recorder = &MockWfRuntimeMementoStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWfRuntimeMementoStorage) EXPECT() *MockWfRuntimeMementoStorageMockRecorder {
	return m.recorder
}

// LoadMementoMap mocks base method.
func (m *MockWfRuntimeMementoStorage) LoadMementoMap(careTakerName string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMementoMap", careTakerName)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMementoMap indicates an expected call of LoadMementoMap.
func (mr *MockWfRuntimeMementoStorageMockRecorder) LoadMementoMap(careTakerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMementoMap", reflect.TypeOf((*MockWfRuntimeMementoStorage)(nil).LoadMementoMap), careTakerName)
}

// Save mocks base method.
func (m *MockWfRuntimeMementoStorage) Save(mementoKey, mementoContent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", mementoKey, mementoContent)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockWfRuntimeMementoStorageMockRecorder) Save(mementoKey, mementoContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockWfRuntimeMementoStorage)(nil).Save), mementoKey, mementoContent)
}
