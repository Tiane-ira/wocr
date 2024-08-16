package utils

import (
	"fmt"
	"testing"
)

func TestGetUserHomeDir(t *testing.T) {
	path := GetProjectPath()
	fmt.Println(path)
}
