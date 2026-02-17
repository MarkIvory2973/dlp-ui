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

func TestParser(t *testing.T) {
	cases := [3]string{
		"https://y.qq.com/n/ryqq/albumDetail/001Ez1U10TYoXS",
		"https://an.unsupported.url/to/parse",
		"test",
	}

	t.Chdir("../..")

	logger, err := log.New("test")
	require.NoError(t, err)
	defer os.Remove("test.log")

	require.FileExists(t, "test.log")

	var parseds sync.Map

	for index := range 3 {
		t.Run(cases[index], func(t *testing.T) {
			logger := logger.WithFields(logrus.Fields{
				"trace": "test",
			})

			parser, err := Parser(cases[index], &parseds)
			require.NoError(t, err)

			parser(logger)

			require.NotEmpty(t, utils.Stom(&parseds))
		})
	}

	raw, err := os.ReadFile("test.log")
	require.NoError(t, err)

	require.NotEmpty(t, raw)
}
