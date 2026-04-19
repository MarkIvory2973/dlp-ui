package ytdlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDownloader(t *testing.T) {
	logger := setup(t)

	tests := []struct {
		Name string
		Case string
	}{
		{"Unsupported", "https://TestNewDownloader/Unsupported"},
		{"Invalid", "TestNewDownloader/Invalid"},
		{"Empty", ""},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			downloads := []Download{{URL: test.Case}}

			downloader, err := NewDownloader("", test.Case, "worst", downloads)
			require.NoError(t, err)

			downloader(logger)

			require.NotEmpty(t, downloads)

			for _, download := range downloads {
				require.NotEmpty(t, download.Errors)
			}
		})
	}
}
