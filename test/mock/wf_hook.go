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
// Source: ./wfengine/wf_hook.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	define "github.com/ApsaraDB/PolarDB-Stack-Workflow/define"
	wfengine "github.com/ApsaraDB/PolarDB-Stack-Workflow/wfengine"
)

// MockWfHook is a mock of WfHook interface.
type MockWfHook struct {
	ctrl     *gomock.Controller
	recorder *MockWfHookMockRecorder
}

// MockWfHookMockRecorder is the mock recorder for MockWfHook.
type MockWfHookMockRecorder struct {
	mock *MockWfHook
}

// NewMockWfHook creates a new mock instance.
func NewMockWfHook(ctrl *gomock.Controller) *MockWfHook {
	mock := &MockWfHook{ctrl: ctrl}
	mock.recorder = &MockWfHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWfHook) EXPECT() *MockWfHookMockRecorder {
	return m.recorder
}

// OnStepCompleted mocks base method.
func (m *MockWfHook) OnStepCompleted(step *wfengine.StepRuntime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnStepCompleted", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnStepCompleted indicates an expected call of OnStepCompleted.
func (mr *MockWfHookMockRecorder) OnStepCompleted(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStepCompleted", reflect.TypeOf((*MockWfHook)(nil).OnStepCompleted), step)
}

// OnStepInit mocks base method.
func (m *MockWfHook) OnStepInit(step *wfengine.StepRuntime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnStepInit", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnStepInit indicates an expected call of OnStepInit.
func (mr *MockWfHookMockRecorder) OnStepInit(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStepInit", reflect.TypeOf((*MockWfHook)(nil).OnStepInit), step)
}

// OnStepWaiting mocks base method.
func (m *MockWfHook) OnStepWaiting(step *wfengine.StepRuntime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnStepWaiting", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnStepWaiting indicates an expected call of OnStepWaiting.
func (mr *MockWfHookMockRecorder) OnStepWaiting(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStepWaiting", reflect.TypeOf((*MockWfHook)(nil).OnStepWaiting), step)
}

// OnWfCompleted mocks base method.
func (m *MockWfHook) OnWfCompleted() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnWfCompleted")
	ret0, _ := ret[0].(error)
	return ret0
}

// OnWfCompleted indicates an expected call of OnWfCompleted.
func (mr *MockWfHookMockRecorder) OnWfCompleted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnWfCompleted", reflect.TypeOf((*MockWfHook)(nil).OnWfCompleted))
}

// OnWfInit mocks base method.
func (m *MockWfHook) OnWfInit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnWfInit")
	ret0, _ := ret[0].(error)
	return ret0
}

// OnWfInit indicates an expected call of OnWfInit.
func (mr *MockWfHookMockRecorder) OnWfInit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnWfInit", reflect.TypeOf((*MockWfHook)(nil).OnWfInit))
}

// OnWfInterrupt mocks base method.
func (m *MockWfHook) OnWfInterrupt(arg0 *define.InterruptError) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnWfInterrupt", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnWfInterrupt indicates an expected call of OnWfInterrupt.
func (mr *MockWfHookMockRecorder) OnWfInterrupt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnWfInterrupt", reflect.TypeOf((*MockWfHook)(nil).OnWfInterrupt), arg0)
}
