package bootstrap

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func initLogger() *logrus.Logger {
	logger = logrus.New()
	if os.Getenv("VERBOSE") != "" {
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}

func GetLogger() *logrus.Logger {
	return logger
}
