package utils

import (
	"fmt"
	"testing"
)

func TestExecShell(t *testing.T) {
	// got := ExecShell("/Users/james/.wocr/rapidocr", "-p", "/Users/james/Documents/ocr/车架照片/IMG_20240401_155613.jpg")
	got := ExecShell("C:\\Users\\irajames\\.wocr\\rapidocr.exe", "-p", "C:\\Users\\irajames\\Documents\\ocr\\车架照片\\IMG_20240401_170211.jpg")
	fmt.Println(got)
}
