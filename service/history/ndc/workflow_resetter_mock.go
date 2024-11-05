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
// Source: workflow_resetter.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../../../LICENSE -package ndc -source workflow_resetter.go -destination workflow_resetter_mock.go
//

// Package ndc is a generated GoMock package.
package ndc

import (
	context "context"
	reflect "reflect"

	enums "go.temporal.io/api/enums/v1"
	history "go.temporal.io/api/history/v1"
	namespace "go.temporal.io/server/common/namespace"
	gomock "go.uber.org/mock/gomock"
)

// MockWorkflowResetter is a mock of WorkflowResetter interface.
type MockWorkflowResetter struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowResetterMockRecorder
}

// MockWorkflowResetterMockRecorder is the mock recorder for MockWorkflowResetter.
type MockWorkflowResetterMockRecorder struct {
	mock *MockWorkflowResetter
}

// NewMockWorkflowResetter creates a new mock instance.
func NewMockWorkflowResetter(ctrl *gomock.Controller) *MockWorkflowResetter {
	mock := &MockWorkflowResetter{ctrl: ctrl}
	mock.recorder = &MockWorkflowResetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflowResetter) EXPECT() *MockWorkflowResetterMockRecorder {
	return m.recorder
}

// ResetWorkflow mocks base method.
func (m *MockWorkflowResetter) ResetWorkflow(ctx context.Context, namespaceID namespace.ID, workflowID, baseRunID string, baseBranchToken []byte, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID int64, resetRunID, resetRequestID string, baseWorkflow, currentWorkflow Workflow, resetReason string, additionalReapplyEvents []*history.HistoryEvent, resetReapplyExcludeTypes map[enums.ResetReapplyExcludeType]struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflow", ctx, namespaceID, workflowID, baseRunID, baseBranchToken, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID, resetRunID, resetRequestID, baseWorkflow, currentWorkflow, resetReason, additionalReapplyEvents, resetReapplyExcludeTypes)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetWorkflow indicates an expected call of ResetWorkflow.
func (mr *MockWorkflowResetterMockRecorder) ResetWorkflow(ctx, namespaceID, workflowID, baseRunID, baseBranchToken, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID, resetRunID, resetRequestID, baseWorkflow, currentWorkflow, resetReason, additionalReapplyEvents, resetReapplyExcludeTypes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflow", reflect.TypeOf((*MockWorkflowResetter)(nil).ResetWorkflow), ctx, namespaceID, workflowID, baseRunID, baseBranchToken, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID, resetRunID, resetRequestID, baseWorkflow, currentWorkflow, resetReason, additionalReapplyEvents, resetReapplyExcludeTypes)
}
