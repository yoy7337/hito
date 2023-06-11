package main

import (
	"hito/cmd"
	"hito/configs"

	log "github.com/sirupsen/logrus"
)

func main() {
	generalConfig()
	// setup command
	cmd.Execute()
}

func generalConfig() {
	// setup logrus
	log.SetLevel(log.Level(configs.GeneralConf.GetInt("logrus.Level")))
}
