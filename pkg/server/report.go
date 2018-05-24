package server

import (
	"net/http"
	"sync"
	"time"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"library.pyc.edu.hk/attendance/pkg/modals"
)

type personalReport struct {
	Records   []modals.Attendance `json:"records"`
	Librarian *modals.Librarian   `json:"librarian"`
}

type ReportHandler struct {
	RecordStore    *modals.AttendanceStore `inject:""`
	LibrarianStore *modals.LibrarianStore  `inject:""`
	Logger         *logrus.Entry           `inject:"api logger"`

	router    *httprouter.Router
	bootstrap sync.Once
}

func (h *ReportHandler) init() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/report/personal", h.personalReport)
	h.router = router
}

func (h *ReportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.bootstrap.Do(h.init)
	h.router.ServeHTTP(w, r)
}

func (h *ReportHandler) personalReport(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	h.Logger.Debug(query)
	pycid := query.Get("pycid")
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")
	if pycid == "" || startDate == "" || endDate == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		h.Logger.Debug(startDate)
		http.Error(w, "bad startTime in request", http.StatusBadRequest)
		return
	}

	endTime, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		http.Error(w, "bad endTime in request", http.StatusBadRequest)
		return
	}

	l, err := h.LibrarianStore.GetByID(pycid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	records, err := h.RecordStore.GetAll()
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when loading", http.StatusInternalServerError)
		return
	}

	fulfilled := make([]modals.Attendance, 0, len(records))

	for _, record := range records {
		h.Logger.Debug(record.Pycid == pycid, startTime.Before(*record.CheckIn), endTime.After(*record.CheckIn))

		if record.Pycid == pycid && startTime.Before(*record.CheckIn) && endTime.After(*record.CheckIn) {
			fulfilled = append(fulfilled, *record)
		}
	}

	body, err := json.Marshal(&personalReport{
		Librarian: l,
		Records:   fulfilled,
	})
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when encoding", http.StatusInternalServerError)
		return
	}
	h.Logger.Debugln(string(body))
	w.Write(body)
}
