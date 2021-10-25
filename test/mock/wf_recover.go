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
// Source: ./wfengine/wf_recover.go

// Package mock_wfengine is a generated GoMock package.
package mock

import (
	reflect "reflect"

	define "github.com/ApsaraDB/PolarDB-Stack-Workflow/define"
	statemachine "github.com/ApsaraDB/PolarDB-Stack-Workflow/statemachine"
	gomock "github.com/golang/mock/gomock"
)

// MockRecover is a mock of Recover interface.
type MockRecover struct {
	ctrl     *gomock.Controller
	recorder *MockRecoverMockRecorder
}

// MockRecoverMockRecorder is the mock recorder for MockRecover.
type MockRecoverMockRecorder struct {
	mock *MockRecover
}

// NewMockRecover creates a new mock instance.
func NewMockRecover(ctrl *gomock.Controller) *MockRecover {
	mock := &MockRecover{ctrl: ctrl}
	mock.recorder = &MockRecoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecover) EXPECT() *MockRecoverMockRecorder {
	return m.recorder
}

// CancelResourceRecover mocks base method.
func (m *MockRecover) CancelResourceRecover(resource statemachine.StateResource, conf map[define.WFConfKey]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelResourceRecover", resource, conf)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelResourceRecover indicates an expected call of CancelResourceRecover.
func (mr *MockRecoverMockRecorder) CancelResourceRecover(resource, conf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelResourceRecover", reflect.TypeOf((*MockRecover)(nil).CancelResourceRecover), resource, conf)
}

// GetResourceRecoverInfo mocks base method.
func (m *MockRecover) GetResourceRecoverInfo(resource statemachine.StateResource, conf map[define.WFConfKey]string) (bool, statemachine.State) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceRecoverInfo", resource, conf)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(statemachine.State)
	return ret0, ret1
}

// GetResourceRecoverInfo indicates an expected call of GetResourceRecoverInfo.
func (mr *MockRecoverMockRecorder) GetResourceRecoverInfo(resource, conf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceRecoverInfo", reflect.TypeOf((*MockRecover)(nil).GetResourceRecoverInfo), resource, conf)
}

// SaveResourceInterruptInfo mocks base method.
func (m *MockRecover) SaveResourceInterruptInfo(resource statemachine.StateResource, conf map[define.WFConfKey]string, reason, message, prevState string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResourceInterruptInfo", resource, conf, reason, message, prevState)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveResourceInterruptInfo indicates an expected call of SaveResourceInterruptInfo.
func (mr *MockRecoverMockRecorder) SaveResourceInterruptInfo(resource, conf, reason, message, prevState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResourceInterruptInfo", reflect.TypeOf((*MockRecover)(nil).SaveResourceInterruptInfo), resource, conf, reason, message, prevState)
}
