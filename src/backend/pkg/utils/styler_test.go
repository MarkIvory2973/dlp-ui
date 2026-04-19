package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyle(t *testing.T) {
	tests := []struct {
		Name string
		Case []string
		Want string
	}{
		{"Mode", []string{Strong}, "\033[1mtest\033[0m"},
		{"Color", []string{Green}, "\033[32mtest\033[0m"},
		{"Empty", []string{}, "test\033[0m"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := Style("test", test.Case...)
			require.Equal(t, test.Want, got)
		})
	}
}
