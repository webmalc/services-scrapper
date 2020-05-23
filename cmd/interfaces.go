package cmd

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// Runner runs the command
type Runner interface {
	Run(names []string)
}
