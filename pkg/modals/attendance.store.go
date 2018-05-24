package modals

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"sync"
	"time"
)

const AttendanceFile = "attendance.data"

var ErrMultipleCheckIn = errors.New("already check in")
var ErrRecordNotFound = errors.New("check in record not found")

type AttendanceStore struct {
	AbstractStore

	LibrarianStore *LibrarianStore    `inject:""`
	Logger         logrus.FieldLogger `inject:"checkin logger"`

	checkedIn sync.Map
	cached    bool
}

func (s *AttendanceStore) filename() string {
	return path.Join(s.getRootDir(), AttendanceFile)
}

func (s *AttendanceStore) alreadyCheckIn(pycid string) (_ bool, err error) {
	if !s.cached {
		ls, err := s.GetCheckedIn()
		if err != nil {
			return false, err
		}

		for _, l := range ls {
			if l.Pycid == pycid {
				return true, nil
			}
		}

		return false, nil
	}
	_, exists := s.checkedIn.Load(pycid)
	return exists, nil
}

func (s *AttendanceStore) GetCheckedIn() (_ []Librarian, err error) {
	results := make([]Librarian, 0)
	var checkedIn sync.Map
	records, err := s.GetAll()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	for _, record := range records {
		if record.CheckOut == nil {
			l, err := s.LibrarianStore.getByID(record.Pycid)
			if err != nil {
				return nil, err
			}
			checkedIn.Store(record.Pycid, *l)
			results = append(results, *l)
		}
	}

	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.checkedIn = checkedIn

	return results, nil
}

func (s *AttendanceStore) GetAll() (_ []*Attendance, err error) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	records := new(Records)

	if err := records.Unmarshal(data); err != nil {
		return nil, err
	}

	return records.Records, nil
}

func (s *AttendanceStore) CheckIn(pycid string) (err error) {
	exists, err := s.alreadyCheckIn(pycid)
	if err != nil {
		return
	} else if exists {
		return ErrMultipleCheckIn
	}

	l, err := s.LibrarianStore.getByID(pycid)
	if err != nil {
		return err
	} else if l == nil {
		return ErrNotFound
	}

	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	records := new(Records)

	if err := records.Unmarshal(data); err != nil {
		return err
	}

	now := time.Now()
	records.Records = append(records.Records, &Attendance{
		Pycid:   pycid,
		CheckIn: &now,
	})

	data, err = records.Marshal()
	if err != nil {
		return
	}

	err = ioutil.WriteFile(s.filename(), data, 0644)
	s.checkedIn.Store(pycid, l)
	return
}

func (s *AttendanceStore) CheckOut(pycid string) (err error) {
	exists, err := s.alreadyCheckIn(pycid)
	if err != nil {
		return
	} else if !exists {
		return ErrRecordNotFound
	}

	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	records := new(Records)

	if err := records.Unmarshal(data); err != nil {
		return err
	}

	now := time.Now()
	modified := false

	for i, record := range records.Records {
		if record.CheckOut == nil && record.Pycid == pycid {
			record.CheckOut = &now
			record.Rank = Rank_fair
			records.Records[i] = record
			modified = true
			break
		}
	}

	if !modified {
		return ErrRecordNotFound
	}

	data, err = records.Marshal()
	if err != nil {
		return
	}

	err = ioutil.WriteFile(s.filename(), data, 0644)
	s.checkedIn.Delete(pycid)
	return
}

func (s *AttendanceStore) CheckOutAll() (err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	records := new(Records)

	if err := records.Unmarshal(data); err != nil {
		return err
	}

	now := time.Now()
	modified := false

	for i, record := range records.Records {
		if record.CheckOut == nil {
			record.CheckOut = &now
			record.Rank = Rank_fair
			records.Records[i] = record
			modified = true
			break
		}
	}

	if !modified {
		return ErrRecordNotFound
	}

	data, err = records.Marshal()
	if err != nil {
		return
	}

	err = ioutil.WriteFile(s.filename(), data, 0644)
	s.checkedIn = sync.Map{}
	return
}
