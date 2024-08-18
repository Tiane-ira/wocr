package ocr

import (
	"context"
	"encoding/json"
	"fmt"
	"wocr/backend/model"
	"wocr/backend/utils"
)

type LocalOcr struct {
	ctx          *context.Context
	executorPath string
	ocrPath      string
	recurcive    bool
}

func NewLocalOcr(ctx *context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	path := utils.GetLocalOcrPath()
	if !utils.Exists(path) {
		return nil, fmt.Errorf("未找到本地OCR执行器")
	}
	return &LocalOcr{
		ctx:          ctx,
		executorPath: utils.GetLocalOcrPath(),
		ocrPath:      param.OcrPath,
		recurcive:    param.Recursive,
	}, nil
}

// GetInvoiceFileList implements Ocr.
func (t *LocalOcr) GetInvoiceFileList() ([]string, error) {
	return nil, fmt.Errorf("暂不支持此功能")
}

// OcrInvoice implements Ocr.
func (t *LocalOcr) OcrInvoice(filename string) (ex *model.InvocieEx, err error) {
	return nil, fmt.Errorf("暂不支持此功能")
}

// OcrVin implements Ocr.
func (t *LocalOcr) OcrVin(filename string) (ex *model.VinEx, err error) {
	output := utils.ExecShell(t.executorPath, "-p", filename)
	res := &model.LocalOcr{}
	err = json.Unmarshal([]byte(output), &res)
	if err != nil {
		return
	}
	if res.Err != "" {
		err = fmt.Errorf(res.Err)
		return
	}
	match, matchCode := utils.GetVincode(res.Result)
	if match {
		ex = res.ToVinEx(filename, matchCode)
	} else {
		err = fmt.Errorf("本地OCR,未匹配到VIN码")
	}
	return
}

// OcrItinerary implements Ocr.
func (t *LocalOcr) OcrItinerary(filename string) (ex []model.ItineraryEx, err error) {
	return nil, fmt.Errorf("暂不支持此功能")
}
