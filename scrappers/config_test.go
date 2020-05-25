package scrappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Contains(t, c.kijijiURLs[0], "ontario")
	assert.Contains(t, c.kijijiURLs[1], "quebec")
}
