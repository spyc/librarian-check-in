package main

import (
	"os"
	"path"

	"crypto/rand"
	"encoding/base32"
	"io/ioutil"
	"library.pyc.edu.hk/attendance/pkg/modals"
	"strings"
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

		if err := ioutil.WriteFile(path.Join(dataPath, modals.LibrarianFile), content, 0644); err != nil {
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

		if err := ioutil.WriteFile(path.Join(dataPath, modals.AttendanceFile), content, 0644); err != nil {
			panic(err)
		}
	}

	{
		secret := make([]byte, 10)
		rand.Read(secret)
		secretString := strings.TrimRight(base32.StdEncoding.EncodeToString(secret), "=")
		if err := ioutil.WriteFile(path.Join(dataPath, "token"), []byte(secretString), 0644); err != nil {
			panic(err)
		}

	}
}
