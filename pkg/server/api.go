package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"library.pyc.edu.hk/attendance/pkg/modals"
)

type ApiHandler struct {
	LibrarianStore *modals.LibrarianStore `inject:""`
	Logger         *logrus.Entry          `inject:"api logger"`

	handler   http.Handler
	bootstrap sync.Once
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.bootstrap.Do(func() {
		router := httprouter.New()
		h.bindLibrarianRoute(router)

		handler := handlers.CombinedLoggingHandler(h.Logger.Writer(), router)
		handler = handlers.RecoveryHandler(
			handlers.RecoveryLogger(h.Logger.WithField("source", "recover")),
			handlers.PrintRecoveryStack(true),
		)(handler)

		h.handler = handler
	})
	h.handler.ServeHTTP(w, r)
}
