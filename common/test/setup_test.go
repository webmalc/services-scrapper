package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert.Equal(t, "test", os.Getenv("SS_ENV"))
}

func TestMain(m *testing.M) {
	Run(m)
}
