package utils

import (
	"os/exec"
	"runtime"
	"syscall"
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
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "error"
	}
	return string(output)
}
