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
// Source: ./statemachine/statemachine.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	statemachine "github.com/ApsaraDB/PolarDB-Stack-Workflow/statemachine"
	gomock "github.com/golang/mock/gomock"
)

// MockStateResource is a mock of StateResource interface.
type MockStateResource struct {
	ctrl     *gomock.Controller
	recorder *MockStateResourceMockRecorder
}

// MockStateResourceMockRecorder is the mock recorder for MockStateResource.
type MockStateResourceMockRecorder struct {
	mock *MockStateResource
}

// NewMockStateResource creates a new mock instance.
func NewMockStateResource(ctrl *gomock.Controller) *MockStateResource {
	mock := &MockStateResource{ctrl: ctrl}
	mock.recorder = &MockStateResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateResource) EXPECT() *MockStateResourceMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockStateResource) Fetch() (statemachine.StateResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch")
	ret0, _ := ret[0].(statemachine.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockStateResourceMockRecorder) Fetch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockStateResource)(nil).Fetch))
}

// GetName mocks base method.
func (m *MockStateResource) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockStateResourceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockStateResource)(nil).GetName))
}

// GetNamespace mocks base method.
func (m *MockStateResource) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockStateResourceMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockStateResource)(nil).GetNamespace))
}

// GetState mocks base method.
func (m *MockStateResource) GetState() statemachine.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(statemachine.State)
	return ret0
}

// GetState indicates an expected call of GetState.
func (mr *MockStateResourceMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockStateResource)(nil).GetState))
}

// IsCancelled mocks base method.
func (m *MockStateResource) IsCancelled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCancelled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCancelled indicates an expected call of IsCancelled.
func (mr *MockStateResourceMockRecorder) IsCancelled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCancelled", reflect.TypeOf((*MockStateResource)(nil).IsCancelled))
}

// UpdateState mocks base method.
func (m *MockStateResource) UpdateState(arg0 statemachine.State) (statemachine.StateResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateState", arg0)
	ret0, _ := ret[0].(statemachine.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateState indicates an expected call of UpdateState.
func (mr *MockStateResourceMockRecorder) UpdateState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateState", reflect.TypeOf((*MockStateResource)(nil).UpdateState), arg0)
}

// MockStateMainEnter is a mock of StateMainEnter interface.
type MockStateMainEnter struct {
	ctrl     *gomock.Controller
	recorder *MockStateMainEnterMockRecorder
}

// MockStateMainEnterMockRecorder is the mock recorder for MockStateMainEnter.
type MockStateMainEnterMockRecorder struct {
	mock *MockStateMainEnter
}

// NewMockStateMainEnter creates a new mock instance.
func NewMockStateMainEnter(ctrl *gomock.Controller) *MockStateMainEnter {
	mock := &MockStateMainEnter{ctrl: ctrl}
	mock.recorder = &MockStateMainEnterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateMainEnter) EXPECT() *MockStateMainEnterMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockStateMainEnter) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockStateMainEnterMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockStateMainEnter)(nil).GetName))
}

// MainEnter mocks base method.
func (m *MockStateMainEnter) MainEnter(obj statemachine.StateResource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainEnter", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// MainEnter indicates an expected call of MainEnter.
func (mr *MockStateMainEnterMockRecorder) MainEnter(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainEnter", reflect.TypeOf((*MockStateMainEnter)(nil).MainEnter), obj)
}
