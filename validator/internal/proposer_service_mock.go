// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: ProposerServiceClient)

package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	grpc "google.golang.org/grpc"
)

// MockProposerServiceClient is a mock of ProposerServiceClient interface
type MockProposerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProposerServiceClientMockRecorder
}

// MockProposerServiceClientMockRecorder is the mock recorder for MockProposerServiceClient
type MockProposerServiceClientMockRecorder struct {
	mock *MockProposerServiceClient
}

// NewMockProposerServiceClient creates a new mock instance
func NewMockProposerServiceClient(ctrl *gomock.Controller) *MockProposerServiceClient {
	mock := &MockProposerServiceClient{ctrl: ctrl}
	mock.recorder = &MockProposerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProposerServiceClient) EXPECT() *MockProposerServiceClientMockRecorder {
	return m.recorder
}

// ProposeBlock mocks base method
func (m *MockProposerServiceClient) ProposeBlock(arg0 context.Context, arg1 *v1.ProposeRequest, arg2 ...grpc.CallOption) (*v1.ProposeResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProposeBlock", varargs...)
	ret0, _ := ret[0].(*v1.ProposeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeBlock indicates an expected call of ProposeBlock
func (mr *MockProposerServiceClientMockRecorder) ProposeBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeBlock", reflect.TypeOf((*MockProposerServiceClient)(nil).ProposeBlock), varargs...)
}
