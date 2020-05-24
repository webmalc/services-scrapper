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
	kijiji := &mocks.Scrapper{}
	yandex := &mocks.Scrapper{}
	runner.scrappers = map[string]Scrapper{kijijiID: kijiji, yandexID: yandex}
	log.On("Infof", mock.Anything, mock.Anything).Return(nil).Once()
	kijiji.On("Scrap", mock.Anything, mock.Anything).Return(nil).Once()
	yandex.On("Scrap", mock.Anything, mock.Anything).Return(nil).Once()
	runner.Run([]string{})
	log.AssertExpectations(t)
}

// Should create a runner object.
func TestNewRunner(t *testing.T) {
	log := &mocks.Logger{}
	runner := NewRunner(log)
	assert.Equal(t, runner.logger, log)
	assert.Len(t, runner.scrappers, 2)
}

// Should return a list of scrappers
func TestRunner_getScrappers(t *testing.T) {
	log := &mocks.Logger{}
	runner := NewRunner(log)
	scrappers := runner.getScrappers([]string{})
	assert.Len(t, scrappers, 2)
	assert.Equal(t, scrappers[0], NewKijiji())

	scrappers = runner.getScrappers([]string{yandexID})
	assert.Len(t, scrappers, 1)
	assert.Equal(t, scrappers[0], NewYandex())
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
