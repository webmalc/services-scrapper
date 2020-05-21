package cmd

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}
