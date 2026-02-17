package ytdlp

import (
	"dlp-ui/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Chdir("../..")

	require.DirExists(t, "bin")

	command, stdout, stderr, err := new("--version")
	require.NoError(t, err)

	err = command.Start()
	require.NoError(t, err)

	outputs, err := utils.ReadLines(stdout)
	require.NoError(t, err)

	errors, err := utils.ReadLines(stderr)
	require.NoError(t, err)

	require.NotEmpty(t, outputs)
	require.Empty(t, errors)
}
