// Code generated by MockGen. DO NOT EDIT.
// Source: api_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
	tfc_api "github.com/zapier/tfbuddy/pkg/tfc_api"
)

// MockApiClient is a mock of ApiClient interface.
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient.
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance.
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// AddTags mocks base method.
func (m *MockApiClient) AddTags(ctx context.Context, workspace, prefix, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTags", ctx, workspace, prefix, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTags indicates an expected call of AddTags.
func (mr *MockApiClientMockRecorder) AddTags(ctx, workspace, prefix, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTags", reflect.TypeOf((*MockApiClient)(nil).AddTags), ctx, workspace, prefix, value)
}

// CreateRunFromSource mocks base method.
func (m *MockApiClient) CreateRunFromSource(ctx context.Context, opts *tfc_api.ApiRunOptions) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRunFromSource", ctx, opts)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRunFromSource indicates an expected call of CreateRunFromSource.
func (mr *MockApiClientMockRecorder) CreateRunFromSource(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRunFromSource", reflect.TypeOf((*MockApiClient)(nil).CreateRunFromSource), ctx, opts)
}

// GetPlanOutput mocks base method.
func (m *MockApiClient) GetPlanOutput(id string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanOutput", id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlanOutput indicates an expected call of GetPlanOutput.
func (mr *MockApiClientMockRecorder) GetPlanOutput(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanOutput", reflect.TypeOf((*MockApiClient)(nil).GetPlanOutput), id)
}

// GetRun mocks base method.
func (m *MockApiClient) GetRun(ctx context.Context, id string) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRun", ctx, id)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun.
func (mr *MockApiClientMockRecorder) GetRun(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockApiClient)(nil).GetRun), ctx, id)
}

// GetTagsByQuery mocks base method.
func (m *MockApiClient) GetTagsByQuery(ctx context.Context, workspace, query string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsByQuery", ctx, workspace, query)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsByQuery indicates an expected call of GetTagsByQuery.
func (mr *MockApiClientMockRecorder) GetTagsByQuery(ctx, workspace, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsByQuery", reflect.TypeOf((*MockApiClient)(nil).GetTagsByQuery), ctx, workspace, query)
}

// GetWorkspaceById mocks base method.
func (m *MockApiClient) GetWorkspaceById(ctx context.Context, id string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceById", ctx, id)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaceById indicates an expected call of GetWorkspaceById.
func (mr *MockApiClientMockRecorder) GetWorkspaceById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceById", reflect.TypeOf((*MockApiClient)(nil).GetWorkspaceById), ctx, id)
}

// GetWorkspaceByName mocks base method.
func (m *MockApiClient) GetWorkspaceByName(ctx context.Context, org, name string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceByName", ctx, org, name)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaceByName indicates an expected call of GetWorkspaceByName.
func (mr *MockApiClientMockRecorder) GetWorkspaceByName(ctx, org, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceByName", reflect.TypeOf((*MockApiClient)(nil).GetWorkspaceByName), ctx, org, name)
}

// LockUnlockWorkspace mocks base method.
func (m *MockApiClient) LockUnlockWorkspace(ctx context.Context, workspace, reason, tag string, lock bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockUnlockWorkspace", ctx, workspace, reason, tag, lock)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockUnlockWorkspace indicates an expected call of LockUnlockWorkspace.
func (mr *MockApiClientMockRecorder) LockUnlockWorkspace(ctx, workspace, reason, tag, lock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockUnlockWorkspace", reflect.TypeOf((*MockApiClient)(nil).LockUnlockWorkspace), ctx, workspace, reason, tag, lock)
}

// RemoveTagsByQuery mocks base method.
func (m *MockApiClient) RemoveTagsByQuery(ctx context.Context, workspace, query string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTagsByQuery", ctx, workspace, query)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTagsByQuery indicates an expected call of RemoveTagsByQuery.
func (mr *MockApiClientMockRecorder) RemoveTagsByQuery(ctx, workspace, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagsByQuery", reflect.TypeOf((*MockApiClient)(nil).RemoveTagsByQuery), ctx, workspace, query)
}
