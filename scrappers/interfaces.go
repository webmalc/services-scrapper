package scrappers

import (
	"sync"

	"github.com/gocolly/colly"
	"github.com/webmalc/services-scrapper/models"
)

// Logger logs the information
type logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

// Scrapper scraps websites
type scrapper interface {
	Scrap(wg *sync.WaitGroup)
}

// ServiceRepository interface
type serviceRepository interface {
	NewService(url, source, title string) *models.Service
	CreateService(service *models.Service) []error
	IsServiceWithURLExists(url string) bool
	AppendImage(src string, service *models.Service)
}

type urlProcessorType func(string, *colly.Collector, serviceRepository)
