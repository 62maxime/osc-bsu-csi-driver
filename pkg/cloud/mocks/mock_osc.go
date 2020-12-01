// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud (interfaces: EC2)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	"github.com/outscale/osc-sdk-go/osc"
)




// MockOscInterface is a mock of OscInterface interface
type MockOscInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOscInterfaceMockRecorder
}

// MockOscInterfaceMockRecorder is the mock recorder for MockOscInterface
type MockOscInterfaceMockRecorder struct {
	mock *MockOscInterface
}

// NewMockOscInterface creates a new mock instance
func NewMockOscInterface(ctrl *gomock.Controller) *MockOscInterface {
	mock := &MockOscInterface{ctrl: ctrl}
	mock.recorder = &MockOscInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOscInterface) EXPECT() *MockOscInterfaceMockRecorder {
	return m.recorder
}

// CreateVolume mocks base method
func (m *MockOscInterface) CreateVolume(ctx context.Context, localVarOptionals *osc.CreateVolumeOpts) (osc.CreateVolumeResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.CreateVolumeResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockOscInterfaceMockRecorder) CreateVolume(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockOscInterface)(nil).CreateVolume), ctx, localVarOptionals)
}

// CreateTags mocks base method
func (m *MockOscInterface) CreateTags(ctx context.Context, localVarOptionals *osc.CreateTagsOpts) (osc.CreateTagsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.CreateTagsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTags indicates an expected call of CreateTags
func (mr *MockOscInterfaceMockRecorder) CreateTags(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockOscInterface)(nil).CreateTags), ctx, localVarOptionals)
}

// ReadVolumes mocks base method
func (m *MockOscInterface) ReadVolumes(ctx context.Context, localVarOptionals *osc.ReadVolumesOpts) (osc.ReadVolumesResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadVolumes", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.ReadVolumesResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadVolumes indicates an expected call of ReadVolumes
func (mr *MockOscInterfaceMockRecorder) ReadVolumes(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadVolumes", reflect.TypeOf((*MockOscInterface)(nil).ReadVolumes), ctx, localVarOptionals)
}

// DeleteVolume mocks base method
func (m *MockOscInterface) DeleteVolume(ctx context.Context, localVarOptionals *osc.DeleteVolumeOpts) (osc.DeleteVolumeResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.DeleteVolumeResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockOscInterfaceMockRecorder) DeleteVolume(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockOscInterface)(nil).DeleteVolume), ctx, localVarOptionals)
}

// LinkVolume mocks base method
func (m *MockOscInterface) LinkVolume(ctx context.Context, localVarOptionals *osc.LinkVolumeOpts) (osc.LinkVolumeResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkVolume", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.LinkVolumeResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LinkVolume indicates an expected call of LinkVolume
func (mr *MockOscInterfaceMockRecorder) LinkVolume(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkVolume", reflect.TypeOf((*MockOscInterface)(nil).LinkVolume), ctx, localVarOptionals)
}


// UnlinkVolume mocks base method
func (m *MockOscInterface) UnlinkVolume(ctx context.Context, localVarOptionals *osc.UnlinkVolumeOpts) (osc.UnlinkVolumeResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlinkVolume", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.UnlinkVolumeResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UnlinkVolume indicates an expected call of UnlinkVolume
func (mr *MockOscInterfaceMockRecorder) UnlinkVolume(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlinkVolume", reflect.TypeOf((*MockOscInterface)(nil).UnlinkVolume), ctx, localVarOptionals)
}

// CreateSnapshot mocks base method
func (m *MockOscInterface) CreateSnapshot(ctx context.Context, localVarOptionals *osc.CreateSnapshotOpts) (osc.CreateSnapshotResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.CreateSnapshotResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSnapshot indicates an expected call of CreateSnapshot
func (mr *MockOscInterfaceMockRecorder) CreateSnapshot(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockOscInterface)(nil).CreateSnapshot), ctx, localVarOptionals)
}

// ReadSnapshots mocks base method
func (m *MockOscInterface) ReadSnapshots(ctx context.Context, localVarOptionals *osc.ReadSnapshotsOpts) (osc.ReadSnapshotsResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSnapshots", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.ReadSnapshotsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadSnapshots indicates an expected call of ReadSnapshots
func (mr *MockOscInterfaceMockRecorder) ReadSnapshots(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSnapshots", reflect.TypeOf((*MockOscInterface)(nil).ReadSnapshots), ctx, localVarOptionals)
}

// DeleteSnapshot mocks base method
func (m *MockOscInterface) DeleteSnapshot(ctx context.Context, localVarOptionals *osc.DeleteSnapshotOpts) (osc.DeleteSnapshotResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.DeleteSnapshotResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot
func (mr *MockOscInterfaceMockRecorder) DeleteSnapshot(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockOscInterface)(nil).DeleteSnapshot), ctx, localVarOptionals)
}


// ReadSubregions mocks base method
func (m *MockOscInterface) ReadSubregions(ctx context.Context, localVarOptionals *osc.ReadSubregionsOpts) (osc.ReadSubregionsResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSubregions", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.ReadSubregionsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadSubregions indicates an expected call of ReadSubregions
func (mr *MockOscInterfaceMockRecorder) ReadSubregions(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSubregions", reflect.TypeOf((*MockOscInterface)(nil).ReadSubregions), ctx, localVarOptionals)
}


// ReadVms mocks base method
func (m *MockOscInterface) ReadVms(ctx context.Context, localVarOptionals *osc.ReadVmsOpts) (osc.ReadVmsResponse, *http.Response, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadVms", ctx, localVarOptionals)
	ret0, _ := ret[0].(osc.ReadVmsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadVms indicates an expected call of ReadVms
func (mr *MockOscInterfaceMockRecorder) ReadVms(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadVms", reflect.TypeOf((*MockOscInterface)(nil).ReadVms), ctx, localVarOptionals)
}