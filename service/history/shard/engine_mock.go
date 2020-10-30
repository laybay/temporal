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
// Source: engine.go

// Package shard is a generated GoMock package.
package shard

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "go.temporal.io/api/common/v1"
	historypb "go.temporal.io/api/history/v1"

	historyservice "go.temporal.io/server/api/historyservice/v1"
	repication "go.temporal.io/server/api/replication/v1"
	persistence "go.temporal.io/server/common/persistence"
	events "go.temporal.io/server/service/history/events"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockEngine) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockEngineMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start))
}

// Stop mocks base method.
func (m *MockEngine) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockEngineMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockEngine)(nil).Stop))
}

// StartWorkflowExecution mocks base method.
func (m *MockEngine) StartWorkflowExecution(ctx context.Context, request *historyservice.StartWorkflowExecutionRequest) (*historyservice.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*historyservice.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflowExecution indicates an expected call of StartWorkflowExecution.
func (mr *MockEngineMockRecorder) StartWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).StartWorkflowExecution), ctx, request)
}

// GetMutableState mocks base method.
func (m *MockEngine) GetMutableState(ctx context.Context, request *historyservice.GetMutableStateRequest) (*historyservice.GetMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMutableState", ctx, request)
	ret0, _ := ret[0].(*historyservice.GetMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMutableState indicates an expected call of GetMutableState.
func (mr *MockEngineMockRecorder) GetMutableState(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMutableState", reflect.TypeOf((*MockEngine)(nil).GetMutableState), ctx, request)
}

// PollMutableState mocks base method.
func (m *MockEngine) PollMutableState(ctx context.Context, request *historyservice.PollMutableStateRequest) (*historyservice.PollMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollMutableState", ctx, request)
	ret0, _ := ret[0].(*historyservice.PollMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollMutableState indicates an expected call of PollMutableState.
func (mr *MockEngineMockRecorder) PollMutableState(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollMutableState", reflect.TypeOf((*MockEngine)(nil).PollMutableState), ctx, request)
}

// DescribeMutableState mocks base method.
func (m *MockEngine) DescribeMutableState(ctx context.Context, request *historyservice.DescribeMutableStateRequest) (*historyservice.DescribeMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMutableState", ctx, request)
	ret0, _ := ret[0].(*historyservice.DescribeMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMutableState indicates an expected call of DescribeMutableState.
func (mr *MockEngineMockRecorder) DescribeMutableState(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMutableState", reflect.TypeOf((*MockEngine)(nil).DescribeMutableState), ctx, request)
}

// ResetStickyTaskQueue mocks base method.
func (m *MockEngine) ResetStickyTaskQueue(ctx context.Context, resetRequest *historyservice.ResetStickyTaskQueueRequest) (*historyservice.ResetStickyTaskQueueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetStickyTaskQueue", ctx, resetRequest)
	ret0, _ := ret[0].(*historyservice.ResetStickyTaskQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetStickyTaskQueue indicates an expected call of ResetStickyTaskQueue.
func (mr *MockEngineMockRecorder) ResetStickyTaskQueue(ctx, resetRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStickyTaskQueue", reflect.TypeOf((*MockEngine)(nil).ResetStickyTaskQueue), ctx, resetRequest)
}

// DescribeWorkflowExecution mocks base method.
func (m *MockEngine) DescribeWorkflowExecution(ctx context.Context, request *historyservice.DescribeWorkflowExecutionRequest) (*historyservice.DescribeWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*historyservice.DescribeWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkflowExecution indicates an expected call of DescribeWorkflowExecution.
func (mr *MockEngineMockRecorder) DescribeWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).DescribeWorkflowExecution), ctx, request)
}

// RecordWorkflowTaskStarted mocks base method.
func (m *MockEngine) RecordWorkflowTaskStarted(ctx context.Context, request *historyservice.RecordWorkflowTaskStartedRequest) (*historyservice.RecordWorkflowTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowTaskStarted", ctx, request)
	ret0, _ := ret[0].(*historyservice.RecordWorkflowTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordWorkflowTaskStarted indicates an expected call of RecordWorkflowTaskStarted.
func (mr *MockEngineMockRecorder) RecordWorkflowTaskStarted(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowTaskStarted", reflect.TypeOf((*MockEngine)(nil).RecordWorkflowTaskStarted), ctx, request)
}

// RecordActivityTaskStarted mocks base method.
func (m *MockEngine) RecordActivityTaskStarted(ctx context.Context, request *historyservice.RecordActivityTaskStartedRequest) (*historyservice.RecordActivityTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskStarted", ctx, request)
	ret0, _ := ret[0].(*historyservice.RecordActivityTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskStarted indicates an expected call of RecordActivityTaskStarted.
func (mr *MockEngineMockRecorder) RecordActivityTaskStarted(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskStarted", reflect.TypeOf((*MockEngine)(nil).RecordActivityTaskStarted), ctx, request)
}

// RespondWorkflowTaskCompleted mocks base method.
func (m *MockEngine) RespondWorkflowTaskCompleted(ctx context.Context, request *historyservice.RespondWorkflowTaskCompletedRequest) (*historyservice.RespondWorkflowTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondWorkflowTaskCompleted", ctx, request)
	ret0, _ := ret[0].(*historyservice.RespondWorkflowTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondWorkflowTaskCompleted indicates an expected call of RespondWorkflowTaskCompleted.
func (mr *MockEngineMockRecorder) RespondWorkflowTaskCompleted(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondWorkflowTaskCompleted", reflect.TypeOf((*MockEngine)(nil).RespondWorkflowTaskCompleted), ctx, request)
}

// RespondWorkflowTaskFailed mocks base method.
func (m *MockEngine) RespondWorkflowTaskFailed(ctx context.Context, request *historyservice.RespondWorkflowTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondWorkflowTaskFailed", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondWorkflowTaskFailed indicates an expected call of RespondWorkflowTaskFailed.
func (mr *MockEngineMockRecorder) RespondWorkflowTaskFailed(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondWorkflowTaskFailed", reflect.TypeOf((*MockEngine)(nil).RespondWorkflowTaskFailed), ctx, request)
}

// RespondActivityTaskCompleted mocks base method.
func (m *MockEngine) RespondActivityTaskCompleted(ctx context.Context, request *historyservice.RespondActivityTaskCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCompleted", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCompleted indicates an expected call of RespondActivityTaskCompleted.
func (mr *MockEngineMockRecorder) RespondActivityTaskCompleted(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCompleted", reflect.TypeOf((*MockEngine)(nil).RespondActivityTaskCompleted), ctx, request)
}

// RespondActivityTaskFailed mocks base method.
func (m *MockEngine) RespondActivityTaskFailed(ctx context.Context, request *historyservice.RespondActivityTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskFailed", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskFailed indicates an expected call of RespondActivityTaskFailed.
func (mr *MockEngineMockRecorder) RespondActivityTaskFailed(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskFailed", reflect.TypeOf((*MockEngine)(nil).RespondActivityTaskFailed), ctx, request)
}

// RespondActivityTaskCanceled mocks base method.
func (m *MockEngine) RespondActivityTaskCanceled(ctx context.Context, request *historyservice.RespondActivityTaskCanceledRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceled", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCanceled indicates an expected call of RespondActivityTaskCanceled.
func (mr *MockEngineMockRecorder) RespondActivityTaskCanceled(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCanceled", reflect.TypeOf((*MockEngine)(nil).RespondActivityTaskCanceled), ctx, request)
}

// RecordActivityTaskHeartbeat mocks base method.
func (m *MockEngine) RecordActivityTaskHeartbeat(ctx context.Context, request *historyservice.RecordActivityTaskHeartbeatRequest) (*historyservice.RecordActivityTaskHeartbeatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeat", ctx, request)
	ret0, _ := ret[0].(*historyservice.RecordActivityTaskHeartbeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskHeartbeat indicates an expected call of RecordActivityTaskHeartbeat.
func (mr *MockEngineMockRecorder) RecordActivityTaskHeartbeat(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskHeartbeat", reflect.TypeOf((*MockEngine)(nil).RecordActivityTaskHeartbeat), ctx, request)
}

// RequestCancelWorkflowExecution mocks base method.
func (m *MockEngine) RequestCancelWorkflowExecution(ctx context.Context, request *historyservice.RequestCancelWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCancelWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestCancelWorkflowExecution indicates an expected call of RequestCancelWorkflowExecution.
func (mr *MockEngineMockRecorder) RequestCancelWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCancelWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).RequestCancelWorkflowExecution), ctx, request)
}

// SignalWorkflowExecution mocks base method.
func (m *MockEngine) SignalWorkflowExecution(ctx context.Context, request *historyservice.SignalWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignalWorkflowExecution indicates an expected call of SignalWorkflowExecution.
func (mr *MockEngineMockRecorder) SignalWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).SignalWorkflowExecution), ctx, request)
}

// SignalWithStartWorkflowExecution mocks base method.
func (m *MockEngine) SignalWithStartWorkflowExecution(ctx context.Context, request *historyservice.SignalWithStartWorkflowExecutionRequest) (*historyservice.SignalWithStartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*historyservice.SignalWithStartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignalWithStartWorkflowExecution indicates an expected call of SignalWithStartWorkflowExecution.
func (mr *MockEngineMockRecorder) SignalWithStartWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWithStartWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).SignalWithStartWorkflowExecution), ctx, request)
}

// RemoveSignalMutableState mocks base method.
func (m *MockEngine) RemoveSignalMutableState(ctx context.Context, request *historyservice.RemoveSignalMutableStateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSignalMutableState", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSignalMutableState indicates an expected call of RemoveSignalMutableState.
func (mr *MockEngineMockRecorder) RemoveSignalMutableState(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSignalMutableState", reflect.TypeOf((*MockEngine)(nil).RemoveSignalMutableState), ctx, request)
}

// TerminateWorkflowExecution mocks base method.
func (m *MockEngine) TerminateWorkflowExecution(ctx context.Context, request *historyservice.TerminateWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateWorkflowExecution indicates an expected call of TerminateWorkflowExecution.
func (mr *MockEngineMockRecorder) TerminateWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).TerminateWorkflowExecution), ctx, request)
}

// ResetWorkflowExecution mocks base method.
func (m *MockEngine) ResetWorkflowExecution(ctx context.Context, request *historyservice.ResetWorkflowExecutionRequest) (*historyservice.ResetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*historyservice.ResetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution.
func (mr *MockEngineMockRecorder) ResetWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockEngine)(nil).ResetWorkflowExecution), ctx, request)
}

// ScheduleWorkflowTask mocks base method.
func (m *MockEngine) ScheduleWorkflowTask(ctx context.Context, request *historyservice.ScheduleWorkflowTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleWorkflowTask", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleWorkflowTask indicates an expected call of ScheduleWorkflowTask.
func (mr *MockEngineMockRecorder) ScheduleWorkflowTask(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleWorkflowTask", reflect.TypeOf((*MockEngine)(nil).ScheduleWorkflowTask), ctx, request)
}

// RecordChildExecutionCompleted mocks base method.
func (m *MockEngine) RecordChildExecutionCompleted(ctx context.Context, request *historyservice.RecordChildExecutionCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordChildExecutionCompleted", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordChildExecutionCompleted indicates an expected call of RecordChildExecutionCompleted.
func (mr *MockEngineMockRecorder) RecordChildExecutionCompleted(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordChildExecutionCompleted", reflect.TypeOf((*MockEngine)(nil).RecordChildExecutionCompleted), ctx, request)
}

// ReplicateEventsV2 mocks base method.
func (m *MockEngine) ReplicateEventsV2(ctx context.Context, request *historyservice.ReplicateEventsV2Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateEventsV2", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateEventsV2 indicates an expected call of ReplicateEventsV2.
func (mr *MockEngineMockRecorder) ReplicateEventsV2(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateEventsV2", reflect.TypeOf((*MockEngine)(nil).ReplicateEventsV2), ctx, request)
}

// SyncShardStatus mocks base method.
func (m *MockEngine) SyncShardStatus(ctx context.Context, request *historyservice.SyncShardStatusRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncShardStatus", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncShardStatus indicates an expected call of SyncShardStatus.
func (mr *MockEngineMockRecorder) SyncShardStatus(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncShardStatus", reflect.TypeOf((*MockEngine)(nil).SyncShardStatus), ctx, request)
}

// SyncActivity mocks base method.
func (m *MockEngine) SyncActivity(ctx context.Context, request *historyservice.SyncActivityRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncActivity", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncActivity indicates an expected call of SyncActivity.
func (mr *MockEngineMockRecorder) SyncActivity(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncActivity", reflect.TypeOf((*MockEngine)(nil).SyncActivity), ctx, request)
}

// GetReplicationMessages mocks base method.
func (m *MockEngine) GetReplicationMessages(ctx context.Context, pollingCluster string, lastReadMessageID int64) (*repication.ReplicationMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicationMessages", ctx, pollingCluster, lastReadMessageID)
	ret0, _ := ret[0].(*repication.ReplicationMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicationMessages indicates an expected call of GetReplicationMessages.
func (mr *MockEngineMockRecorder) GetReplicationMessages(ctx, pollingCluster, lastReadMessageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationMessages", reflect.TypeOf((*MockEngine)(nil).GetReplicationMessages), ctx, pollingCluster, lastReadMessageID)
}

// GetDLQReplicationMessages mocks base method.
func (m *MockEngine) GetDLQReplicationMessages(ctx context.Context, taskInfos []*repication.ReplicationTaskInfo) ([]*repication.ReplicationTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDLQReplicationMessages", ctx, taskInfos)
	ret0, _ := ret[0].([]*repication.ReplicationTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLQReplicationMessages indicates an expected call of GetDLQReplicationMessages.
func (mr *MockEngineMockRecorder) GetDLQReplicationMessages(ctx, taskInfos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLQReplicationMessages", reflect.TypeOf((*MockEngine)(nil).GetDLQReplicationMessages), ctx, taskInfos)
}

// QueryWorkflow mocks base method.
func (m *MockEngine) QueryWorkflow(ctx context.Context, request *historyservice.QueryWorkflowRequest) (*historyservice.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryWorkflow", ctx, request)
	ret0, _ := ret[0].(*historyservice.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow.
func (mr *MockEngineMockRecorder) QueryWorkflow(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockEngine)(nil).QueryWorkflow), ctx, request)
}

// ReapplyEvents mocks base method.
func (m *MockEngine) ReapplyEvents(ctx context.Context, namespaceUUID, workflowID, runID string, events []*historypb.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReapplyEvents", ctx, namespaceUUID, workflowID, runID, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReapplyEvents indicates an expected call of ReapplyEvents.
func (mr *MockEngineMockRecorder) ReapplyEvents(ctx, namespaceUUID, workflowID, runID, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReapplyEvents", reflect.TypeOf((*MockEngine)(nil).ReapplyEvents), ctx, namespaceUUID, workflowID, runID, events)
}

// GetDLQMessages mocks base method.
func (m *MockEngine) GetDLQMessages(ctx context.Context, messagesRequest *historyservice.GetDLQMessagesRequest) (*historyservice.GetDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDLQMessages", ctx, messagesRequest)
	ret0, _ := ret[0].(*historyservice.GetDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLQMessages indicates an expected call of GetDLQMessages.
func (mr *MockEngineMockRecorder) GetDLQMessages(ctx, messagesRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLQMessages", reflect.TypeOf((*MockEngine)(nil).GetDLQMessages), ctx, messagesRequest)
}

// PurgeDLQMessages mocks base method.
func (m *MockEngine) PurgeDLQMessages(ctx context.Context, messagesRequest *historyservice.PurgeDLQMessagesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeDLQMessages", ctx, messagesRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeDLQMessages indicates an expected call of PurgeDLQMessages.
func (mr *MockEngineMockRecorder) PurgeDLQMessages(ctx, messagesRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeDLQMessages", reflect.TypeOf((*MockEngine)(nil).PurgeDLQMessages), ctx, messagesRequest)
}

// MergeDLQMessages mocks base method.
func (m *MockEngine) MergeDLQMessages(ctx context.Context, messagesRequest *historyservice.MergeDLQMessagesRequest) (*historyservice.MergeDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeDLQMessages", ctx, messagesRequest)
	ret0, _ := ret[0].(*historyservice.MergeDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeDLQMessages indicates an expected call of MergeDLQMessages.
func (mr *MockEngineMockRecorder) MergeDLQMessages(ctx, messagesRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeDLQMessages", reflect.TypeOf((*MockEngine)(nil).MergeDLQMessages), ctx, messagesRequest)
}

// RefreshWorkflowTasks mocks base method.
func (m *MockEngine) RefreshWorkflowTasks(ctx context.Context, namespaceUUID string, execution common.WorkflowExecution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWorkflowTasks", ctx, namespaceUUID, execution)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWorkflowTasks indicates an expected call of RefreshWorkflowTasks.
func (mr *MockEngineMockRecorder) RefreshWorkflowTasks(ctx, namespaceUUID, execution interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWorkflowTasks", reflect.TypeOf((*MockEngine)(nil).RefreshWorkflowTasks), ctx, namespaceUUID, execution)
}

// NotifyNewHistoryEvent mocks base method.
func (m *MockEngine) NotifyNewHistoryEvent(event *events.Notification) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewHistoryEvent", event)
}

// NotifyNewHistoryEvent indicates an expected call of NotifyNewHistoryEvent.
func (mr *MockEngineMockRecorder) NotifyNewHistoryEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewHistoryEvent", reflect.TypeOf((*MockEngine)(nil).NotifyNewHistoryEvent), event)
}

// NotifyNewTransferTasks mocks base method.
func (m *MockEngine) NotifyNewTransferTasks(tasks []persistence.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTransferTasks", tasks)
}

// NotifyNewTransferTasks indicates an expected call of NotifyNewTransferTasks.
func (mr *MockEngineMockRecorder) NotifyNewTransferTasks(tasks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTransferTasks", reflect.TypeOf((*MockEngine)(nil).NotifyNewTransferTasks), tasks)
}

// NotifyNewReplicationTasks mocks base method.
func (m *MockEngine) NotifyNewReplicationTasks(tasks []persistence.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewReplicationTasks", tasks)
}

// NotifyNewReplicationTasks indicates an expected call of NotifyNewReplicationTasks.
func (mr *MockEngineMockRecorder) NotifyNewReplicationTasks(tasks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewReplicationTasks", reflect.TypeOf((*MockEngine)(nil).NotifyNewReplicationTasks), tasks)
}

// NotifyNewTimerTasks mocks base method.
func (m *MockEngine) NotifyNewTimerTasks(tasks []persistence.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTimerTasks", tasks)
}

// NotifyNewTimerTasks indicates an expected call of NotifyNewTimerTasks.
func (mr *MockEngineMockRecorder) NotifyNewTimerTasks(tasks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTimerTasks", reflect.TypeOf((*MockEngine)(nil).NotifyNewTimerTasks), tasks)
}
