package routers

import (
	"hito/configs"
	v1 "hito/routers/apis/v1"

	middleware "hito/routers/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	// setup router
	router := gin.New()
	router.Use(cors.New(corsConfig()))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// setup error handler
	router.Use(middleware.ErrorHandler())

	// setup validator
	initValidator()

	publicV1R := router.Group(configs.GeneralConf.GetString("app.api.prefix")).Group("/v1")
	privateV1R := router.Group(configs.GeneralConf.GetString("app.api.prefix")).Group("/v1")
	privateV1R.Use(middleware.AuthJWT)

	v1.MountUserApis(publicV1R.Group("/user"), privateV1R.Group("/user"))

	return router
}

func corsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "Content-Length", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host"}

	return config
}
