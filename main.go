//+build !test

package main

import (
	"github.com/webmalc/services-scrapper/cmd"
	"github.com/webmalc/services-scrapper/common/config"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/common/logger"
	"github.com/webmalc/services-scrapper/models"
	"github.com/webmalc/services-scrapper/scrappers"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	defer conn.Close()
	models.Migrate(conn)
	repo := models.NewServiceRepository(conn)
	cmdRouter := cmd.NewCommandRouter(log, scrappers.NewRunner(log, repo))
	cmdRouter.Run()
}
