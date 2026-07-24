package server

import (
	"dlp-ui/cmd"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	if cmd.GetDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func New() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())

	if cmd.GetDebug() {
		router.Use(cors.Default())
	}

	return router
}
