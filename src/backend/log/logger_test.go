package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetEnvLevel(t *testing.T) {
	cases := [7]string{
		"panic",
		"fatal",
		"error",
		"warn",
		"info",
		"debug",
		"trace",
	}
	wants := [7]logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}

	for index := range 7 {
		t.Run(cases[index], func(t *testing.T) {
			t.Setenv("LOG_LEVEL", cases[index])

			got := getEnvLevel()
			require.Equal(t, wants[index], got)
		})
	}

	t.Run("default", func(t *testing.T) {
		got := getEnvLevel()
		require.Equal(t, logrus.InfoLevel, got)
	})
}

func TestNew(t *testing.T) {
	t.Chdir("..")

	logger, err := New("test")
	require.NoError(t, err)
	defer os.Remove("test.log")

	require.FileExists(t, "test.log")

	logger.Info("test")

	raw, err := os.ReadFile("test.log")
	require.NoError(t, err)

	require.NotEmpty(t, raw)
}
