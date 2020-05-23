//+build !test

package main

import (
	"github.com/webmalc/services-scrapper/cmd"
	"github.com/webmalc/services-scrapper/common/config"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/common/logger"
	"github.com/webmalc/services-scrapper/scrappers"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	defer conn.Close()
	cmdRouter := cmd.NewCommandRouter(log, scrappers.NewRunner(log))
	cmdRouter.Run()
}
