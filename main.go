//+build !test

package main

import (
	"github.com/webmalc/services-scrapper/common/cmd"
	"github.com/webmalc/services-scrapper/common/config"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/common/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	defer conn.Close()
	cmdRouter := cmd.NewCommandRouter(log)
	cmdRouter.Run()
}
