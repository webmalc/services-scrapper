package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const prefix = "ss"

// getFilename returns filename based on the environment variable
func getFilename() string {
	fileName := "config"
	env := os.Getenv(strings.ToUpper(prefix + "_env"))
	if env != "" {
		fileName += "." + env
	}
	return fileName
}

// setDefaults sets default values
func setDefaults(baseDir string) {
	viper.Set("base_dir", filepath.Dir(filepath.Dir(baseDir))+"/")
	viper.SetDefault("is_prod", false)
	viper.SetDefault("log_path", "logs/app.log")
	viper.SetDefault("scrappers", []string{
		"kijiji", "yandex", "yellowpages", "uslugio",
	})
}

// setPaths set paths
func setPaths(baseDir string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(baseDir + "/yaml")
	viper.AddConfigPath("/etc/ss")
	viper.AddConfigPath("$HOME/.ss")
}

// setEnv sets the environment
func setEnv() {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
}

// getBaseDir get the base directory
func getBaseDir() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("config: unable to determine the caller.")
	}
	return filepath.Dir(b)
}

// read reads configuration
func read() {
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// Setup initializes the main configuration.
func Setup() {
	baseDir := getBaseDir()
	viper.SetConfigName(getFilename())
	setPaths(baseDir)
	setEnv()
	setDefaults(baseDir)
	read()
}
