package middleware

import (
	"dlp-ui/log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Chdir("../..")

	logger, err := log.New("test")
	require.NoError(t, err)
	defer os.Remove("test.log")

	require.FileExists(t, "test.log")

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.SetTrustedProxies(nil)
	router.Use(Logger(logger))
	router.Use(gin.Recovery())

	router.GET("/test", func(context *gin.Context) {
		logger := context.MustGet("logger").(*logrus.Entry)

		logger.Info("test")

		context.Status(http.StatusOK)
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)

	raw, err := os.ReadFile("test.log")
	require.NoError(t, err)

	require.NotEmpty(t, raw)
}
