// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	http "net/http"
	reflect "reflect"

	client "github.com/go-bricks/bricks/v2/interfaces/http/client"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockHTTPClientBuilder is a mock of HTTPClientBuilder interface.
type MockHTTPClientBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientBuilderMockRecorder
}

// MockHTTPClientBuilderMockRecorder is the mock recorder for MockHTTPClientBuilder.
type MockHTTPClientBuilderMockRecorder struct {
	mock *MockHTTPClientBuilder
}

// NewMockHTTPClientBuilder creates a new mock instance.
func NewMockHTTPClientBuilder(ctrl *gomock.Controller) *MockHTTPClientBuilder {
	mock := &MockHTTPClientBuilder{ctrl: ctrl}
	mock.recorder = &MockHTTPClientBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClientBuilder) EXPECT() *MockHTTPClientBuilderMockRecorder {
	return m.recorder
}

// AddInterceptors mocks base method.
func (m *MockHTTPClientBuilder) AddInterceptors(arg0 ...client.HTTPClientInterceptor) client.HTTPClientBuilder {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddInterceptors", varargs...)
	ret0, _ := ret[0].(client.HTTPClientBuilder)
	return ret0
}

// AddInterceptors indicates an expected call of AddInterceptors.
func (mr *MockHTTPClientBuilderMockRecorder) AddInterceptors(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInterceptors", reflect.TypeOf((*MockHTTPClientBuilder)(nil).AddInterceptors), arg0...)
}

// Build mocks base method.
func (m *MockHTTPClientBuilder) Build() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockHTTPClientBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockHTTPClientBuilder)(nil).Build))
}

// WithPreconfiguredClient mocks base method.
func (m *MockHTTPClientBuilder) WithPreconfiguredClient(arg0 *http.Client) client.HTTPClientBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithPreconfiguredClient", arg0)
	ret0, _ := ret[0].(client.HTTPClientBuilder)
	return ret0
}

// WithPreconfiguredClient indicates an expected call of WithPreconfiguredClient.
func (mr *MockHTTPClientBuilderMockRecorder) WithPreconfiguredClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithPreconfiguredClient", reflect.TypeOf((*MockHTTPClientBuilder)(nil).WithPreconfiguredClient), arg0)
}

// MockGRPCClientConnectionWrapper is a mock of GRPCClientConnectionWrapper interface.
type MockGRPCClientConnectionWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCClientConnectionWrapperMockRecorder
}

// MockGRPCClientConnectionWrapperMockRecorder is the mock recorder for MockGRPCClientConnectionWrapper.
type MockGRPCClientConnectionWrapperMockRecorder struct {
	mock *MockGRPCClientConnectionWrapper
}

// NewMockGRPCClientConnectionWrapper creates a new mock instance.
func NewMockGRPCClientConnectionWrapper(ctrl *gomock.Controller) *MockGRPCClientConnectionWrapper {
	mock := &MockGRPCClientConnectionWrapper{ctrl: ctrl}
	mock.recorder = &MockGRPCClientConnectionWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCClientConnectionWrapper) EXPECT() *MockGRPCClientConnectionWrapperMockRecorder {
	return m.recorder
}

// Dial mocks base method.
func (m *MockGRPCClientConnectionWrapper) Dial(ctx context.Context, target string, extraOptions ...grpc.DialOption) (grpc.ClientConnInterface, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, target}
	for _, a := range extraOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dial", varargs...)
	ret0, _ := ret[0].(grpc.ClientConnInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dial indicates an expected call of Dial.
func (mr *MockGRPCClientConnectionWrapperMockRecorder) Dial(ctx, target interface{}, extraOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, target}, extraOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*MockGRPCClientConnectionWrapper)(nil).Dial), varargs...)
}

// MockGRPCClientConnectionBuilder is a mock of GRPCClientConnectionBuilder interface.
type MockGRPCClientConnectionBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCClientConnectionBuilderMockRecorder
}

// MockGRPCClientConnectionBuilderMockRecorder is the mock recorder for MockGRPCClientConnectionBuilder.
type MockGRPCClientConnectionBuilderMockRecorder struct {
	mock *MockGRPCClientConnectionBuilder
}

// NewMockGRPCClientConnectionBuilder creates a new mock instance.
func NewMockGRPCClientConnectionBuilder(ctrl *gomock.Controller) *MockGRPCClientConnectionBuilder {
	mock := &MockGRPCClientConnectionBuilder{ctrl: ctrl}
	mock.recorder = &MockGRPCClientConnectionBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCClientConnectionBuilder) EXPECT() *MockGRPCClientConnectionBuilderMockRecorder {
	return m.recorder
}

// AddOptions mocks base method.
func (m *MockGRPCClientConnectionBuilder) AddOptions(opts ...grpc.DialOption) client.GRPCClientConnectionBuilder {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddOptions", varargs...)
	ret0, _ := ret[0].(client.GRPCClientConnectionBuilder)
	return ret0
}

// AddOptions indicates an expected call of AddOptions.
func (mr *MockGRPCClientConnectionBuilderMockRecorder) AddOptions(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOptions", reflect.TypeOf((*MockGRPCClientConnectionBuilder)(nil).AddOptions), opts...)
}

// Build mocks base method.
func (m *MockGRPCClientConnectionBuilder) Build() client.GRPCClientConnectionWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(client.GRPCClientConnectionWrapper)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockGRPCClientConnectionBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockGRPCClientConnectionBuilder)(nil).Build))
}
