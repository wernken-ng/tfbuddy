// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	runstream "github.com/zapier/tfbuddy/pkg/runstream"
)

// MockStreamClient is a mock of StreamClient interface.
type MockStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockStreamClientMockRecorder
}

// MockStreamClientMockRecorder is the mock recorder for MockStreamClient.
type MockStreamClientMockRecorder struct {
	mock *MockStreamClient
}

// NewMockStreamClient creates a new mock instance.
func NewMockStreamClient(ctrl *gomock.Controller) *MockStreamClient {
	mock := &MockStreamClient{ctrl: ctrl}
	mock.recorder = &MockStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamClient) EXPECT() *MockStreamClientMockRecorder {
	return m.recorder
}

// AddRunMeta mocks base method.
func (m *MockStreamClient) AddRunMeta(rmd runstream.RunMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRunMeta", rmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRunMeta indicates an expected call of AddRunMeta.
func (mr *MockStreamClientMockRecorder) AddRunMeta(rmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRunMeta", reflect.TypeOf((*MockStreamClient)(nil).AddRunMeta), rmd)
}

// GetRunMeta mocks base method.
func (m *MockStreamClient) GetRunMeta(runID string) (runstream.RunMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunMeta", runID)
	ret0, _ := ret[0].(runstream.RunMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunMeta indicates an expected call of GetRunMeta.
func (mr *MockStreamClientMockRecorder) GetRunMeta(runID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunMeta", reflect.TypeOf((*MockStreamClient)(nil).GetRunMeta), runID)
}

// HealthCheck mocks base method.
func (m *MockStreamClient) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockStreamClientMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockStreamClient)(nil).HealthCheck))
}

// NewTFRunPollingTask mocks base method.
func (m *MockStreamClient) NewTFRunPollingTask(meta runstream.RunMetadata, delay time.Duration) runstream.RunPollingTask {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTFRunPollingTask", meta, delay)
	ret0, _ := ret[0].(runstream.RunPollingTask)
	return ret0
}

// NewTFRunPollingTask indicates an expected call of NewTFRunPollingTask.
func (mr *MockStreamClientMockRecorder) NewTFRunPollingTask(meta, delay interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTFRunPollingTask", reflect.TypeOf((*MockStreamClient)(nil).NewTFRunPollingTask), meta, delay)
}

// PublishTFRunEvent mocks base method.
func (m *MockStreamClient) PublishTFRunEvent(ctx context.Context, re runstream.RunEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishTFRunEvent", ctx, re)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishTFRunEvent indicates an expected call of PublishTFRunEvent.
func (mr *MockStreamClientMockRecorder) PublishTFRunEvent(ctx, re interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTFRunEvent", reflect.TypeOf((*MockStreamClient)(nil).PublishTFRunEvent), ctx, re)
}

// SubscribeTFRunEvents mocks base method.
func (m *MockStreamClient) SubscribeTFRunEvents(queue string, cb func(runstream.RunEvent) bool) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeTFRunEvents", queue, cb)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeTFRunEvents indicates an expected call of SubscribeTFRunEvents.
func (mr *MockStreamClientMockRecorder) SubscribeTFRunEvents(queue, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeTFRunEvents", reflect.TypeOf((*MockStreamClient)(nil).SubscribeTFRunEvents), queue, cb)
}

// SubscribeTFRunPollingTasks mocks base method.
func (m *MockStreamClient) SubscribeTFRunPollingTasks(cb func(runstream.RunPollingTask) bool) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeTFRunPollingTasks", cb)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeTFRunPollingTasks indicates an expected call of SubscribeTFRunPollingTasks.
func (mr *MockStreamClientMockRecorder) SubscribeTFRunPollingTasks(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeTFRunPollingTasks", reflect.TypeOf((*MockStreamClient)(nil).SubscribeTFRunPollingTasks), cb)
}

// MockRunEvent is a mock of RunEvent interface.
type MockRunEvent struct {
	ctrl     *gomock.Controller
	recorder *MockRunEventMockRecorder
}

// MockRunEventMockRecorder is the mock recorder for MockRunEvent.
type MockRunEventMockRecorder struct {
	mock *MockRunEvent
}

// NewMockRunEvent creates a new mock instance.
func NewMockRunEvent(ctrl *gomock.Controller) *MockRunEvent {
	mock := &MockRunEvent{ctrl: ctrl}
	mock.recorder = &MockRunEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunEvent) EXPECT() *MockRunEventMockRecorder {
	return m.recorder
}

// GetContext mocks base method.
func (m *MockRunEvent) GetContext() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockRunEventMockRecorder) GetContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockRunEvent)(nil).GetContext))
}

