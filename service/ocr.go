package service

import (
	"context"
	"fmt"
	"wocr/model"
)

type Ocr interface {
	OcrInvoice() (string, error)
}

func NewOcr(ctx context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	switch param.Type {
	case "百度云":
		ocr, err = NewBaiduOcr(ctx, param)
	case "腾讯云":
		ocr, err = NewTencentOcr(ctx, param)
	case "阿里云":
		ocr, err = NewAliOcr(ctx, param)
	default:
		err = fmt.Errorf("暂不支持此厂商:%s", param.Type)
	}
	return
}
