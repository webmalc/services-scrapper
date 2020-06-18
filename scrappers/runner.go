package scrappers

import (
	"sync"
)

// Runner is the scrappers runner.
type Runner struct {
	logger    logger
	scrappers map[string]scrapper
}

// getScrappers get a list of scrappers to run
func (r *Runner) getScrappers(names []string) []scrapper {
	result := make([]scrapper, 0, len(r.scrappers))
	if len(names) == 0 {
		for _, scrapper := range r.scrappers {
			result = append(result, scrapper)
		}
		return result
	}
	for _, id := range names {
		if scrapper, ok := r.scrappers[id]; ok {
			result = append(result, scrapper)
		}
	}
	return result
}

// Run runs the scrappers.
func (r *Runner) Run(names []string) {
	r.logger.Infof("Start the scrappers: %v", names)

	var wg sync.WaitGroup
	for _, scrapper := range r.getScrappers(names) {
		wg.Add(1)
		go scrapper.Scrap(&wg)
	}
	wg.Wait()
}

// NewRunner creates a new Runner instance.
func NewRunner(log logger, repo serviceRepository) *Runner {
	return &Runner{
		logger: log,
		scrappers: map[string]scrapper{
			kijijiID:      NewKijiji(log, repo),
			yellowpagesID: NewYellowpages(log, repo),
			yandexID:      NewYandex(log),
		},
	}
}