// GetMetadata mocks base method.
func (m *MockRunEvent) GetMetadata() runstream.RunMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata")
	ret0, _ := ret[0].(runstream.RunMetadata)
	return ret0
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockRunEventMockRecorder) GetMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockRunEvent)(nil).GetMetadata))
}

// GetNewStatus mocks base method.
func (m *MockRunEvent) GetNewStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNewStatus indicates an expected call of GetNewStatus.
func (mr *MockRunEventMockRecorder) GetNewStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewStatus", reflect.TypeOf((*MockRunEvent)(nil).GetNewStatus))
}

// GetRunID mocks base method.
func (m *MockRunEvent) GetRunID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunID indicates an expected call of GetRunID.
func (mr *MockRunEventMockRecorder) GetRunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunID", reflect.TypeOf((*MockRunEvent)(nil).GetRunID))
}

// SetCarrier mocks base method.
func (m *MockRunEvent) SetCarrier(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCarrier", arg0)
}

// SetCarrier indicates an expected call of SetCarrier.
func (mr *MockRunEventMockRecorder) SetCarrier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCarrier", reflect.TypeOf((*MockRunEvent)(nil).SetCarrier), arg0)
}

// SetContext mocks base method.
func (m *MockRunEvent) SetContext(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContext", arg0)
}

// SetContext indicates an expected call of SetContext.
func (mr *MockRunEventMockRecorder) SetContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContext", reflect.TypeOf((*MockRunEvent)(nil).SetContext), arg0)
}

// SetMetadata mocks base method.
func (m *MockRunEvent) SetMetadata(arg0 runstream.RunMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMetadata", arg0)
}

// SetMetadata indicates an expected call of SetMetadata.
func (mr *MockRunEventMockRecorder) SetMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetadata", reflect.TypeOf((*MockRunEvent)(nil).SetMetadata), arg0)
}

// MockRunMetadata is a mock of RunMetadata interface.
type MockRunMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockRunMetadataMockRecorder
}

// MockRunMetadataMockRecorder is the mock recorder for MockRunMetadata.
type MockRunMetadataMockRecorder struct {
	mock *MockRunMetadata
}

// NewMockRunMetadata creates a new mock instance.
func NewMockRunMetadata(ctrl *gomock.Controller) *MockRunMetadata {
	mock := &MockRunMetadata{ctrl: ctrl}
	mock.recorder = &MockRunMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunMetadata) EXPECT() *MockRunMetadataMockRecorder {
	return m.recorder
}

// GetAction mocks base method.
func (m *MockRunMetadata) GetAction() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAction")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAction indicates an expected call of GetAction.
func (mr *MockRunMetadataMockRecorder) GetAction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAction", reflect.TypeOf((*MockRunMetadata)(nil).GetAction))
}

// GetCommitSHA mocks base method.
func (m *MockRunMetadata) GetCommitSHA() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitSHA")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCommitSHA indicates an expected call of GetCommitSHA.
func (mr *MockRunMetadataMockRecorder) GetCommitSHA() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitSHA", reflect.TypeOf((*MockRunMetadata)(nil).GetCommitSHA))
}

// GetDiscussionID mocks base method.
func (m *MockRunMetadata) GetDiscussionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiscussionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDiscussionID indicates an expected call of GetDiscussionID.
func (mr *MockRunMetadataMockRecorder) GetDiscussionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiscussionID", reflect.TypeOf((*MockRunMetadata)(nil).GetDiscussionID))
}

// GetMRInternalID mocks base method.
func (m *MockRunMetadata) GetMRInternalID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMRInternalID")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMRInternalID indicates an expected call of GetMRInternalID.
func (mr *MockRunMetadataMockRecorder) GetMRInternalID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMRInternalID", reflect.TypeOf((*MockRunMetadata)(nil).GetMRInternalID))
}

// GetMRProjectNameWithNamespace mocks base method.
func (m *MockRunMetadata) GetMRProjectNameWithNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMRProjectNameWithNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMRProjectNameWithNamespace indicates an expected call of GetMRProjectNameWithNamespace.
func (mr *MockRunMetadataMockRecorder) GetMRProjectNameWithNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMRProjectNameWithNamespace", reflect.TypeOf((*MockRunMetadata)(nil).GetMRProjectNameWithNamespace))
}

// GetOrganization mocks base method.
func (m *MockRunMetadata) GetOrganization() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganization")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrganization indicates an expected call of GetOrganization.
func (mr *MockRunMetadataMockRecorder) GetOrganization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganization", reflect.TypeOf((*MockRunMetadata)(nil).GetOrganization))
}

