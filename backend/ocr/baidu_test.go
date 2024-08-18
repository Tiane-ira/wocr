package ocr

import (
	"log"
	"testing"
	"wocr/backend/model"
)

func TestBaiduOcrItinerary(t *testing.T) {
	b, _ := NewBaidu(nil, &model.OcrParam{SkConfig: model.SkConfig{Ak: "", Sk: ""}})
	daas, err := b.OcrItinerary("37483790行程单.pdf")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(daas)
}
