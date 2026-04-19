package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func setup() (*gin.Engine, *logrus.Logger) {
	logger := &logrus.Logger{}

	gin.SetMode(gin.TestMode)

	router := gin.New()

	return router, logger
}
