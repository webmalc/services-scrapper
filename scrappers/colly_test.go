package scrappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_requestJSON(t *testing.T) {
	type test struct {
		One string
		Key string
	}
	s := &test{}
	_ = requestJSON("http://echo.jsontest.com/key/value/one/two", s, nil)
	assert.Equal(t, s.One, "two")
	assert.Equal(t, s.Key, "value")
}
