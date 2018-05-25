package bootstrap

import (
	"os"

	"fmt"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
	"time"
)

var logger *logrus.Logger

func initLogger() *logrus.Logger {
	logger = logrus.New()

	filename := fmt.Sprintf("log/%s.log", time.Now().Format(time.RFC3339))
	out, _ := os.Create(filename)
	logger.Out = out

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
		&inject.Object{Value: logger.WithField("source", "checkin"), Name: "checkin logger"},
		&inject.Object{Value: logger.WithField("source", "librarian"), Name: "librarian logger"},
		&inject.Object{Value: logger.WithField("source", "auth"), Name: "auth logger"},
	)
}
