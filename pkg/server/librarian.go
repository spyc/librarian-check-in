package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"library.pyc.edu.hk/attendance/pkg/modals"
)

type LibrarianHandler struct {
	LibrarianStore *modals.LibrarianStore `inject:""`
	Logger         *logrus.Entry          `inject:"api logger"`
	Handler        *AuthMiddleware        `inject:""`

	bootstrap sync.Once
}

func (h *LibrarianHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.bootstrap.Do(h.init)
	h.Handler.ServeHTTP(w, r)
}

func (h *LibrarianHandler) init() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/librarian", h.getLibrarians)
	router.HandlerFunc("POST", "/librarian", h.addLibrarian)
	router.HandlerFunc("PUT", "/librarian/:pycid", h.updateLibrarian)
	router.HandlerFunc("DELETE", "/librarian/:pycid", h.deleteLibrarian)
	h.Handler.Handler = router
}

func (h *LibrarianHandler) getLibrarians(w http.ResponseWriter, r *http.Request) {
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

func (h *LibrarianHandler) addLibrarian(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when sending", http.StatusUnprocessableEntity)
		return
	}

	var l modals.Librarian
	if err := json.Unmarshal(body, &l); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when parsing", http.StatusUnprocessableEntity)
		return
	}

	if err := h.LibrarianStore.Add(&l); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when processing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LibrarianHandler) updateLibrarian(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("pycid")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when sending", http.StatusUnprocessableEntity)
		return
	}

	var l modals.Librarian
	if err := json.Unmarshal(body, &l); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when parsing", http.StatusUnprocessableEntity)
		return
	}

	if err := h.LibrarianStore.Update(id, &l); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when processing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LibrarianHandler) deleteLibrarian(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("pycid")

	if err := h.LibrarianStore.Remove(id); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when processing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
