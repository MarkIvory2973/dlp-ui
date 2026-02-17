package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListIPs(t *testing.T) {
	ips, err := ListIPs()
	require.NoError(t, err)

	require.NotEmpty(t, ips)
}
