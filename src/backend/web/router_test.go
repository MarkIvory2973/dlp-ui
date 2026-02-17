package web

import (
	"dlp-ui/log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetEnvMode(t *testing.T) {
	cases := [3]string{
		"debug",
		"release",
		"test",
	}
	wants := [3]string{
		gin.DebugMode,
		gin.ReleaseMode,
		gin.TestMode,
	}

	for index := range 3 {
		t.Run(cases[index], func(t *testing.T) {
			t.Setenv("MODE", cases[index])

			got := getEnvMode()
			require.Equal(t, wants[index], got)
		})
	}

	t.Run("default", func(t *testing.T) {
		got := getEnvMode()
		require.Equal(t, gin.ReleaseMode, got)
	})
}

func TestNew(t *testing.T) {
	t.Chdir("..")

	logger, err := log.New("test")
	require.NoError(t, err)
	defer os.Remove("test.log")

	require.FileExists(t, "test.log")

	New(logger)
}
