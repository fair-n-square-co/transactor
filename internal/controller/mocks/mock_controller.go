// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go
//
// Generated by this command:
//
//	mockgen -source=controller.go -destination=mocks/mock_controller.go -package=controllermocks
//

// Package controllermocks is a generated GoMock package.
package controllermocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transactions/v1alpha1"
	gomock "go.uber.org/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method.
func (m *MockController) CreateGroup(ctx context.Context, req *v1alpha1.CreateGroupRequest) (*v1alpha1.CreateGroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, req)
	ret0, _ := ret[0].(*v1alpha1.CreateGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockControllerMockRecorder) CreateGroup(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockController)(nil).CreateGroup), ctx, req)
}

// ListGroups mocks base method.
func (m *MockController) ListGroups(ctx context.Context, req *v1alpha1.ListGroupsRequest) (*v1alpha1.ListGroupsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", ctx, req)
	ret0, _ := ret[0].(*v1alpha1.ListGroupsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockControllerMockRecorder) ListGroups(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockController)(nil).ListGroups), ctx, req)
}
