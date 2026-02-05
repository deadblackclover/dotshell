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
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func ExecuteCmd(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("ExecuteCmd: command is empty")
	}

	var shell string
	var param string

	switch runtime.GOOS {
	case "windows":
		shell = "powershell"
		param = "-command"
	default:
		shell = "sh"
		param = "-c"
	}

	cmd := exec.Command(shell, param, command)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
