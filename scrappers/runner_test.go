package scrappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/services-scrapper/common/test"
	"github.com/webmalc/services-scrapper/scrappers/mocks"
)

// Should run the scrappers.
func TestRunner_Run(t *testing.T) {
	log := &mocks.Logger{}
	runner := NewRunner(log)
	log.On("Infof", mock.Anything, mock.Anything).Return(nil).Once()
	runner.Run([]string{"kijiji", "yandex"})
	log.AssertExpectations(t)
}

// Should create a runner object.
func TestNewRunner(t *testing.T) {
	log := &mocks.Logger{}
	runner := NewRunner(log)
	assert.Equal(t, runner.logger, log)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
