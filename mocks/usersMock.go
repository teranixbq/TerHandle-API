// Code generated by MockGen. DO NOT EDIT.
// Source: terhandle/internal/features/user/entity (interfaces: UserRepositoryInterface)

// Package test is a generated GoMock package.
package test

import (
	reflect "reflect"
	entity "terhandle/internal/features/user/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// DeleteById mocks base method.
func (m *MockUserRepositoryInterface) DeleteById(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockUserRepositoryInterfaceMockRecorder) DeleteById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).DeleteById), arg0)
}

// Insert mocks base method.
func (m *MockUserRepositoryInterface) Insert(arg0 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockUserRepositoryInterfaceMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserRepositoryInterface)(nil).Insert), arg0)
}

// Login mocks base method.
func (m *MockUserRepositoryInterface) Login(arg0, arg1 string) (entity.Core, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(entity.Core)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockUserRepositoryInterfaceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserRepositoryInterface)(nil).Login), arg0, arg1)
}

// SelectAll mocks base method.
func (m *MockUserRepositoryInterface) SelectAll() ([]entity.Core, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]entity.Core)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectAll))
}

// SelectByIdWithFeedback mocks base method.
func (m *MockUserRepositoryInterface) SelectByIdWithFeedback(arg0 uint) ([]entity.Core, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByIdWithFeedback", arg0)
	ret0, _ := ret[0].([]entity.Core)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByIdWithFeedback indicates an expected call of SelectByIdWithFeedback.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectByIdWithFeedback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByIdWithFeedback", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectByIdWithFeedback), arg0)
}

// SelectUserById mocks base method.
func (m *MockUserRepositoryInterface) SelectUserById(arg0 uint, arg1 string) (entity.Core, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserById", arg0, arg1)
	ret0, _ := ret[0].(entity.Core)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserById indicates an expected call of SelectUserById.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectUserById), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserRepositoryInterface) Update(arg0 uint, arg1 entity.Core) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepositoryInterface)(nil).Update), arg0, arg1)
}

// UpdateField mocks base method.
func (m *MockUserRepositoryInterface) UpdateField(arg0 uint, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateField", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateField indicates an expected call of UpdateField.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdateField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateField", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdateField), arg0, arg1, arg2)
}