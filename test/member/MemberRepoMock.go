// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repos/interfaces/MemberInterface.go

// Package member is a generated GoMock package.
package member

import (
	models "golang/pkg/repos/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMemberRepo is a mock of MemberRepo interface.
type MockMemberRepo struct {
	ctrl     *gomock.Controller
	recorder *MockMemberRepoMockRecorder
}

// MockMemberRepoMockRecorder is the mock recorder for MockMemberRepo.
type MockMemberRepoMockRecorder struct {
	mock *MockMemberRepo
}

// NewMockMemberRepo creates a new mock instance.
func NewMockMemberRepo(ctrl *gomock.Controller) *MockMemberRepo {
	mock := &MockMemberRepo{ctrl: ctrl}
	mock.recorder = &MockMemberRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberRepo) EXPECT() *MockMemberRepoMockRecorder {
	return m.recorder
}

// ChangePwd mocks base method.
func (m *MockMemberRepo) ChangePwd(memberId, newPwd string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePwd", memberId, newPwd)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePwd indicates an expected call of ChangePwd.
func (mr *MockMemberRepoMockRecorder) ChangePwd(memberId, newPwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePwd", reflect.TypeOf((*MockMemberRepo)(nil).ChangePwd), memberId, newPwd)
}

// Create mocks base method.
func (m *MockMemberRepo) Create(member models.MemberModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", member)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMemberRepoMockRecorder) Create(member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMemberRepo)(nil).Create), member)
}

// GetMember mocks base method.
func (m *MockMemberRepo) GetMember(memberId string) (models.MemberModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMember", memberId)
	ret0, _ := ret[0].(models.MemberModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMember indicates an expected call of GetMember.
func (mr *MockMemberRepoMockRecorder) GetMember(memberId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockMemberRepo)(nil).GetMember), memberId)
}

// GetMemberByEmail mocks base method.
func (m *MockMemberRepo) GetMemberByEmail(email string) (models.MemberModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberByEmail", email)
	ret0, _ := ret[0].(models.MemberModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberByEmail indicates an expected call of GetMemberByEmail.
func (mr *MockMemberRepoMockRecorder) GetMemberByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberByEmail", reflect.TypeOf((*MockMemberRepo)(nil).GetMemberByEmail), email)
}

// IsEmailExist mocks base method.
func (m *MockMemberRepo) IsEmailExist(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmailExist", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmailExist indicates an expected call of IsEmailExist.
func (mr *MockMemberRepoMockRecorder) IsEmailExist(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmailExist", reflect.TypeOf((*MockMemberRepo)(nil).IsEmailExist), email)
}
