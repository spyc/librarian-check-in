package modals

import (
	"os"
	"path"
	"sync"
)

type AbstractStore struct {
	Mutex sync.RWMutex `inject:"private"`
}

func (s *AbstractStore) getRootDir() string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "data")
}
