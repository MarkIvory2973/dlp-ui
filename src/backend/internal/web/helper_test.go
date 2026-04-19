package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func setup() *logrus.Logger {
	logger := &logrus.Logger{}

	gin.SetMode(gin.TestMode)

	return logger
}
