package scrappers

// ScrappersRunner is the scrappers runner.
type Runner struct {
	logger Logger
}

// Run run the scrappers.
func (r *Runner) Run(names []string) {
	r.logger.Infof("Start the scrappers: %v", names)
}

// NewRunner creates a new Runner instance.
func NewRunner(log Logger) *Runner {
	return &Runner{
		logger: log,
	}
}
