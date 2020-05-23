package scrappers

// Logger logs the information
type Logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}
