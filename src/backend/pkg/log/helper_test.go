package log

import "testing"

func setup(t *testing.T) {
	path := t.TempDir()
	t.Chdir(path)
}
