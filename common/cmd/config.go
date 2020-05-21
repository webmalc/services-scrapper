package cmd

import (
	"github.com/spf13/viper"
)

// Config is the cmd configuration struct.
type Config struct {
	scrappers []string
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		scrappers: viper.GetStringSlice("scrappers"),
	}
	return config
}
