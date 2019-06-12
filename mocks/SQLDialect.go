// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import exp "github.com/doug-martin/goqu/v7/exp"

import mock "github.com/stretchr/testify/mock"
import sb "github.com/doug-martin/goqu/v7/internal/sb"

// SQLDialect is an autogenerated mock type for the SQLDialect type
type SQLDialect struct {
	mock.Mock
}

// ToDeleteSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToDeleteSQL(b sb.SQLBuilder, clauses exp.Clauses) {
	_m.Called(b, clauses)
}

// ToInsertSQL provides a mock function with given fields: b, clauses, ie
func (_m *SQLDialect) ToInsertSQL(b sb.SQLBuilder, clauses exp.Clauses, ie exp.InsertExpression) {
	_m.Called(b, clauses, ie)
}

// ToSelectSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToSelectSQL(b sb.SQLBuilder, clauses exp.Clauses) {
	_m.Called(b, clauses)
}

// ToTruncateSQL provides a mock function with given fields: b, clauses, options
func (_m *SQLDialect) ToTruncateSQL(b sb.SQLBuilder, clauses exp.Clauses, options exp.TruncateOptions) {
	_m.Called(b, clauses, options)
}

// ToUpdateSQL provides a mock function with given fields: b, clauses, update
func (_m *SQLDialect) ToUpdateSQL(b sb.SQLBuilder, clauses exp.Clauses, update interface{}) {
	_m.Called(b, clauses, update)
}