package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"

	"library.pyc.edu.hk/attendance/pkg/modals"
)

type ApiHandler struct {
	LibrarianStore   *modals.LibrarianStore `inject:""`
	Logger           *logrus.Entry          `inject:"api logger"`
	CheckInHandler   *CheckInHandler        `inject:""`
	LibrarianHandler *LibrarianHandler      `inject:""`
	ReportHandler    *ReportHandler         `inject:""`
	AuthHandler      *AuthHandler           `inject:""`

	handler   http.Handler
	bootstrap sync.Once
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.bootstrap.Do(func() {
		mux := http.NewServeMux()
		mux.Handle("/checkin", h.CheckInHandler)
		mux.Handle("/checkin/", h.CheckInHandler)
		mux.Handle("/checkout", h.CheckInHandler)
		mux.Handle("/checkout/", h.CheckInHandler)

		mux.Handle("/librarian", h.LibrarianHandler)
		mux.Handle("/librarian/", h.LibrarianHandler)

		mux.Handle("/report/", h.ReportHandler)

		mux.Handle("/auth", h.AuthHandler)

		handler := handlers.CombinedLoggingHandler(h.Logger.Writer(), mux)
		handler = handlers.RecoveryHandler(
			handlers.RecoveryLogger(h.Logger.WithField("source", "recover")),
			handlers.PrintRecoveryStack(true),
		)(handler)

		h.handler = handler
	})
	h.handler.ServeHTTP(w, r)
}
