// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typusomega/semantic-changelog-gen/pkg/git (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	changelog "github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetLog mocks base method
func (m *MockRepository) GetLog() ([]*changelog.SemanticCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLog")
	ret0, _ := ret[0].([]*changelog.SemanticCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLog indicates an expected call of GetLog
func (mr *MockRepositoryMockRecorder) GetLog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLog", reflect.TypeOf((*MockRepository)(nil).GetLog))
}
