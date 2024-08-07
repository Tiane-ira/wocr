package utils

import (
	"log"
	"testing"
	"wocr/model"
)

func TestExport(t *testing.T) {
	datalist := []model.InvocieEx{
		{InvoiceNum: "123", TotalAmount: "2123.35"},
		{InvoiceNum: "567", TotalAmount: "999.35"},
	}
	err := Export("test.xlsx", []string{"InvoiceNum", "TotalAmount"}, datalist)
	if err != nil {
		log.Fatal(err)
	}
}
