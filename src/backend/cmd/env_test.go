package cmd

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetMode(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want string
	}{
		{"Debug", "debug", gin.DebugMode},
		{"Release", "release", gin.ReleaseMode},
		{"Test", "test", gin.TestMode},
		{"Upper", "Debug", gin.DebugMode},
		{"Default", "", gin.ReleaseMode},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Setenv("MODE", test.Case)

			got := GetMode()
			require.Equal(t, test.Want, got)
		})
	}
}
