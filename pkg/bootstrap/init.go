package bootstrap

import (
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/gorilla/handlers"
	"library.pyc.edu.hk/attendance/pkg/server"
)

var handler server.Handler

func init() {
	var graph inject.Graph
	graph.Logger = initLogger()

	if err := graph.Provide(&inject.Object{Value: &handler}); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
}

func GetHandler() (h http.Handler) {
	h = &handler

	h = handlers.CombinedLoggingHandler(logger.Writer(), h)
	h = handlers.RecoveryHandler(
		handlers.RecoveryLogger(logger.WithField("source", "recover")),
		handlers.PrintRecoveryStack(true),
	)(h)

	return
}
