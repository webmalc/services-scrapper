package test

import (
	"os"
	"testing"

	"github.com/webmalc/services-scrapper/common/config"
)

// Setups the tests
func setUp() {
	os.Setenv("SS_ENV", "test")
	config.Setup()
}

// Run setups, runs and teardowns the tests
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
