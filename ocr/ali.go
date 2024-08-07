package ocr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"wocr/model"
	"wocr/utils"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ocr "github.com/alibabacloud-go/ocr-api-20210707/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Ali struct {
	ctx       *context.Context
	client    *ocr.Client
	ocrPath   string
	recurcive bool
}

func NewAli(ctx *context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	a := &Ali{
		ctx:       ctx,
		ocrPath:   param.OcrPath,
		recurcive: param.Recursive,
	}
	a.client, err = a.CreateClient(param.Ak, param.Sk)
	return a, err
}

func (a *Ali) CreateClient(ak, sk string) (client *ocr.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(ak),
		AccessKeySecret: tea.String(sk),
	}
	config.Endpoint = tea.String("ocr-api.cn-hangzhou.aliyuncs.com")
	client = &ocr.Client{}
	client, err = ocr.NewClient(config)
	return
}

// GetFileList implements Ocr.
func (a *Ali) GetFileList() ([]string, error) {
	exts := []string{jpg, jpeg, png, bmp, gif, tiff, webp, pdf, ofd}
	return utils.GetFileList(a.ocrPath, exts, a.recurcive)
}

// OcrInvoice implements Ocr.
func (a *Ali) OcrInvoice(filename string) (ex *model.InvocieEx, err error) {
	var f *os.File
	f, err = os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	request := &ocr.RecognizeInvoiceRequest{Body: io.Reader(f)}
	resp, err := a.client.RecognizeInvoice(request)
	if err != nil {
		return
	}
	if resp.Body.Message != nil {
		err = fmt.Errorf("阿里云OCR接口异常: %s:%s", *resp.Body.Code, *resp.Body.Data)
		return
	}
	aliData := &model.RespAliInvoice{}
	err = json.Unmarshal([]byte(*resp.Body.Data), aliData)
	if err != nil {
		return
	}
	ex = aliData.Data.AliToInvoiceEx(filename)
	return
}
