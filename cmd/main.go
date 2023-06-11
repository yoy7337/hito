package cmd

import (
	"fmt"
	"hito/configs"
	"hito/routers"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "hito",
	Short: "run hitto backend service",
	Long:  "run hitto backend service",
	Run: func(_ *cobra.Command, _ []string) {

		// setup router
		router := routers.InitRouters()

		// setup http server
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", configs.GeneralConf.GetInt("app.api.port")),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		log.Infof("hix-server started")

		s.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(mainCmd)
}
