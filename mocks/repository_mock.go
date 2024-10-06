// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/repository.go -destination=mocks/repository_mock.go -package=mocks StorageRepository,QueueRepository,TodoRepository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	service_models "codechallenge/internal/service/service_models"
	utils "codechallenge/utils"
	context "context"
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStorageRepository is a mock of StorageRepository interface.
type MockStorageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStorageRepositoryMockRecorder
}

// MockStorageRepositoryMockRecorder is the mock recorder for MockStorageRepository.
type MockStorageRepositoryMockRecorder struct {
	mock *MockStorageRepository
}

// NewMockStorageRepository creates a new mock instance.
func NewMockStorageRepository(ctrl *gomock.Controller) *MockStorageRepository {
	mock := &MockStorageRepository{ctrl: ctrl}
	mock.recorder = &MockStorageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageRepository) EXPECT() *MockStorageRepositoryMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *MockStorageRepository) Upload(ctx context.Context, file io.ReadCloser, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, file, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockStorageRepositoryMockRecorder) Upload(ctx, file, filename any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockStorageRepository)(nil).Upload), ctx, file, filename)
}

// MockQueueRepository is a mock of QueueRepository interface.
type MockQueueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQueueRepositoryMockRecorder
}

// MockQueueRepositoryMockRecorder is the mock recorder for MockQueueRepository.
type MockQueueRepositoryMockRecorder struct {
	mock *MockQueueRepository
}

// NewMockQueueRepository creates a new mock instance.
func NewMockQueueRepository(ctrl *gomock.Controller) *MockQueueRepository {
	mock := &MockQueueRepository{ctrl: ctrl}
	mock.recorder = &MockQueueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueRepository) EXPECT() *MockQueueRepositoryMockRecorder {
	return m.recorder
}

// PushTodoItem mocks base method.
func (m *MockQueueRepository) PushTodoItem(ctx context.Context, message service_models.TodoItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushTodoItem", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushTodoItem indicates an expected call of PushTodoItem.
func (mr *MockQueueRepositoryMockRecorder) PushTodoItem(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushTodoItem", reflect.TypeOf((*MockQueueRepository)(nil).PushTodoItem), ctx, message)
}

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// CreateWithTX mocks base method.
func (m *MockTodoRepository) CreateWithTX(ctx context.Context, todoItem service_models.TodoItem) (utils.DbTransaction, service_models.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithTX", ctx, todoItem)
	ret0, _ := ret[0].(utils.DbTransaction)
	ret1, _ := ret[1].(service_models.TodoItem)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateWithTX indicates an expected call of CreateWithTX.
func (mr *MockTodoRepositoryMockRecorder) CreateWithTX(ctx, todoItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithTX", reflect.TypeOf((*MockTodoRepository)(nil).CreateWithTX), ctx, todoItem)
}

// Get mocks base method.
func (m *MockTodoRepository) Get(ctx context.Context, id string) (service_models.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(service_models.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTodoRepositoryMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodoRepository)(nil).Get), ctx, id)
}
