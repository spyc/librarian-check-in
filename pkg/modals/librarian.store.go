package modals

import (
	"errors"
	"io/ioutil"
	"path"
)

const LibrarianFile = "librarian.data"

var ErrNotFound = errors.New("librarian not found")

type LibrarianStore struct {
	AbstractStore `inject:"inline"`
}

func (s *LibrarianStore) filename() string {
	return path.Join(s.getRootDir(), "librarian.data")
}

func (s *LibrarianStore) GetAll() (_ []*Librarian, err error) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	data, err := ioutil.ReadFile(s.filename())
	if err != nil {
		return nil, err
	}

	librarians := new(Librarians)

	if err := librarians.Unmarshal(data); err != nil {
		return nil, err
	}

	return librarians.Members, nil
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

	return ioutil.WriteFile(s.filename(), data, 0644)
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
			return ioutil.WriteFile(s.filename(), data, 0644)
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
			return ioutil.WriteFile(s.filename(), data, 0644)
		}
	}

	return ErrNotFound
}
