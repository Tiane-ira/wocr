package utils

import (
	"fmt"
	"testing"
)

func TestGetUserHomeDir(t *testing.T) {
	path := GetProjectPath()
	fmt.Println(path)
}

func TestExists(t *testing.T) {
	got := Exists(GetLocalOcrPath())
	fmt.Println(got)
}
