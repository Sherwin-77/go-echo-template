// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/role.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/role.go -destination=test/mock/./repository/role.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/sherwin-77/go-echo-template/internal/entity"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockRoleRepository is a mock of RoleRepository interface.
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepositoryMockRecorder
	isgomock struct{}
}

// MockRoleRepositoryMockRecorder is the mock recorder for MockRoleRepository.
type MockRoleRepositoryMockRecorder struct {
	mock *MockRoleRepository
}

// NewMockRoleRepository creates a new mock instance.
func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &MockRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleRepository) EXPECT() *MockRoleRepositoryMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockRoleRepository) BeginTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockRoleRepositoryMockRecorder) BeginTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockRoleRepository)(nil).BeginTransaction))
}

// Commit mocks base method.
func (m *MockRoleRepository) Commit(tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockRoleRepositoryMockRecorder) Commit(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockRoleRepository)(nil).Commit), tx)
}

// CreateRole mocks base method.
func (m *MockRoleRepository) CreateRole(ctx context.Context, tx *gorm.DB, role *entity.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", ctx, tx, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRoleRepositoryMockRecorder) CreateRole(ctx, tx, role any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRoleRepository)(nil).CreateRole), ctx, tx, role)
}

// DeleteRole mocks base method.
func (m *MockRoleRepository) DeleteRole(ctx context.Context, tx *gorm.DB, role *entity.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", ctx, tx, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRoleRepositoryMockRecorder) DeleteRole(ctx, tx, role any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRoleRepository)(nil).DeleteRole), ctx, tx, role)
}

// GetRoleByID mocks base method.
func (m *MockRoleRepository) GetRoleByID(ctx context.Context, tx *gorm.DB, id string) (*entity.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByID", ctx, tx, id)
	ret0, _ := ret[0].(*entity.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleByID indicates an expected call of GetRoleByID.
func (mr *MockRoleRepositoryMockRecorder) GetRoleByID(ctx, tx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByID", reflect.TypeOf((*MockRoleRepository)(nil).GetRoleByID), ctx, tx, id)
}

// GetRoles mocks base method.
func (m *MockRoleRepository) GetRoles(ctx context.Context, tx *gorm.DB) ([]entity.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", ctx, tx)
	ret0, _ := ret[0].([]entity.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles.
func (mr *MockRoleRepositoryMockRecorder) GetRoles(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockRoleRepository)(nil).GetRoles), ctx, tx)
}

// GetRolesFiltered mocks base method.
func (m *MockRoleRepository) GetRolesFiltered(ctx context.Context, tx *gorm.DB, limit, offset int, order, query any, args ...any) ([]entity.Role, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, tx, limit, offset, order, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRolesFiltered", varargs...)
	ret0, _ := ret[0].([]entity.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRolesFiltered indicates an expected call of GetRolesFiltered.
func (mr *MockRoleRepositoryMockRecorder) GetRolesFiltered(ctx, tx, limit, offset, order, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, limit, offset, order, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRolesFiltered", reflect.TypeOf((*MockRoleRepository)(nil).GetRolesFiltered), varargs...)
}

// Rollback mocks base method.
func (m *MockRoleRepository) Rollback(tx *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", tx)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockRoleRepositoryMockRecorder) Rollback(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockRoleRepository)(nil).Rollback), tx)
}

// SingleTransaction mocks base method.
func (m *MockRoleRepository) SingleTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingleTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// SingleTransaction indicates an expected call of SingleTransaction.
func (mr *MockRoleRepositoryMockRecorder) SingleTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingleTransaction", reflect.TypeOf((*MockRoleRepository)(nil).SingleTransaction))
}

// UpdateRole mocks base method.
func (m *MockRoleRepository) UpdateRole(ctx context.Context, tx *gorm.DB, role *entity.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", ctx, tx, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRoleRepositoryMockRecorder) UpdateRole(ctx, tx, role any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRoleRepository)(nil).UpdateRole), ctx, tx, role)
}

// WithTransaction mocks base method.
func (m *MockRoleRepository) WithTransaction(fn func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockRoleRepositoryMockRecorder) WithTransaction(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockRoleRepository)(nil).WithTransaction), fn)
}
