package bootstrap

import (
	"os"

	"github.com/facebookgo/inject"
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

func injectLogger(graph *inject.Graph) error {
	return graph.Provide(
		&inject.Object{Value: logger.WithField("source", "api"), Name: "api logger"},
	)
}