// GetRootNoteID mocks base method.
func (m *MockRunMetadata) GetRootNoteID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRootNoteID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetRootNoteID indicates an expected call of GetRootNoteID.
func (mr *MockRunMetadataMockRecorder) GetRootNoteID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRootNoteID", reflect.TypeOf((*MockRunMetadata)(nil).GetRootNoteID))
}

// GetRunID mocks base method.
func (m *MockRunMetadata) GetRunID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunID indicates an expected call of GetRunID.
func (mr *MockRunMetadataMockRecorder) GetRunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunID", reflect.TypeOf((*MockRunMetadata)(nil).GetRunID))
}

// GetVcsProvider mocks base method.
func (m *MockRunMetadata) GetVcsProvider() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVcsProvider")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVcsProvider indicates an expected call of GetVcsProvider.
func (mr *MockRunMetadataMockRecorder) GetVcsProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVcsProvider", reflect.TypeOf((*MockRunMetadata)(nil).GetVcsProvider))
}

// GetWorkspace mocks base method.
func (m *MockRunMetadata) GetWorkspace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockRunMetadataMockRecorder) GetWorkspace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockRunMetadata)(nil).GetWorkspace))
}

// MockRunPollingTask is a mock of RunPollingTask interface.
type MockRunPollingTask struct {
	ctrl     *gomock.Controller
	recorder *MockRunPollingTaskMockRecorder
}

// MockRunPollingTaskMockRecorder is the mock recorder for MockRunPollingTask.
type MockRunPollingTaskMockRecorder struct {
	mock *MockRunPollingTask
}

// NewMockRunPollingTask creates a new mock instance.
func NewMockRunPollingTask(ctrl *gomock.Controller) *MockRunPollingTask {
	mock := &MockRunPollingTask{ctrl: ctrl}
	mock.recorder = &MockRunPollingTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunPollingTask) EXPECT() *MockRunPollingTaskMockRecorder {
	return m.recorder
}

// Completed mocks base method.
func (m *MockRunPollingTask) Completed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Completed")
	ret0, _ := ret[0].(error)
	return ret0
}

// Completed indicates an expected call of Completed.
func (mr *MockRunPollingTaskMockRecorder) Completed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Completed", reflect.TypeOf((*MockRunPollingTask)(nil).Completed))
}

// GetContext mocks base method.
func (m *MockRunPollingTask) GetContext() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockRunPollingTaskMockRecorder) GetContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockRunPollingTask)(nil).GetContext))
}

// GetLastStatus mocks base method.
func (m *MockRunPollingTask) GetLastStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLastStatus indicates an expected call of GetLastStatus.
func (mr *MockRunPollingTaskMockRecorder) GetLastStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastStatus", reflect.TypeOf((*MockRunPollingTask)(nil).GetLastStatus))
}

// GetRunID mocks base method.
func (m *MockRunPollingTask) GetRunID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunID indicates an expected call of GetRunID.
func (mr *MockRunPollingTaskMockRecorder) GetRunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunID", reflect.TypeOf((*MockRunPollingTask)(nil).GetRunID))
}

// GetRunMetaData mocks base method.
func (m *MockRunPollingTask) GetRunMetaData() runstream.RunMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunMetaData")
	ret0, _ := ret[0].(runstream.RunMetadata)
	return ret0
}

// GetRunMetaData indicates an expected call of GetRunMetaData.
func (mr *MockRunPollingTaskMockRecorder) GetRunMetaData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunMetaData", reflect.TypeOf((*MockRunPollingTask)(nil).GetRunMetaData))
}

// Reschedule mocks base method.
func (m *MockRunPollingTask) Reschedule(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reschedule", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reschedule indicates an expected call of Reschedule.
func (mr *MockRunPollingTaskMockRecorder) Reschedule(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reschedule", reflect.TypeOf((*MockRunPollingTask)(nil).Reschedule), ctx)
}

// Schedule mocks base method.
func (m *MockRunPollingTask) Schedule(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Schedule indicates an expected call of Schedule.
func (mr *MockRunPollingTaskMockRecorder) Schedule(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockRunPollingTask)(nil).Schedule), ctx)
}

// SetCarrier mocks base method.
func (m *MockRunPollingTask) SetCarrier(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCarrier", arg0)
}

// SetCarrier indicates an expected call of SetCarrier.
func (mr *MockRunPollingTaskMockRecorder) SetCarrier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCarrier", reflect.TypeOf((*MockRunPollingTask)(nil).SetCarrier), arg0)
}

// SetLastStatus mocks base method.
func (m *MockRunPollingTask) SetLastStatus(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastStatus", arg0)
}

// SetLastStatus indicates an expected call of SetLastStatus.
func (mr *MockRunPollingTaskMockRecorder) SetLastStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastStatus", reflect.TypeOf((*MockRunPollingTask)(nil).SetLastStatus), arg0)
}
