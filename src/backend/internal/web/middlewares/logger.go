package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		trace := uuid.NewString()
		logger := logger.WithField("trace", trace)

		context.Header("X-Trace", trace)
		context.Set("logger", logger)

		context.Next()
	}
}
