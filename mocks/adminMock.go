// Code generated by MockGen. DO NOT EDIT.
// Source: terhandle/internal/features/admin/entity (interfaces: AdminRepositoryInterface)

// Package test is a generated GoMock package.
package test

import (
	reflect "reflect"
	entity "terhandle/internal/features/admin/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockAdminRepositoryInterface is a mock of AdminRepositoryInterface interface.
type MockAdminRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryInterfaceMockRecorder
}

// MockAdminRepositoryInterfaceMockRecorder is the mock recorder for MockAdminRepositoryInterface.
type MockAdminRepositoryInterfaceMockRecorder struct {
	mock *MockAdminRepositoryInterface
}

// NewMockAdminRepositoryInterface creates a new mock instance.
func NewMockAdminRepositoryInterface(ctrl *gomock.Controller) *MockAdminRepositoryInterface {
	mock := &MockAdminRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepositoryInterface) EXPECT() *MockAdminRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAdminRepositoryInterface) Delete(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminRepositoryInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).Delete), arg0)
}

// Insert mocks base method.
func (m *MockAdminRepositoryInterface) Insert(arg0 entity.CoreProfesi) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAdminRepositoryInterfaceMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).Insert), arg0)
}

// InsertBiaya mocks base method.
func (m *MockAdminRepositoryInterface) InsertBiaya(arg0 entity.CoreTransport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBiaya", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBiaya indicates an expected call of InsertBiaya.
func (mr *MockAdminRepositoryInterfaceMockRecorder) InsertBiaya(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBiaya", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).InsertBiaya), arg0)
}

// SelectAll mocks base method.
func (m *MockAdminRepositoryInterface) SelectAll() ([]entity.CoreProfesi, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]entity.CoreProfesi)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockAdminRepositoryInterfaceMockRecorder) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).SelectAll))
}

// Update mocks base method.
func (m *MockAdminRepositoryInterface) Update(arg0 uint, arg1 entity.CoreProfesi) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAdminRepositoryInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).Update), arg0, arg1)
}

// UpdateBiaya mocks base method.
func (m *MockAdminRepositoryInterface) UpdateBiaya(arg0 uint, arg1 entity.CoreTransport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBiaya", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBiaya indicates an expected call of UpdateBiaya.
func (mr *MockAdminRepositoryInterfaceMockRecorder) UpdateBiaya(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBiaya", reflect.TypeOf((*MockAdminRepositoryInterface)(nil).UpdateBiaya), arg0, arg1)
}
