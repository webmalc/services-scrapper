package scrappers

import (
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type urlProcessorType func(string, *colly.Collector)

// Scrapper is the base scrapper.
type baseScrapper struct {
	id             string
	logger         Logger
	urls           []string
	allowedDomains []string
	processURL     urlProcessorType
}

func (b *baseScrapper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(b.allowedDomains...),
		colly.Async(true),
	)
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		RandomDelay: 3 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	extensions.RandomUserAgent(c)

	return c
}

func (b *baseScrapper) Scrap(wg *sync.WaitGroup) {
	for _, url := range b.urls {
		b.processURL(url, b.getCollector())
	}
	wg.Done()
}
