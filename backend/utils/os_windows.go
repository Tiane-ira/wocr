//go:build windows

package utils

import (
	"os/exec"
	"syscall"
)

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
