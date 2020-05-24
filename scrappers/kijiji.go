package scrappers

import (
	"fmt"
	"sync"
	"time"
)

const kijijiID = "kijiji"

// Kijiji is the kijiji scrapper.
type Kijiji struct {
	baseScrapper
}

// Run runs the scrappers.
func (k *Kijiji) Scrap(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Duration(1) * time.Millisecond)
		fmt.Println(k.ID)
	}
	wg.Done()
}

// NewKijiji creates a new Kijiji instance.
func NewKijiji() *Kijiji {
	return &Kijiji{baseScrapper: baseScrapper{ID: kijijiID}}
}
