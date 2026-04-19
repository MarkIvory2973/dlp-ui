package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func New(name string, level logrus.Level) (*logrus.Logger, error) {
	file, err := os.Create(fmt.Sprintf("./%s.log", name))
	if err != nil {
		return nil, err
	}

	formatter := logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000 Z07:00",
	}

	logger := &logrus.Logger{
		Out:          file,
		Formatter:    &formatter,
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	return logger, nil
}
