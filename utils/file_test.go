package utils

import (
	"fmt"
	"testing"
)

func TestGetInvoiceFileList(t *testing.T) {
	pdfs := GetFileList("C://Users//irajames//Documents//ocr//发票", []string{".pdf"}, true)
	fmt.Printf("%v", pdfs)
}

func TestGetSavePath(t *testing.T) {
	got := GetSavePath("C://Users//irajames//Documents//ocr")
	fmt.Printf("%s", got)
}

func TestBase64(t *testing.T) {
	fmt.Printf("Base64(\"C:/Users/irajames/Documents/ocr/发票/95799760.pdf\", false): %v\n", Base64("C:/Users/irajames/Documents/ocr/发票/95799760.pdf", false))
}
