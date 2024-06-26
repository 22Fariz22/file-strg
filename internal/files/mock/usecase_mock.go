// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/AleksK1NG/api-mc/internal/models"
	utils "github.com/AleksK1NG/api-mc/pkg/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUseCase) Delete(ctx context.Context, file_id *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, file_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(ctx, file_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, file_id)
}

// Download mocks base method.
func (m *MockUseCase) Download(ctx context.Context, file_id *[]byte) (*models.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, file_id)
	ret0, _ := ret[0].(*models.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockUseCaseMockRecorder) Download(ctx, file_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockUseCase)(nil).Download), ctx, file_id)
}

// GetAllFiles mocks base method.
func (m *MockUseCase) GetAllFiles(ctx context.Context, pq *utils.PaginationQuery) (*models.FileList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFiles", ctx, pq)
	ret0, _ := ret[0].(*models.FileList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFiles indicates an expected call of GetAllFiles.
func (mr *MockUseCaseMockRecorder) GetAllFiles(ctx, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFiles", reflect.TypeOf((*MockUseCase)(nil).GetAllFiles), ctx, pq)
}

// Share mocks base method.
func (m *MockUseCase) Share(ctx context.Context, share *models.Share) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Share", ctx, share)
	ret0, _ := ret[0].(error)
	return ret0
}

// Share indicates an expected call of Share.
func (mr *MockUseCaseMockRecorder) Share(ctx, share interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Share", reflect.TypeOf((*MockUseCase)(nil).Share), ctx, share)
}

// Upload mocks base method.
func (m *MockUseCase) Upload(ctx context.Context, filename string, filesize int64, content *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, filename, filesize, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockUseCaseMockRecorder) Upload(ctx, filename, filesize, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockUseCase)(nil).Upload), ctx, filename, filesize, content)
}
