// Code generated by MockGen. DO NOT EDIT.
// Source: bricks.go

// Package mock_monitor is a generated GoMock package.
package mock_monitor

import (
	context "context"
	reflect "reflect"

	monitor "github.com/go-bricks/bricks/interfaces/monitor"
	gomock "github.com/golang/mock/gomock"
)

// MockBrickMetric is a mock of BrickMetric interface.
type MockBrickMetric struct {
	ctrl     *gomock.Controller
	recorder *MockBrickMetricMockRecorder
}

// MockBrickMetricMockRecorder is the mock recorder for MockBrickMetric.
type MockBrickMetricMockRecorder struct {
	mock *MockBrickMetric
}

// NewMockBrickMetric creates a new mock instance.
func NewMockBrickMetric(ctrl *gomock.Controller) *MockBrickMetric {
	mock := &MockBrickMetric{ctrl: ctrl}
	mock.recorder = &MockBrickMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrickMetric) EXPECT() *MockBrickMetricMockRecorder {
	return m.recorder
}

// MockBricksCounter is a mock of BricksCounter interface.
type MockBricksCounter struct {
	ctrl     *gomock.Controller
	recorder *MockBricksCounterMockRecorder
}

// MockBricksCounterMockRecorder is the mock recorder for MockBricksCounter.
type MockBricksCounterMockRecorder struct {
	mock *MockBricksCounter
}

// NewMockBricksCounter creates a new mock instance.
func NewMockBricksCounter(ctrl *gomock.Controller) *MockBricksCounter {
	mock := &MockBricksCounter{ctrl: ctrl}
	mock.recorder = &MockBricksCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBricksCounter) EXPECT() *MockBricksCounterMockRecorder {
	return m.recorder
}

