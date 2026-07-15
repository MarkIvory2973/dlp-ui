package server

import (
	"dlp-ui/cmd"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	mode := cmd.GetMode()
	gin.SetMode(mode)
}

func New() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())

	mode := cmd.GetMode()
	if mode != gin.ReleaseMode {
		router.Use(cors.Default())
	}

	return router
}
