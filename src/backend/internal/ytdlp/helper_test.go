package ytdlp

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func setup(t *testing.T) *logrus.Entry {
	path := t.TempDir()

	err := os.CopyFS(path, os.DirFS("testdata"))
	if err != nil {
		t.Fatal(err)
	}

	t.Chdir(path)

	logger := &logrus.Entry{
		Logger: &logrus.Logger{},
	}

	return logger
}
