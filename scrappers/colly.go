package scrappers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

// requestJSON get JSON by the URL
func requestJSON(url string, target interface{}, client *http.Client) error {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// visitLinkBySelector visits the link by selector
func visitLinkBySelector(selector string, c *colly.Collector) {
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		_ = e.Request.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})
}

// visitLinkNotVisited visits the link if it has not been visited yet
func visitLinkNotVisitedBySelector(
	selector string,
	c *colly.Collector,
	repo serviceRepository,
) {
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("href"))
		if !repo.IsServiceWithURLExists(url) {
			_ = e.Request.Visit(url)
		}
	})
}
