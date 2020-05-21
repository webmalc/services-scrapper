package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/services-scrapper/common/test"
)

const databaseKey = "database_uri"

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Contains(t, c.DatabaseURI, ":memory:")
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
