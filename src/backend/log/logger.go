package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var formatter = &logrus.JSONFormatter{
	// 2006-01-02 15:04:05.000 +07:00
	TimestampFormat: "2006-01-02 15:04:05.000 Z07:00",
}

func getEnvLevel() logrus.Level {
	// get level from env
	level := os.Getenv("LOG_LEVEL")

	switch level {
	// panic level
	case "panic":
		return logrus.PanicLevel
	// fatal level
	case "fatal":
		return logrus.FatalLevel
	// error level
	case "error":
		return logrus.ErrorLevel
	// warn level
	case "warn":
		return logrus.WarnLevel
	// info level
	case "info":
		return logrus.InfoLevel
	// debug level
	case "debug":
		return logrus.DebugLevel
	// trace level
	case "trace":
		return logrus.TraceLevel
	// info level (default)
	default:
		return logrus.InfoLevel
	}
}

func New(name string) (*logrus.Logger, error) {
	// create a file
	path := fmt.Sprintf("%s.log", name)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	// get level from env
	level := getEnvLevel()

	// create a new logger
	logger := logrus.New()

	// configs
	logger.SetOutput(file)
	logger.SetFormatter(formatter)
	logger.SetLevel(level)

	return logger, nil
}
