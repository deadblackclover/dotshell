package utils

import (
	"os"
	"runtime"
)

type System struct {
	Os          string
	Arch        string
	Hostname    string
	UserHomeDir string
	GoVersion   string
}

func GetSystemData() System {
	var data System

	data.Os = runtime.GOOS
	data.Arch = runtime.GOARCH

	if hostname, err := os.Hostname(); err == nil {
		data.Hostname = hostname
	}

	if userHomeDir, err := os.UserHomeDir(); err == nil {
		data.UserHomeDir = userHomeDir
	}

	data.GoVersion = runtime.Version()

	return data
}
