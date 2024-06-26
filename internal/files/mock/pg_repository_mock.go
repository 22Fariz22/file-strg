// Code generated by MockGen. DO NOT EDIT.
// Source: pg_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/AleksK1NG/api-mc/internal/models"
	utils "github.com/AleksK1NG/api-mc/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, user_id, file uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, user_id, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, user_id, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, user_id, file)
}

// Download mocks base method.
func (m *MockRepository) Download(ctx context.Context, file *models.File) (*models.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, file)
	ret0, _ := ret[0].(*models.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockRepositoryMockRecorder) Download(ctx, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockRepository)(nil).Download), ctx, file)
}

// GetAllFiles mocks base method.
func (m *MockRepository) GetAllFiles(ctx context.Context, user *models.User, pq *utils.PaginationQuery) (*models.FileList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFiles", ctx, user, pq)
	ret0, _ := ret[0].(*models.FileList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFiles indicates an expected call of GetAllFiles.
func (mr *MockRepositoryMockRecorder) GetAllFiles(ctx, user, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFiles", reflect.TypeOf((*MockRepository)(nil).GetAllFiles), ctx, user, pq)
}

// Share mocks base method.
func (m *MockRepository) Share(ctx context.Context, share *models.Share) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Share", ctx, share)
	ret0, _ := ret[0].(error)
	return ret0
}

// Share indicates an expected call of Share.
func (mr *MockRepositoryMockRecorder) Share(ctx, share interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Share", reflect.TypeOf((*MockRepository)(nil).Share), ctx, share)
}

// Upload mocks base method.
func (m *MockRepository) Upload(ctx context.Context, file *models.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockRepositoryMockRecorder) Upload(ctx, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockRepository)(nil).Upload), ctx, file)
}
