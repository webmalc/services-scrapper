package scrappers

import (
	"encoding/json"

	"github.com/gocolly/colly"
)

const yellowpagesID = "yellowpages"
const yellowpagesCountry = "Canada"

// YellowpagesImagesData is the image data struct
type YellowpagesImagesData struct {
	Medias []struct {
		IsPhoto bool   `json:"isPhoto"`
		URL     string `json:"url"`
	}
}

// getYellowpagesImages return a list of images
func getYellowpagesImages(e *colly.HTMLElement) []string {
	result := []string{}
	scriptText := e.DOM.Closest("body").Find(
		"script.jsImagesConfigHandlebar",
	).Text()
	if scriptText == "" {
		return result
	}
	data := &YellowpagesImagesData{}
	err := json.Unmarshal([]byte(scriptText), data)
	if err != nil {
		return result
	}
	for _, d := range data.Medias {
		if d.IsPhoto {
			result = append(result, d.URL)
		}
	}
	return result
}

func processYellowpagesURL(
	url string, c *colly.Collector, repo serviceRepository,
) {
	c.OnHTML("div.merchant__header--root", func(e *colly.HTMLElement) {
		service := repo.NewService(
			e.Request.URL.String(),
			yellowpagesID,
			e.ChildText("h1.merchant__title span[itemprop=name]"),
		)
		service.Country = yellowpagesCountry
		service.Description = e.ChildText("div#businessSection")
		service.City = e.ChildText(
			"div.merchant__details__section--address span[itemprop=addressLocality]",
		)
		service.Address = e.ChildText(
			"div.merchant__details__section--address div.merchant__address",
		)
		avatar := e.ChildAttr("a.merchant-logo-link img", "src")
		if avatar != "" {
			service.Avatar = avatar
		}
		e.ForEach(
			"div.middle_section li.mlr__item--phone span.mlr__sub-text",
			func(i int, el *colly.HTMLElement) {
				repo.AppendPhone(el.Text, service)
			},
		)
		e.ForEach(
			"div.middle_section li.mlr__item--website span.mlr__sub-text",
			func(i int, el *colly.HTMLElement) {
				repo.AppendLink(el.Text, service)
			},
		)
		for _, i := range getYellowpagesImages(e) {
			repo.AppendImage(i, service)
		}
		repo.CreateService(service)
	})
	visitLinkNotVisitedBySelector("a.listing__link", c, repo)
	visitLinkBySelector("span.pageCount + a.pageButton", c)
	_ = c.Visit(url)
	c.Wait()
}

// Yellowpages is the yellowpages scrapper.
type Yellowpages struct {
	baseScrapper
}

// NewYellowpages creates a new yellowpages instance.
func NewYellowpages(log logger, repo serviceRepository) *Yellowpages {
	config := NewConfig()
	return &Yellowpages{baseScrapper: baseScrapper{
		id:             yellowpagesID,
		logger:         log,
		repo:           repo,
		urls:           config.yellowpagesURLs,
		processURL:     processYellowpagesURL,
		allowedDomains: []string{"www.yellowpages.ca", "yellowpages.ca"},
	}}
}
