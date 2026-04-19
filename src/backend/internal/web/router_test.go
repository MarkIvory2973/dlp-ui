package web

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

func TestNew(t *testing.T) {
	logger := setup()

	router := New(logger)

	router.GET("/TestNew", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	apitest.New().
		Handler(router).
		Get("/TestNew").
		Expect(t).
		Status(http.StatusOK).
		End()
}
