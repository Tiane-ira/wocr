//go:build darwin

package utils

import (
	"os/exec"
)

func ExecShell(bin string, param string, arg string) string {
	cmd := exec.Command(bin, param, arg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "error"
	}
	return string(output)
}
