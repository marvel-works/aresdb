// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import topology "github.com/uber/aresdb/cluster/topology"

// HealthTrackingDynamicTopoloy is an autogenerated mock type for the HealthTrackingDynamicTopoloy type
type HealthTrackingDynamicTopoloy struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *HealthTrackingDynamicTopoloy) Close() {
	_m.Called()
}

// Get provides a mock function with given fields:
func (_m *HealthTrackingDynamicTopoloy) Get() topology.Map {
	ret := _m.Called()

	var r0 topology.Map
	if rf, ok := ret.Get(0).(func() topology.Map); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(topology.Map)
		}
	}

	return r0
}

// MarkHostHealthy provides a mock function with given fields: host
func (_m *HealthTrackingDynamicTopoloy) MarkHostHealthy(host topology.Host) error {
	ret := _m.Called(host)

	var r0 error
	if rf, ok := ret.Get(0).(func(topology.Host) error); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkHostUnhealthy provides a mock function with given fields: host
func (_m *HealthTrackingDynamicTopoloy) MarkHostUnhealthy(host topology.Host) error {
	ret := _m.Called(host)

	var r0 error
	if rf, ok := ret.Get(0).(func(topology.Host) error); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields:
func (_m *HealthTrackingDynamicTopoloy) Watch() (topology.MapWatch, error) {
	ret := _m.Called()

	var r0 topology.MapWatch
	if rf, ok := ret.Get(0).(func() topology.MapWatch); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(topology.MapWatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
