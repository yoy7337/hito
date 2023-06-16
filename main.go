package main

import (
	"hito/cmd"
	"hito/configs"
	"hito/models"

	log "github.com/sirupsen/logrus"
)

func main() {
	// init models
	models.Init()

	// general config
	generalConfig()
	// setup command
	cmd.Execute()
}

func generalConfig() {
	// setup logrus
	log.SetLevel(log.Level(configs.GeneralConf.GetInt("logrus.Level")))
}
