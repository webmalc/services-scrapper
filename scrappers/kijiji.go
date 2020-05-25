package scrappers

import (
	"github.com/gocolly/colly"
)

const kijijiID = "kijiji"

func processKijijiURL(url string, c *colly.Collector) {
	/* complete the scrapping
	c.OnHTML("a.title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		serviceURL := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Println(serviceURL)
		e.Request.Visit(serviceURL)
	})
	c.OnHTML("div[itemProp=description]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML("div.pagination a[title=Next]", func(e *colly.HTMLElement) {
		fmt.Println("go to --------------------------------------------------")
		pageURL := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Println(pageURL)
		e.Request.Visit(pageURL)
	})
	c.Visit(url)
	c.Wait()
	*/
}

// Kijiji is the kijiji scrapper.
type Kijiji struct {
	baseScrapper
}

// NewKijiji creates a new Kijiji instance.
func NewKijiji(log Logger) *Kijiji {
	config := NewConfig()
	return &Kijiji{baseScrapper: baseScrapper{
		id:             kijijiID,
		logger:         log,
		urls:           config.kijijiURLs,
		processURL:     processKijijiURL,
		allowedDomains: []string{"www.kijiji.ca", "kijiji.ca"},
	}}
}
