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
		{"Invalid", "invalid", false},
		{"Empty", " ", false},
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

func TestGetBrowser(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want string
	}{
		{"Brave", "brave", "brave"},
		{"Chrome", "chrome", "chrome"},
		{"Chromium", "chromium", "chromium"},
		{"Edge", "edge", "edge"},
		{"Firefox", "firefox", "firefox"},
		{"Opera", "opera", "opera"},
		{"Safari", "safari", "safari"},
		{"Vivaldi", "vivaldi", "vivaldi"},
		{"Whale", "whale", "whale"},
		{"Upper", "BravE", "brave"},
		{"Invalid", "invalid", ""},
		{"Empty", " ", ""},
		{"None", "", ""},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Setenv("BROWSER", test.Case)

			got := GetBrowser()
			require.Equal(t, test.Want, got)
		})
	}
}
