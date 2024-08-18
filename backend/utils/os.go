package utils

import (
	"os/exec"
	"runtime"
)

func GetOS() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "Windows"
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	default:
		return "Unknown OS"
	}
}

func ExecShell(bin string, param string, arg string) string {
	cmd := exec.Command(bin, param, arg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "error"
	}
	return string(output)
}
