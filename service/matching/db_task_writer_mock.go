// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: db_task_writer.go

// Package matching is a generated GoMock package.
package matching

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	persistence "go.temporal.io/server/api/persistence/v1"
	future "go.temporal.io/server/common/future"
)

// MockdbTaskWriter is a mock of dbTaskWriter interface.
type MockdbTaskWriter struct {
	ctrl     *gomock.Controller
	recorder *MockdbTaskWriterMockRecorder
}

// MockdbTaskWriterMockRecorder is the mock recorder for MockdbTaskWriter.
type MockdbTaskWriterMockRecorder struct {
	mock *MockdbTaskWriter
}

// NewMockdbTaskWriter creates a new mock instance.
func NewMockdbTaskWriter(ctrl *gomock.Controller) *MockdbTaskWriter {
	mock := &MockdbTaskWriter{ctrl: ctrl}
	mock.recorder = &MockdbTaskWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdbTaskWriter) EXPECT() *MockdbTaskWriterMockRecorder {
	return m.recorder
}

// appendTask mocks base method.
func (m *MockdbTaskWriter) appendTask(task *persistence.TaskInfo) future.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "appendTask", task)
	ret0, _ := ret[0].(future.Future)
	return ret0
}

// appendTask indicates an expected call of appendTask.
func (mr *MockdbTaskWriterMockRecorder) appendTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "appendTask", reflect.TypeOf((*MockdbTaskWriter)(nil).appendTask), task)
}

// flushTasks mocks base method.
func (m *MockdbTaskWriter) flushTasks(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "flushTasks", ctx)
}

// flushTasks indicates an expected call of flushTasks.
func (mr *MockdbTaskWriterMockRecorder) flushTasks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "flushTasks", reflect.TypeOf((*MockdbTaskWriter)(nil).flushTasks), ctx)
}

// notifyFlushChan mocks base method.
func (m *MockdbTaskWriter) notifyFlushChan() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "notifyFlushChan")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// notifyFlushChan indicates an expected call of notifyFlushChan.
func (mr *MockdbTaskWriterMockRecorder) notifyFlushChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "notifyFlushChan", reflect.TypeOf((*MockdbTaskWriter)(nil).notifyFlushChan))
}
