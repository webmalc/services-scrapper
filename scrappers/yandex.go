package scrappers

import (
	"fmt"
	"sync"
	"time"
)

const yandexID = "yandex"

// Yandex is the yandex scrapper.
type Yandex struct {
	baseScrapper
}

// Run runs the scrappers.
func (k *Yandex) Scrap(wg *sync.WaitGroup) {
	n := 2
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Println(k.ID)
	}
	wg.Done()
}

// NewYandex creates a new Yandex instance.
func NewYandex() *Yandex {
	return &Yandex{baseScrapper: baseScrapper{ID: yandexID}}
}
