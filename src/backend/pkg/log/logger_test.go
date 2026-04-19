package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	setup(t)

	logger, err := New("test", logrus.DebugLevel)
	require.NoError(t, err)

	require.FileExists(t, "./test.log")

	logger.Debug("test")

	content, err := os.ReadFile("./test.log")
	require.NoError(t, err)

	require.NotEmpty(t, content)
}
