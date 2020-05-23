package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Logger logs errors.
type Logger struct {
	mock.Mock
}

// Error is a method mock
func (m *Logger) Error(args ...interface{}) {
	m.Called(args...)
}

// Infof is a method mock
func (m *Logger) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}