// WithTags mocks base method.
func (m *MockBricksCounter) WithTags(tags map[string]string) (monitor.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTags", tags)
	ret0, _ := ret[0].(monitor.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTags indicates an expected call of WithTags.
func (mr *MockBricksCounterMockRecorder) WithTags(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTags", reflect.TypeOf((*MockBricksCounter)(nil).WithTags), tags)
}

// MockBricksGauge is a mock of BricksGauge interface.
type MockBricksGauge struct {
	ctrl     *gomock.Controller
	recorder *MockBricksGaugeMockRecorder
}

// MockBricksGaugeMockRecorder is the mock recorder for MockBricksGauge.
type MockBricksGaugeMockRecorder struct {
	mock *MockBricksGauge
}

// NewMockBricksGauge creates a new mock instance.
func NewMockBricksGauge(ctrl *gomock.Controller) *MockBricksGauge {
	mock := &MockBricksGauge{ctrl: ctrl}
	mock.recorder = &MockBricksGaugeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBricksGauge) EXPECT() *MockBricksGaugeMockRecorder {
	return m.recorder
}

// WithTags mocks base method.
func (m *MockBricksGauge) WithTags(tags map[string]string) (monitor.Gauge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTags", tags)
	ret0, _ := ret[0].(monitor.Gauge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTags indicates an expected call of WithTags.
func (mr *MockBricksGaugeMockRecorder) WithTags(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTags", reflect.TypeOf((*MockBricksGauge)(nil).WithTags), tags)
}

// MockBricksHistogram is a mock of BricksHistogram interface.
type MockBricksHistogram struct {
	ctrl     *gomock.Controller
	recorder *MockBricksHistogramMockRecorder
}

// MockBricksHistogramMockRecorder is the mock recorder for MockBricksHistogram.
type MockBricksHistogramMockRecorder struct {
	mock *MockBricksHistogram
}

// NewMockBricksHistogram creates a new mock instance.
func NewMockBricksHistogram(ctrl *gomock.Controller) *MockBricksHistogram {
	mock := &MockBricksHistogram{ctrl: ctrl}
	mock.recorder = &MockBricksHistogramMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBricksHistogram) EXPECT() *MockBricksHistogramMockRecorder {
	return m.recorder
}

// WithTags mocks base method.
func (m *MockBricksHistogram) WithTags(tags map[string]string) (monitor.Histogram, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTags", tags)
	ret0, _ := ret[0].(monitor.Histogram)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTags indicates an expected call of WithTags.
func (mr *MockBricksHistogramMockRecorder) WithTags(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTags", reflect.TypeOf((*MockBricksHistogram)(nil).WithTags), tags)
}

// MockBricksTimer is a mock of BricksTimer interface.
type MockBricksTimer struct {
	ctrl     *gomock.Controller
	recorder *MockBricksTimerMockRecorder
}

// MockBricksTimerMockRecorder is the mock recorder for MockBricksTimer.
type MockBricksTimerMockRecorder struct {
	mock *MockBricksTimer
}

// NewMockBricksTimer creates a new mock instance.
func NewMockBricksTimer(ctrl *gomock.Controller) *MockBricksTimer {
	mock := &MockBricksTimer{ctrl: ctrl}
	mock.recorder = &MockBricksTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBricksTimer) EXPECT() *MockBricksTimerMockRecorder {
	return m.recorder
}

// WithTags mocks base method.
func (m *MockBricksTimer) WithTags(tags map[string]string) (monitor.Timer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTags", tags)
	ret0, _ := ret[0].(monitor.Timer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTags indicates an expected call of WithTags.
func (mr *MockBricksTimerMockRecorder) WithTags(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTags", reflect.TypeOf((*MockBricksTimer)(nil).WithTags), tags)
}

// MockBricksMetrics is a mock of BricksMetrics interface.
type MockBricksMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockBricksMetricsMockRecorder
}

// MockBricksMetricsMockRecorder is the mock recorder for MockBricksMetrics.
type MockBricksMetricsMockRecorder struct {
	mock *MockBricksMetrics
}

// NewMockBricksMetrics creates a new mock instance.
func NewMockBricksMetrics(ctrl *gomock.Controller) *MockBricksMetrics {
	mock := &MockBricksMetrics{ctrl: ctrl}
	mock.recorder = &MockBricksMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBricksMetrics) EXPECT() *MockBricksMetricsMockRecorder {
	return m.recorder
}

// Counter mocks base method.
func (m *MockBricksMetrics) Counter(name, desc string, tagKeys ...string) (monitor.BricksCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, desc}
	for _, a := range tagKeys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Counter", varargs...)
	ret0, _ := ret[0].(monitor.BricksCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Counter indicates an expected call of Counter.
func (mr *MockBricksMetricsMockRecorder) Counter(name, desc interface{}, tagKeys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, desc}, tagKeys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Counter", reflect.TypeOf((*MockBricksMetrics)(nil).Counter), varargs...)
}

// Gauge mocks base method.
func (m *MockBricksMetrics) Gauge(name, desc string, tagKeys ...string) (monitor.BricksGauge, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, desc}
	for _, a := range tagKeys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Gauge", varargs...)
	ret0, _ := ret[0].(monitor.BricksGauge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gauge indicates an expected call of Gauge.
func (mr *MockBricksMetricsMockRecorder) Gauge(name, desc interface{}, tagKeys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, desc}, tagKeys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockBricksMetrics)(nil).Gauge), varargs...)
}

// Histogram mocks base method.
func (m *MockBricksMetrics) Histogram(name, desc string, buckets []float64, tagKeys ...string) (monitor.BricksHistogram, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, desc, buckets}
	for _, a := range tagKeys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Histogram", varargs...)
	ret0, _ := ret[0].(monitor.BricksHistogram)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Histogram indicates an expected call of Histogram.
func (mr *MockBricksMetricsMockRecorder) Histogram(name, desc, buckets interface{}, tagKeys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, desc, buckets}, tagKeys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Histogram", reflect.TypeOf((*MockBricksMetrics)(nil).Histogram), varargs...)
}

// Remove mocks base method.
func (m *MockBricksMetrics) Remove(metric monitor.BrickMetric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", metric)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockBricksMetricsMockRecorder) Remove(metric interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockBricksMetrics)(nil).Remove), metric)
}

// Timer mocks base method.
func (m *MockBricksMetrics) Timer(name, desc string, tagKeys ...string) (monitor.BricksTimer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, desc}
	for _, a := range tagKeys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Timer", varargs...)
	ret0, _ := ret[0].(monitor.BricksTimer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Timer indicates an expected call of Timer.
func (mr *MockBricksMetricsMockRecorder) Timer(name, desc interface{}, tagKeys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, desc}, tagKeys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timer", reflect.TypeOf((*MockBricksMetrics)(nil).Timer), varargs...)
}

// MockBricksReporter is a mock of BricksReporter interface.
type MockBricksReporter struct {
	ctrl     *gomock.Controller
	recorder *MockBricksReporterMockRecorder
}

// MockBricksReporterMockRecorder is the mock recorder for MockBricksReporter.
type MockBricksReporterMockRecorder struct {
	mock *MockBricksReporter
}

// NewMockBricksReporter creates a new mock instance.
func NewMockBricksReporter(ctrl *gomock.Controller) *MockBricksReporter {
	mock := &MockBricksReporter{ctrl: ctrl}
	mock.recorder = &MockBricksReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBricksReporter) EXPECT() *MockBricksReporterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBricksReporter) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockBricksReporterMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBricksReporter)(nil).Close), ctx)
}

// Connect mocks base method.
func (m *MockBricksReporter) Connect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockBricksReporterMockRecorder) Connect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockBricksReporter)(nil).Connect), ctx)
}

// Metrics mocks base method.
func (m *MockBricksReporter) Metrics() monitor.BricksMetrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metrics")
	ret0, _ := ret[0].(monitor.BricksMetrics)
	return ret0
}

// Metrics indicates an expected call of Metrics.
func (mr *MockBricksReporterMockRecorder) Metrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockBricksReporter)(nil).Metrics))
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockBuilder) Build() monitor.BricksReporter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(monitor.BricksReporter)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build))
}