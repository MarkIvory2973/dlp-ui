package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDebug(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want bool
	}{
		{"True", "true", true},
		{"False", "false", false},
		{"Upper", "TruE", true},
		{"None", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Setenv("DEBUG", test.Case)

			got := GetDebug()
			require.Equal(t, test.Want, got)
		})
	}
}
