package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Contains(t, c.scrappers, "kijiji")
	assert.Contains(t, c.scrappers, "yandex")
}
