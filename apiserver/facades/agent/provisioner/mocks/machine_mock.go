// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/provisioner (interfaces: Machine)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	set "github.com/juju/collections/set"
	constraints "github.com/juju/juju/constraints"
	instance "github.com/juju/juju/instance"
	containerizer "github.com/juju/juju/network/containerizer"
	state "github.com/juju/juju/state"
	names_v2 "gopkg.in/juju/names.v2"
	reflect "reflect"
)

// MockMachine is a mock of Machine interface
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// AllLinkLayerDevices mocks base method
func (m *MockMachine) AllLinkLayerDevices() ([]containerizer.LinkLayerDevice, error) {
	ret := m.ctrl.Call(m, "AllLinkLayerDevices")
	ret0, _ := ret[0].([]containerizer.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLinkLayerDevices indicates an expected call of AllLinkLayerDevices
func (mr *MockMachineMockRecorder) AllLinkLayerDevices() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLinkLayerDevices", reflect.TypeOf((*MockMachine)(nil).AllLinkLayerDevices))
}

// AllSpaces mocks base method
func (m *MockMachine) AllSpaces() (set.Strings, error) {
	ret := m.ctrl.Call(m, "AllSpaces")
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaces indicates an expected call of AllSpaces
func (mr *MockMachineMockRecorder) AllSpaces() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaces", reflect.TypeOf((*MockMachine)(nil).AllSpaces))
}

// ContainerType mocks base method
func (m *MockMachine) ContainerType() instance.ContainerType {
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType
func (mr *MockMachineMockRecorder) ContainerType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockMachine)(nil).ContainerType))
}

// DesiredSpaces mocks base method
func (m *MockMachine) DesiredSpaces() (set.Strings, error) {
	ret := m.ctrl.Call(m, "DesiredSpaces")
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DesiredSpaces indicates an expected call of DesiredSpaces
func (mr *MockMachineMockRecorder) DesiredSpaces() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DesiredSpaces", reflect.TypeOf((*MockMachine)(nil).DesiredSpaces))
}

// Id mocks base method
func (m *MockMachine) Id() string {
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockMachineMockRecorder) Id() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMachine)(nil).Id))
}

// InstanceId mocks base method
func (m *MockMachine) InstanceId() (instance.Id, error) {
	ret := m.ctrl.Call(m, "InstanceId")
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceId indicates an expected call of InstanceId
func (mr *MockMachineMockRecorder) InstanceId() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceId", reflect.TypeOf((*MockMachine)(nil).InstanceId))
}

// IsManual mocks base method
func (m *MockMachine) IsManual() (bool, error) {
	ret := m.ctrl.Call(m, "IsManual")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManual indicates an expected call of IsManual
func (mr *MockMachineMockRecorder) IsManual() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManual", reflect.TypeOf((*MockMachine)(nil).IsManual))
}

// LinkLayerDevicesForSpaces mocks base method
func (m *MockMachine) LinkLayerDevicesForSpaces(arg0 []string) (map[string][]containerizer.LinkLayerDevice, error) {
	ret := m.ctrl.Call(m, "LinkLayerDevicesForSpaces", arg0)
	ret0, _ := ret[0].(map[string][]containerizer.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkLayerDevicesForSpaces indicates an expected call of LinkLayerDevicesForSpaces
func (mr *MockMachineMockRecorder) LinkLayerDevicesForSpaces(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkLayerDevicesForSpaces", reflect.TypeOf((*MockMachine)(nil).LinkLayerDevicesForSpaces), arg0)
}

// MachineTag mocks base method
func (m *MockMachine) MachineTag() names_v2.MachineTag {
	ret := m.ctrl.Call(m, "MachineTag")
	ret0, _ := ret[0].(names_v2.MachineTag)
	return ret0
}

// MachineTag indicates an expected call of MachineTag
func (mr *MockMachineMockRecorder) MachineTag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineTag", reflect.TypeOf((*MockMachine)(nil).MachineTag))
}

// Raw mocks base method
func (m *MockMachine) Raw() *state.Machine {
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(*state.Machine)
	return ret0
}

// Raw indicates an expected call of Raw
func (mr *MockMachineMockRecorder) Raw() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockMachine)(nil).Raw))
}

// RemoveAllAddresses mocks base method
func (m *MockMachine) RemoveAllAddresses() error {
	ret := m.ctrl.Call(m, "RemoveAllAddresses")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllAddresses indicates an expected call of RemoveAllAddresses
func (mr *MockMachineMockRecorder) RemoveAllAddresses() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllAddresses", reflect.TypeOf((*MockMachine)(nil).RemoveAllAddresses))
}

// SetConstraints mocks base method
func (m *MockMachine) SetConstraints(arg0 constraints.Value) error {
	ret := m.ctrl.Call(m, "SetConstraints", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints
func (mr *MockMachineMockRecorder) SetConstraints(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockMachine)(nil).SetConstraints), arg0)
}

// SetDevicesAddresses mocks base method
func (m *MockMachine) SetDevicesAddresses(arg0 ...state.LinkLayerDeviceAddress) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetDevicesAddresses", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDevicesAddresses indicates an expected call of SetDevicesAddresses
func (mr *MockMachineMockRecorder) SetDevicesAddresses(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDevicesAddresses", reflect.TypeOf((*MockMachine)(nil).SetDevicesAddresses), arg0...)
}

// SetLinkLayerDevices mocks base method
func (m *MockMachine) SetLinkLayerDevices(arg0 ...state.LinkLayerDeviceArgs) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetLinkLayerDevices", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLinkLayerDevices indicates an expected call of SetLinkLayerDevices
func (mr *MockMachineMockRecorder) SetLinkLayerDevices(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLinkLayerDevices", reflect.TypeOf((*MockMachine)(nil).SetLinkLayerDevices), arg0...)
}

// SetParentLinkLayerDevicesBeforeTheirChildren mocks base method
func (m *MockMachine) SetParentLinkLayerDevicesBeforeTheirChildren(arg0 []state.LinkLayerDeviceArgs) error {
	ret := m.ctrl.Call(m, "SetParentLinkLayerDevicesBeforeTheirChildren", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetParentLinkLayerDevicesBeforeTheirChildren indicates an expected call of SetParentLinkLayerDevicesBeforeTheirChildren
func (mr *MockMachineMockRecorder) SetParentLinkLayerDevicesBeforeTheirChildren(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParentLinkLayerDevicesBeforeTheirChildren", reflect.TypeOf((*MockMachine)(nil).SetParentLinkLayerDevicesBeforeTheirChildren), arg0)
}

// Units mocks base method
func (m *MockMachine) Units() ([]containerizer.Unit, error) {
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]containerizer.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockMachineMockRecorder) Units() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockMachine)(nil).Units))
}