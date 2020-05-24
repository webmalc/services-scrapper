package scrappers

import "sync"

// Logger logs the information
type Logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

// Scrapper scraps websites
type Scrapper interface {
	Scrap(wg *sync.WaitGroup)
}
