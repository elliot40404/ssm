package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func OpenWithDefaultTextEditor(filename string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filename)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filename)
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func ParseInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return i
}
