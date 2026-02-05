// Copyright (c) 2026, DEADBLACKCLOVER.

// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
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
