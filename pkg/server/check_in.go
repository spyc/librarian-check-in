package server

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"library.pyc.edu.hk/attendance/pkg/modals"
	"sync"
)

type CheckInHandler struct {
	Store  *modals.AttendanceStore `inject:""`
	Logger *logrus.Entry           `inject:"api logger"`

	router    *httprouter.Router
	bootstrap sync.Once
}

func (h *CheckInHandler) init() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/checkin", h.getAll)
	router.HandlerFunc("POST", "/checkin/:pycid", h.checkIn)
	router.HandlerFunc("POST", "/checkout/:pycid", h.checkOut)
	router.HandlerFunc("POST", "/checkout", h.checkOutAll)
	h.router = router
}

func (h *CheckInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.bootstrap.Do(h.init)
	h.router.ServeHTTP(w, r)
}

func (h *CheckInHandler) getAll(w http.ResponseWriter, r *http.Request) {
	results, err := h.Store.GetCheckedIn()
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when loading", http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		w.Write([]byte("[]"))
		return
	}

	body, err := json.Marshal(results)
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when encoding", http.StatusInternalServerError)
		return
	}

	w.Write(body)
}

func (h *CheckInHandler) checkIn(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("pycid")
	if err := h.Store.CheckIn(id); err != nil {
		var status int
		switch err {
		case modals.ErrNotFound:
			status = http.StatusNotFound
			break
		case modals.ErrMultipleCheckIn:
			status = http.StatusConflict
			break
		default:
			status = http.StatusInternalServerError
			h.Logger.Error(err)
			break
		}
		w.WriteHeader(status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CheckInHandler) checkOut(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("pycid")
	if err := h.Store.CheckOut(id); err != nil {
		var status int
		switch err {
		case modals.ErrRecordNotFound:
			status = http.StatusNotFound
			break
		default:
			status = http.StatusInternalServerError
			h.Logger.Error(err)
			break
		}
		w.WriteHeader(status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CheckInHandler) checkOutAll(w http.ResponseWriter, r *http.Request) {
	if err := h.Store.CheckOutAll(); err != nil {
		var status int
		switch err {
		case modals.ErrRecordNotFound:
			status = http.StatusNotFound
			break
		default:
			status = http.StatusInternalServerError
			h.Logger.Error(err)
			break
		}
		w.WriteHeader(status)
		return
	}

	w.WriteHeader(http.StatusOK)
}
