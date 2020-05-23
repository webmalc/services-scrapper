package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Runner mocks the object
type Runner struct {
	mock.Mock
}

// Error is method mock
func (r *Runner) Run(names []string) {
	r.Called(names)
}
