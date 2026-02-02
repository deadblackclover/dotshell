package utils

import (
	"log"
	"os"
	"os/user"
	"runtime"
)

type System struct {
	Os        string
	Arch      string
	Hostname  string
	User      *user.User
	GoVersion string
}

func (s *System) Init() {
	s.Os = runtime.GOOS
	s.Arch = runtime.GOARCH

	if hostname, err := os.Hostname(); err == nil {
		s.Hostname = hostname
	} else {
		log.Println(err)
	}

	if user, err := user.Current(); err == nil {
		s.User = user
	} else {
		log.Println(err)
	}

	s.GoVersion = runtime.Version()
}
