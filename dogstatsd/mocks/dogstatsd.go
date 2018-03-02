// Automatically generated by MockGen. DO NOT EDIT!
// Source: dogstatsd/dogstatsd.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	time "time"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Incr(name string, tags []string, rate float64) error {
	ret := _m.ctrl.Call(_m, "Incr", name, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Incr(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Incr", arg0, arg1, arg2)
}

func (_m *MockClient) Count(name string, value int64, tags []string, rate float64) error {
	ret := _m.ctrl.Call(_m, "Count", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Count(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Count", arg0, arg1, arg2, arg3)
}

func (_m *MockClient) Gauge(name string, value float64, tags []string, rate float64) error {
	ret := _m.ctrl.Call(_m, "Gauge", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Gauge(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Gauge", arg0, arg1, arg2, arg3)
}

func (_m *MockClient) Timing(name string, value time.Duration, tags []string, rate float64) error {
	ret := _m.ctrl.Call(_m, "Timing", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Timing(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Timing", arg0, arg1, arg2, arg3)
}

func (_m *MockClient) Histogram(name string, value float64, tags []string, rate float64) error {
	ret := _m.ctrl.Call(_m, "Histogram", name, value, tags, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Histogram(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Histogram", arg0, arg1, arg2, arg3)
}
