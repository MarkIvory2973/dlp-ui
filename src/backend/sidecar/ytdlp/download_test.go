package ytdlp

import (
	"dlp-ui/log"
	"dlp-ui/utils"
	"os"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestDownloader(t *testing.T) {
	cases := [3][2]string{
		{"https://y.qq.com/n/ryqq/songDetail/002rhFKO3EjKAg", "flac"},
		{"https://an.unsupported.url/to/parse", "best"},
		{"test", "test"},
	}

	t.Chdir("../..")

	logger, err := log.New("test")
	require.NoError(t, err)
	defer os.Remove("test.log")

	require.FileExists(t, "test.log")

	var downloads sync.Map

	for index := range 3 {
		t.Run(cases[index][0], func(t *testing.T) {
			logger := logger.WithFields(logrus.Fields{
				"trace": "test",
			})

			downloader, err := Downloader(cases[index][0], cases[index][1], &downloads)
			require.NoError(t, err)

			downloader(logger)

			require.NotEmpty(t, utils.Stom(&downloads))
		})
	}
	defer os.RemoveAll("down")

	require.DirExists(t, "down")

	entries, err := os.ReadDir("down")
	require.NoError(t, err)

	require.NotEmpty(t, entries)

	for _, entry := range entries {
		raw, err := os.ReadFile("down/" + entry.Name())
		require.NoError(t, err)

		require.NotEmpty(t, raw)
	}

	raw, err := os.ReadFile("test.log")
	require.NoError(t, err)

	require.NotEmpty(t, raw)
}
