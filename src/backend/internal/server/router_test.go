package server

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

func TestNew(t *testing.T) {
	router := New(gin.TestMode)

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	apitest.New().
		Handler(router).
		Get("/ping").
		Expect(t).
		Status(http.StatusOK).
		Body("pong").
		End()
}
