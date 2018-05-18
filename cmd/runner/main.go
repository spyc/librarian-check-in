package main

import (
	"net/http"

	"library.pyc.edu.hk/attendance/pkg/bootstrap"
)

func main() {
	bootstrap.GetLogger().Debugln("Start serving")
	if err := http.ListenAndServe(":8080", bootstrap.GetHandler()); err != nil {
		panic(err)
	}
}
