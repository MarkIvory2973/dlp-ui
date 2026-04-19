package ytdlp

import (
	"dlp-ui/pkg/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	setup(t)

	args := []string{
		"--version",
	}
	command, stdout, stderr, err := new(args...)
	require.NoError(t, err)

	err = command.Start()
	require.NoError(t, err)

	t.Run("Stdout", func(t *testing.T) {
		var contents []string

		utils.ScanLineFunc(stdout, func(content string) {
			contents = append(contents, content)
		})

		require.NotEmpty(t, contents)
	})

	t.Run("Stderr", func(t *testing.T) {
		var contents []string

		utils.ScanLineFunc(stderr, func(content string) {
			contents = append(contents, content)
		})

		require.Empty(t, contents)
	})
}
