package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		// generate a new uuid for each request
		trace := uuid.NewString()
		// attach it to the logger
		logger := logger.WithFields(logrus.Fields{
			"trace": trace,
		})
		// attach it to the handler
		context.Set("logger", logger)
		// attach it to the response
		context.Header("X-Trace", trace)

		// measure handler duration
		start := time.Now()
		context.Next()
		end := time.Now()

		// fields of the logger
		status := context.Writer.Status()
		duration := end.Sub(start).String()
		method := context.Request.Method
		path := context.Request.URL.Path
		ip := context.ClientIP()
		query := context.Request.URL.Query()

		logger = logger.WithFields(logrus.Fields{
			"status":   status,
			"duration": duration,
			"method":   method,
			"path":     path,
			"query":    query,
			"ip":       ip,
		})

		switch {
		// 5xx
		case 500 <= status:
			logger.Error("http_request")
		// 4xx
		case 400 <= status && status < 500:
			logger.Warn("http_request")
		// 2xx or 3xx
		default:
			logger.Info("http_request")
		}
	}
}
