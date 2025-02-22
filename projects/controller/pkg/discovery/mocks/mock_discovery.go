// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/controller/pkg/discovery (interfaces: DiscoveryPlugin)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	discovery "github.com/solo-io/gloo/projects/controller/pkg/discovery"
	plugins "github.com/solo-io/gloo/projects/controller/pkg/plugins"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockDiscoveryPlugin is a mock of DiscoveryPlugin interface.
type MockDiscoveryPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryPluginMockRecorder
}

// MockDiscoveryPluginMockRecorder is the mock recorder for MockDiscoveryPlugin.
type MockDiscoveryPluginMockRecorder struct {
	mock *MockDiscoveryPlugin
}

// NewMockDiscoveryPlugin creates a new mock instance.
func NewMockDiscoveryPlugin(ctrl *gomock.Controller) *MockDiscoveryPlugin {
	mock := &MockDiscoveryPlugin{ctrl: ctrl}
	mock.recorder = &MockDiscoveryPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscoveryPlugin) EXPECT() *MockDiscoveryPluginMockRecorder {
	return m.recorder
}

// DiscoverUpstreams mocks base method.
func (m *MockDiscoveryPlugin) DiscoverUpstreams(arg0 []string, arg1 string, arg2 clients.WatchOpts, arg3 discovery.Opts) (chan v1.UpstreamList, chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverUpstreams", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(chan v1.UpstreamList)
	ret1, _ := ret[1].(chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DiscoverUpstreams indicates an expected call of DiscoverUpstreams.
func (mr *MockDiscoveryPluginMockRecorder) DiscoverUpstreams(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverUpstreams", reflect.TypeOf((*MockDiscoveryPlugin)(nil).DiscoverUpstreams), arg0, arg1, arg2, arg3)
}

// Init mocks base method.
func (m *MockDiscoveryPlugin) Init(arg0 plugins.InitParams) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0)
}

// Init indicates an expected call of Init.
func (mr *MockDiscoveryPluginMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDiscoveryPlugin)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockDiscoveryPlugin) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDiscoveryPluginMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDiscoveryPlugin)(nil).Name))
}

// UpdateUpstream mocks base method.
func (m *MockDiscoveryPlugin) UpdateUpstream(arg0, arg1 *v1.Upstream) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUpstream", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUpstream indicates an expected call of UpdateUpstream.
func (mr *MockDiscoveryPluginMockRecorder) UpdateUpstream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUpstream", reflect.TypeOf((*MockDiscoveryPlugin)(nil).UpdateUpstream), arg0, arg1)
}

// WatchEndpoints mocks base method.
func (m *MockDiscoveryPlugin) WatchEndpoints(arg0 string, arg1 v1.UpstreamList, arg2 clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchEndpoints", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan v1.EndpointList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WatchEndpoints indicates an expected call of WatchEndpoints.
func (mr *MockDiscoveryPluginMockRecorder) WatchEndpoints(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchEndpoints", reflect.TypeOf((*MockDiscoveryPlugin)(nil).WatchEndpoints), arg0, arg1, arg2)
}
