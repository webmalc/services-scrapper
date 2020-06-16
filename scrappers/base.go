package scrappers

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/pkg/errors"
)

// Scrapper is the base scrapper.
type baseScrapper struct {
	id             string
	logger         logger
	repo           serviceRepository
	urls           []string
	allowedDomains []string
	processURL     urlProcessorType
}

// getCollector gets a colly Collector
func (b *baseScrapper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(b.allowedDomains...),
		colly.Async(true),
	)
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	})
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		RandomDelay: 3 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		b.logger.Error(
			"Request URL:", r.Request.URL,
			"failed with response:", r,
			"Error:", errors.Wrap(err, b.id),
		)
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		b.logger.Infof("%s: visiting the URL %s", b.id, r.URL.String())
	})

	extensions.RandomUserAgent(c)

	return c
}

// Scrap scraps the websites
func (b *baseScrapper) Scrap(wg *sync.WaitGroup) {
	for _, url := range b.urls {
		b.processURL(url, b.getCollector(), b.repo)
	}
	wg.Done()
}
