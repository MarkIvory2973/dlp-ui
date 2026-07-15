package ytdlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	extraArgs := []string{
		"--version",
	}
	context, err := new(extraArgs...)
	require.NoError(t, err)

	err = context.Start()
	require.NoError(t, err)

	err = context.Wait()
	require.NoError(t, err)
}
