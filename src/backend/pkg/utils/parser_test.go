package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBrowser(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want string
	}{
		{"Chrome", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Safari/537.36", "chrome"},
		{"Edge", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/103.0.1264.71", "edge"},
		{"Firefox", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Gecko/20100101 Firefox/102.0", "firefox"},
		{"Opera", "Mozilla/5.0 (Windows NT 11.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 OPR/89.0.4447.51", "opera"},
		{"Safari", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15", "safari"},
		{"Unsupported", "curl/8.14.1", ""},
		{"Invalid", "invalid", ""},
		{"Empty", "", ""},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got, err := ParseBrowser(test.Case)
			require.NoError(t, err)

			require.Equal(t, test.Want, got)
		})
	}
}
