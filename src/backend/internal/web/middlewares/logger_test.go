package middlewares

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/steinfletcher/apitest"
)

func TestLogger(t *testing.T) {
	router, logger := setup()

	router.Use(Logger(logger))

	router.GET("/TestLogger/Trace", func(context *gin.Context) {
		trace := context.Writer.Header().Get("X-Trace")
		if trace == "" {
			context.Status(http.StatusInternalServerError)
			return
		}

		context.Status(http.StatusOK)
	})

	router.GET("/TestLogger/Logger", func(context *gin.Context) {
		logger, exists := context.Get("logger")
		if !exists {
			context.Status(http.StatusInternalServerError)
			return
		}

		_, ok := logger.(*logrus.Entry)
		if !ok {
			context.Status(http.StatusInternalServerError)
			return
		}

		context.Status(http.StatusOK)
	})

	t.Run("Trace", func(t *testing.T) {
		apitest.New().
			Handler(router).
			Get("/TestLogger/Trace").
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Logger", func(t *testing.T) {
		apitest.New().
			Handler(router).
			Get("/TestLogger/Logger").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
