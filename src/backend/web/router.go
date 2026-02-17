package web

import (
	"dlp-ui/web/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getEnvMode() string {
	// get mode from env
	mode := os.Getenv("MODE")

	switch mode {
	// debug mode
	case "debug":
		return gin.DebugMode
	// release mode
	case "release":
		return gin.ReleaseMode
	// test mode
	case "test":
		return gin.TestMode
	// release mode (default)
	default:
		return gin.ReleaseMode
	}
}

func New(logger *logrus.Logger) *gin.Engine {
	// set mode
	mode := getEnvMode()
	gin.SetMode(mode)

	// create a new router
	router := gin.New()

	// use cors if the mode is not release
	if mode != gin.ReleaseMode {
		router.Use(cors.Default())
	}

	// configs
	router.SetTrustedProxies(nil)
	// middlewares
	router.Use(middleware.Logger(logger))
	router.Use(gin.Recovery())

	return router
}
