package scrappers

import (
	"github.com/gocolly/colly"
)

const kijijiID = "kijiji"

func processKijijiURL(url string, c *colly.Collector, repo serviceRepository) {
	c.OnHTML("div#ViewItemPage", func(e *colly.HTMLElement) {
		service := repo.NewService(
			e.Request.URL.String(),
			kijijiID,
			e.ChildText("h1[itemProp=name]"),
		)
		service.Description = e.ChildText("div[itemProp=description]")
		service.Country = "Canada"
		service.City = e.ChildText("ul[class*=crumbList] li:first-child")
		service.Address = e.ChildText("span[itemProp=address]")
		repo.AppendImage(e.ChildAttr("img[itemProp=image]", "src"), service)
		repo.CreateService(service)
	})
	visitLinkNotVisitedBySelector("a.title", c, repo)
	visitLinkBySelector("div.pagination a[title=Next]", c)
	_ = c.Visit(url)
	c.Wait()
}

// Kijiji is the kijiji scrapper.
type Kijiji struct {
	baseScrapper
}

// NewKijiji creates a new Kijiji instance.
func NewKijiji(log logger, repo serviceRepository) *Kijiji {
	config := NewConfig()
	return &Kijiji{baseScrapper: baseScrapper{
		id:             kijijiID,
		logger:         log,
		repo:           repo,
		urls:           config.kijijiURLs,
		processURL:     processKijijiURL,
		allowedDomains: []string{"www.kijiji.ca", "kijiji.ca"},
	}}
}
