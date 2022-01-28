package scrappers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/microcosm-cc/bluemonday"
)

const uslugioID = "uslugio"
const delay = 8 * time.Second

type content struct {
	Title  string
	Text   string
	Addres string
}

type phone struct {
	Phone string
}

// getJSONContent gets the conten via JSON request
func getJSONContent(
	id string, e *colly.HTMLElement, httpClient *http.Client,
) (*content, *phone, error) {
	time.Sleep(delay)
	content := &content{}
	phone := &phone{}

	contentURL := e.Request.AbsoluteURL(
		fmt.Sprintf("/?do=show_one&id=%s", id),
	)
	phoneURL := e.Request.AbsoluteURL(
		fmt.Sprintf("/?do=showphone&json=1&h=1&id=%s", id),
	)
	err := requestJSON(contentURL, content, httpClient)
	if err != nil {
		return content, phone, err
	}
	err = requestJSON(phoneURL, phone, httpClient)

	return content, phone, err
}

func processUslugioURL(url string, c *colly.Collector, repo serviceRepository) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	blue := bluemonday.StrictPolicy()
	offset := 50

	c.OnHTML("div.title.showone", func(e *colly.HTMLElement) {
		id := e.Attr("data-id")
		fakeURL := e.Request.AbsoluteURL("/?id=" + id)
		if !repo.IsServiceWithURLExists(fakeURL) {
			content, phone, err := getJSONContent(
				id, e, httpClient,
			)
			if err == nil {
				service := repo.NewService(
					fakeURL,
					uslugioID,
					content.Title,
				)
				service.Description = blue.Sanitize(
					content.Text,
				)
				service.Address = content.Addres
				repo.AppendPhone(phone.Phone, service)
				repo.CreateService(service)
			}
		}
	})

	visitLinkBySelector("a", c)

	c.OnHTML("button.loadmore", func(e *colly.HTMLElement) {
		for i := 1; i < 10; i++ {
			nextURL := e.Request.URL.String() + fmt.Sprintf(
				"?offset=%d", i*offset,
			)
			_ = e.Request.Visit(
				e.Request.AbsoluteURL(nextURL),
			)
		}
	})
	_ = c.Visit(url)
	c.Wait()
}

// Uslugio is the uslugio scrapper.
type Uslugio struct {
	baseScrapper
}

// Uslugio creates a new uslugio instance.
func NewUslugio(log logger, repo serviceRepository) *Uslugio {
	config := NewConfig()
	return &Uslugio{baseScrapper: baseScrapper{
		id:             uslugioID,
		logger:         log,
		repo:           repo,
		urls:           config.uslugioURLs,
		processURL:     processUslugioURL,
		allowedDomains: []string{"www.uslugio.com", "uslugio.com"},
	}}
}
