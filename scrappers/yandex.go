package scrappers

import (
	"sync"
)

const yandexID = "yandex"

// Yandex is the yandex scrapper.
type Yandex struct {
	baseScrapper
}

// Scrap runs the scrapper.
func (k *Yandex) Scrap(wg *sync.WaitGroup) {
	wg.Done()
}

// NewYandex creates a new Yandex instance.
func NewYandex(log logger) *Yandex {
	return &Yandex{baseScrapper: baseScrapper{id: yandexID, logger: log}}
}
