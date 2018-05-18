package main

import (
	"os"
	"path"

	"io/ioutil"
	"library.pyc.edu.hk/attendance/pkg/modals"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dataPath := path.Join(pwd, "data")

	{
		librarians := new(modals.Librarians)
		librarians.Members = make([]*modals.Librarian, 0)

		content, err := librarians.Marshal()
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(path.Join(dataPath, "librarian.data"), content, 0644); err != nil {
			panic(err)
		}
	}

	{
		records := new(modals.Records)
		records.Records = make([]*modals.Attendance, 0)

		content, err := records.Marshal()
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(path.Join(dataPath, "attendance.data"), content, 0644); err != nil {
			panic(err)
		}
	}
}
