// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ClickHouse/clickhouse-go/v2/lib/driver (interfaces: Rows)
//
// Generated by this command:
//
//	mockgen -destination=./mock/rows_mock.go -package mock github.com/ClickHouse/clickhouse-go/v2/lib/driver Rows
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	driver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	gomock "go.uber.org/mock/gomock"
)

// MockRows is a mock of Rows interface.
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows.
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance.
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRows) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRowsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRows)(nil).Close))
}

// ColumnTypes mocks base method.
func (m *MockRows) ColumnTypes() []driver.ColumnType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColumnTypes")
	ret0, _ := ret[0].([]driver.ColumnType)
	return ret0
}

// ColumnTypes indicates an expected call of ColumnTypes.
func (mr *MockRowsMockRecorder) ColumnTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColumnTypes", reflect.TypeOf((*MockRows)(nil).ColumnTypes))
}

// Columns mocks base method.
func (m *MockRows) Columns() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Columns")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Columns indicates an expected call of Columns.
func (mr *MockRowsMockRecorder) Columns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockRows)(nil).Columns))
}

// Err mocks base method.
func (m *MockRows) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockRowsMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRows)(nil).Err))
}

// Next mocks base method.
func (m *MockRows) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRowsMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRows)(nil).Next))
}

// Scan mocks base method.
func (m *MockRows) Scan(arg0 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRowsMockRecorder) Scan(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRows)(nil).Scan), arg0...)
}

// ScanStruct mocks base method.
func (m *MockRows) ScanStruct(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanStruct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanStruct indicates an expected call of ScanStruct.
func (mr *MockRowsMockRecorder) ScanStruct(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanStruct", reflect.TypeOf((*MockRows)(nil).ScanStruct), arg0)
}

// Totals mocks base method.
func (m *MockRows) Totals(arg0 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Totals", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Totals indicates an expected call of Totals.
func (mr *MockRowsMockRecorder) Totals(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Totals", reflect.TypeOf((*MockRows)(nil).Totals), arg0...)
}
