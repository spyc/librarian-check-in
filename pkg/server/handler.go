package server

import (
	"net/http"
	"os"
	"path"
	"sync"
)

type Handler struct {
	Mux *http.ServeMux `inject:""`
	Api *ApiHandler    `inject:""`

	bootstrap sync.Once
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.bootstrap.Do(func() {
		mux := h.Mux

		root, _ := os.Getwd()
		dir := http.Dir(path.Join(root, "ui"))
		mux.Handle("/", http.FileServer(dir))
		mux.Handle("/api/", http.StripPrefix("/api", h.Api))
	})

	h.Mux.ServeHTTP(w, r)
}
