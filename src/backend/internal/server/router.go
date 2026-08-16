package server

import (
	"dlp-ui/cmd"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(mode string) *gin.Engine {
	gin.SetMode(mode)

	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())

	if cmd.GetDebug() {
		router.Use(cors.Default())
	}

	return router
}
