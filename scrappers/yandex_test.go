package scrappers

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/services-scrapper/scrappers/mocks"
)

func TestYandex_Scrap(t *testing.T) {
	log := &mocks.Logger{}
	var wg sync.WaitGroup
	yandex := NewYandex(log)
	wg.Add(1)
	yandex.Scrap(&wg)
	wg.Wait()
}

func TestNewYandex(t *testing.T) {
	log := &mocks.Logger{}
	yandex := NewYandex(log)
	assert.Equal(t, yandex.logger, log)
	assert.Equal(t, yandex.id, yandexID)
}
