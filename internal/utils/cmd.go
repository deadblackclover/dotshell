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
