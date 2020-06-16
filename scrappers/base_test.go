package scrappers

import (
	"sync"
	"testing"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/services-scrapper/scrappers/mocks"
)

func Test_baseScrapper_getCollector(t *testing.T) {
	domain := "example.com"
	log := &mocks.Logger{}
	scrapper := &baseScrapper{
		id:             "test",
		logger:         log,
		allowedDomains: []string{domain},
	}
	c := scrapper.getCollector()
	assert.Contains(t, c.AllowedDomains, domain)
}

func Test_baseScrapper_Scrap(t *testing.T) {
	log := &mocks.Logger{}
	var wg sync.WaitGroup
	count := 0
	scrapper := &baseScrapper{
		id:     "test",
		logger: log,
		urls:   []string{"one", "two"},
		processURL: func(
			url string, c *colly.Collector, r serviceRepository,
		) {
			count++
		},
	}
	wg.Add(1)
	scrapper.Scrap(&wg)
	wg.Wait()
	assert.Equal(t, len(scrapper.urls), count)
}
