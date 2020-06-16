package scrappers

import (
	"github.com/gocolly/colly"
)

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
