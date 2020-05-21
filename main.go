//+build !test

package main

import (
	"github.com/webmalc/services-scrapper/common/config"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/common/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	defer conn.Close()
	log.Info("test message")
}
