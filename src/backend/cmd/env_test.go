package cmd

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetLogLevel(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want logrus.Level
	}{
		{"Panic", "panic", logrus.PanicLevel},
		{"Fatal", "fatal", logrus.FatalLevel},
		{"Error", "error", logrus.ErrorLevel},
		{"Warn", "warn", logrus.WarnLevel},
		{"Info", "info", logrus.InfoLevel},
		{"Debug", "debug", logrus.DebugLevel},
		{"Trace", "trace", logrus.TraceLevel},
		{"Empty", "", logrus.InfoLevel},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Setenv("LOG_LEVEL", test.Case)

			got := GetLogLevel()
			require.Equal(t, test.Want, got)
		})
	}
}

func TestGetMode(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want string
	}{
		{"Debug", "debug", gin.DebugMode},
		{"Release", "release", gin.ReleaseMode},
		{"Test", "test", gin.TestMode},
		{"Empty", "", gin.ReleaseMode},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Setenv("MODE", test.Case)

			got := GetMode()
			require.Equal(t, test.Want, got)
		})
	}
}
