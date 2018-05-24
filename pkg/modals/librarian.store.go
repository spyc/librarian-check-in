package modals

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"sync"
)

const LibrarianFile = "librarian.data"

var ErrNotFound = errors.New("librarian not found")

type LibrarianStore struct {
	AbstractStore

	Logger logrus.FieldLogger `inject:"librarian logger"`

	cache  sync.Map
	cached bool
}

func (s *LibrarianStore) filename() string {
	return path.Join(s.getRootDir(), LibrarianFile)
}

func (s *LibrarianStore) getByID(id string) (_ *Librarian, err error) {
	logger := s.Logger.WithField("func", "getByID")
	if !s.cached {
		s.Logger.Debug("get all")
		if _, err := s.GetAll(); err != nil {
			logger.Error(err)
			return nil, err
		}
	}

	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	val, exists := s.cache.Load(id)

	logger.Debug(val, exists)

	if !exists {
		return nil, nil
	}

	return val.(*Librarian), nil
}

func (s *LibrarianStore) GetAll() (_ []*Librarian, err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return nil, err
	}

	librarians := new(Librarians)

	if err := librarians.Unmarshal(data); err != nil {
		return nil, err
	}

	s.buildCache(librarians.Members)
	return librarians.Members, nil
}

func (s *LibrarianStore) buildCache(librarians []*Librarian) {
	var cache sync.Map

	for _, l := range librarians {
		cache.Store(l.Pycid, l)
		s.Logger.Debugf("Add %s to cache", l.Pycid)
	}

	s.cache = cache
	s.Logger.Debug(s.cache)
	s.cached = true
	s.Logger.Debug("renew cache")
}

func (s *LibrarianStore) Add(l *Librarian) (err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	librarians := new(Librarians)

	if err := librarians.Unmarshal(data); err != nil {
		return err
	}

	librarians.Members = append(librarians.Members, l)

	data, err = librarians.Marshal()

	err = ioutil.WriteFile(s.filename(), data, 0644)
	s.cache.Store(l.Pycid, l)
	return err
}

func (s *LibrarianStore) Remove(pycid string) (err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	librarians := new(Librarians)

	if err := librarians.Unmarshal(data); err != nil {
		return err
	}

	for i, librarian := range librarians.Members {
		if librarian.Pycid == pycid {
			librarians.Members = append(librarians.Members[:i], librarians.Members[i+1:]...)
			data, err = librarians.Marshal()
			err = ioutil.WriteFile(s.filename(), data, 0644)
			if err == nil {
				s.cache.Delete(pycid)
			}
			return err
		}
	}

	return ErrNotFound
}

func (s *LibrarianStore) Update(pycid string, l *Librarian) (err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return
	}

	librarians := new(Librarians)

	if err := librarians.Unmarshal(data); err != nil {
		return err
	}

	for i, librarian := range librarians.Members {
		if librarian.Pycid == pycid {
			librarians.Members = append(librarians.Members[:i], librarians.Members[i+1:]...)
			librarians.Members = append(librarians.Members, l)
			data, err = librarians.Marshal()
			err = ioutil.WriteFile(s.filename(), data, 0644)
			if err == nil {
				s.cache.Store(l.Pycid, l)
			}
			return err
		}
	}

	return ErrNotFound
}
