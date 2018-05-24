package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"library.pyc.edu.hk/attendance/pkg/modals"
)

func (h *ApiHandler) bindLibrarianRoute(router *httprouter.Router) {
	router.HandlerFunc("GET", "/librarian", h.getLibrarians)
	router.HandlerFunc("POST", "/librarian", h.addLibrarian)
	router.HandlerFunc("PUT", "/librarian/:pycid", h.updateLibrarian)
	router.HandlerFunc("DELETE", "/librarian/:pycid", h.deleteLibrarian)
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

func (h *ApiHandler) addLibrarian(w http.ResponseWriter, r *http.Request) {
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

func (h *ApiHandler) updateLibrarian(w http.ResponseWriter, r *http.Request) {
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

func (h *ApiHandler) deleteLibrarian(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("pycid")

	if err := h.LibrarianStore.Remove(id); err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when processing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
