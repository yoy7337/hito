package configs

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var GeneralConf *viper.Viper

func init() {
	GeneralConf = viper.New()
	GeneralConf.SetConfigType("yaml")
	GeneralConf.SetConfigName("generalConf")

	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		configPath = "./configs/"
	}
	log.Infof("Load config files from %s", configPath)
	GeneralConf.AddConfigPath(configPath)
	if err := GeneralConf.ReadInConfig(); err != nil {
		// handle errors
		log.Fatalf("Can not find config file: %s", err.Error())
	}

	setDefault()
}

func setDefault() {

}
