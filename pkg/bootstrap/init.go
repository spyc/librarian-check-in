package bootstrap

import (
	"net/http"

	"github.com/facebookgo/inject"
	"library.pyc.edu.hk/attendance/pkg/server"
)

var handler server.Handler

type injectFunc func(*inject.Graph) error

func init() {
	var graph inject.Graph
	graph.Logger = initLogger()

	if err := graph.Provide(&inject.Object{Value: &handler}); err != nil {
		panic(err)
	}

	for _, fn := range []injectFunc{injectLogger} {
		if err := fn(&graph); err != nil {
			logger.Error(err)
			panic(err)
		}
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
}

func GetHandler() http.Handler {
	return &handler
}
