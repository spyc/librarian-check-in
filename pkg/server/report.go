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

type overviewRow struct {
	Pycid    string  `json:"pycid"`
	Name     string  `json:"name"`
	FairTime float64 `json:"fair_time"`
	GoodTime float64 `json:"good_time"`
	FailTime float64 `json:"fail_time"`
}

type fulfilledRow struct {
	librarian *modals.Librarian
	record    *modals.Attendance
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
	router.HandlerFunc("GET", "/report/overview", h.overviewReport)
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

func (h *ReportHandler) overviewReport(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	h.Logger.Debug(query)
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")
	if startDate == "" || endDate == "" {
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

	records, err := h.RecordStore.GetAll()
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when loading", http.StatusInternalServerError)
		return
	}

	fulfilled := make(map[string]overviewRow, len(records))

	for _, record := range records {
		if startTime.Before(*record.CheckIn) && endTime.After(*record.CheckIn) {
			l, err := h.LibrarianStore.GetByID(record.Pycid)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			row, found := fulfilled[l.Pycid]
			if !found {
				row = overviewRow{
					Pycid:    l.Pycid,
					Name:     l.Name,
					GoodTime: 0,
					FailTime: 0,
					FairTime: 0,
				}
			}

			minutes := record.CheckOut.Sub(*record.CheckIn).Minutes()
			switch record.Rank {
			case modals.Rank_fail:
				row.FailTime += 0.5 * minutes
				break
			case modals.Rank_fair:
				row.FairTime += 1.0 * minutes
				break
			case modals.Rank_good:
				row.GoodTime += 2 * minutes
				break
			}

			fulfilled[l.Pycid] = row
		}
	}

	rows := make([]overviewRow, 0, len(fulfilled))

	for _, v := range fulfilled {
		rows = append(rows, v)
	}

	body, err := json.Marshal(rows)
	if err != nil {
		h.Logger.Error(err)
		http.Error(w, "error when encoding", http.StatusInternalServerError)
		return
	}
	h.Logger.Debugln(string(body))
	w.Write(body)
}
