package modals

import (
	"io/ioutil"
	"path"
)

const LibrarianFile = "librarian.data"

type LibrarianStore struct {
	AbstractStore `inject:"inline"`
}

func (s *LibrarianStore) GetAll() (_ []*Librarian, err error) {
	filename := path.Join(s.getRootDir(), "librarian.data")

	data, err := ioutil.ReadFile(filename)

	librarians := new(Librarians)

	if err := librarians.Unmarshal(data); err != nil {
		return nil, err
	}

	return librarians.Members, nil
}
