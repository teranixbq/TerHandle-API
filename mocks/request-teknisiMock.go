// Code generated by MockGen. DO NOT EDIT.
// Source: terhandle/internal/features/request-teknisi/entity (interfaces: UserRequestRepositoryInterface)

// Package test is a generated GoMock package.
package test

import (
	reflect "reflect"
	entity "terhandle/internal/features/request-teknisi/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRequestRepositoryInterface is a mock of UserRequestRepositoryInterface interface.
type MockUserRequestRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRequestRepositoryInterfaceMockRecorder
}

// MockUserRequestRepositoryInterfaceMockRecorder is the mock recorder for MockUserRequestRepositoryInterface.
type MockUserRequestRepositoryInterfaceMockRecorder struct {
	mock *MockUserRequestRepositoryInterface
}

// NewMockUserRequestRepositoryInterface creates a new mock instance.
func NewMockUserRequestRepositoryInterface(ctrl *gomock.Controller) *MockUserRequestRepositoryInterface {
	mock := &MockUserRequestRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRequestRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRequestRepositoryInterface) EXPECT() *MockUserRequestRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockUserRequestRepositoryInterface) Insert(arg0 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).Insert), arg0)
}

// SelectAllById mocks base method.
func (m *MockUserRequestRepositoryInterface) SelectAllById(arg0 int) ([]entity.Core, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAllById", arg0)
	ret0, _ := ret[0].([]entity.Core)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAllById indicates an expected call of SelectAllById.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) SelectAllById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAllById", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).SelectAllById), arg0)
}

// SelectById mocks base method.
func (m *MockUserRequestRepositoryInterface) SelectById(arg0, arg1 int) ([]entity.Core, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", arg0, arg1)
	ret0, _ := ret[0].([]entity.Core)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) SelectById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).SelectById), arg0, arg1)
}

// SelectByIdAndRole mocks base method.
func (m *MockUserRequestRepositoryInterface) SelectByIdAndRole(arg0, arg1 int, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByIdAndRole", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectByIdAndRole indicates an expected call of SelectByIdAndRole.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) SelectByIdAndRole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByIdAndRole", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).SelectByIdAndRole), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockUserRequestRepositoryInterface) Update(arg0 int, arg1 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).Update), arg0, arg1)
}

// UpdateClaims mocks base method.
func (m *MockUserRequestRepositoryInterface) UpdateClaims(arg0 int, arg1 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClaims", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClaims indicates an expected call of UpdateClaims.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) UpdateClaims(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClaims", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).UpdateClaims), arg0, arg1)
}

// UpdateField mocks base method.
func (m *MockUserRequestRepositoryInterface) UpdateField(arg0 int, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateField", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateField indicates an expected call of UpdateField.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) UpdateField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateField", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).UpdateField), arg0, arg1, arg2)
}

// UpdateStatusBatal mocks base method.
func (m *MockUserRequestRepositoryInterface) UpdateStatusBatal(arg0, arg1 int, arg2 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusBatal", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusBatal indicates an expected call of UpdateStatusBatal.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) UpdateStatusBatal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusBatal", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).UpdateStatusBatal), arg0, arg1, arg2)
}

// UpdateStatusKonfirmBiaya mocks base method.
func (m *MockUserRequestRepositoryInterface) UpdateStatusKonfirmBiaya(arg0, arg1 int, arg2 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusKonfirmBiaya", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusKonfirmBiaya indicates an expected call of UpdateStatusKonfirmBiaya.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) UpdateStatusKonfirmBiaya(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusKonfirmBiaya", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).UpdateStatusKonfirmBiaya), arg0, arg1, arg2)
}

// UpdateStatusSelesai mocks base method.
func (m *MockUserRequestRepositoryInterface) UpdateStatusSelesai(arg0, arg1 int, arg2 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusSelesai", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusSelesai indicates an expected call of UpdateStatusSelesai.
func (mr *MockUserRequestRepositoryInterfaceMockRecorder) UpdateStatusSelesai(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusSelesai", reflect.TypeOf((*MockUserRequestRepositoryInterface)(nil).UpdateStatusSelesai), arg0, arg1, arg2)
}