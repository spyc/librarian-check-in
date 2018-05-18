package server

import (
	"net/http"

	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"library.pyc.edu.hk/attendance/pkg/modals"
	"sync"
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
		router.HandlerFunc("GET", "/librarian", h.getLibrarians)

		handler := handlers.CombinedLoggingHandler(h.Logger.Writer(), router)
		handler = handlers.RecoveryHandler(
			handlers.RecoveryLogger(h.Logger.WithField("source", "recover")),
			handlers.PrintRecoveryStack(true),
		)(handler)

		h.handler = handler
	})
	h.handler.ServeHTTP(w, r)
}

func (h *ApiHandler) getLibrarians(w http.ResponseWriter, r *http.Request) {
	librarians, err := h.LibrarianStore.GetAll()
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when loading", http.StatusInternalServerError)
		return
	}

	if len(librarians) == 0 {
		w.Write([]byte("[]"))
		return
	}

	body, err := json.Marshal(librarians)
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when encoding", http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
