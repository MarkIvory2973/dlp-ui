package web

import (
	"dlp-ui/internal/web/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func New(logger *logrus.Logger) *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies(nil)

	router.Use(middlewares.Logger(logger))
	router.Use(gin.Recovery())

	return router
}
