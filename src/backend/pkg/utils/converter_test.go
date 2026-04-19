package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want int
	}{
		{"Normal", "1", 1},
		{"Invalid", "Invalid", -1},
		{"Empty", "", -1},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := Atoi(test.Case, -1)
			require.Equal(t, test.Want, got)
		})
	}
}
