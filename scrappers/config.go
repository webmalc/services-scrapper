package scrappers

import (
	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	kijijiURLs []string
	yandexURLs []string
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("kijijiURLs", []string{
		"https://www.kijiji.ca/b-services/ontario/c72l9004",
		"https://www.kijiji.ca/b-services/quebec/c72l9001",
	})
	viper.SetDefault("yandexURLs", []string{""})
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		kijijiURLs: viper.GetStringSlice("kijijiURLs"),
		yandexURLs: viper.GetStringSlice("yandexURLs"),
	}
	return config
}
