package ocr

import (
	"context"
	"fmt"
	"strings"
	"wocr/model"
	"wocr/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

type Tencent struct {
	ctx       *context.Context
	secretId  string
	secretKey string
	ocrPath   string
	recurcive bool
}

func NewTencent(ctx *context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	return &Tencent{
		ctx:       ctx,
		secretId:  param.Ak,
		secretKey: param.Sk,
		ocrPath:   param.OcrPath,
		recurcive: param.Recursive,
	}, nil
}

// GetFileList implements Ocr.
func (t *Tencent) GetFileList() ([]string, error) {
	exts := []string{pdf, ofd, jpg, jpeg, png}
	return utils.GetFileList(t.ocrPath, exts, t.recurcive)
}

// OcrInvoice implements Ocr.
func (t *Tencent) OcrInvoice(filename string) (ex *model.InvocieEx, err error) {
	credential := common.NewCredential(
		t.secretId,
		t.secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-beijing", cpf)
	if strings.HasSuffix(filename, ofd) {
		request := ocr.NewVerifyOfdVatInvoiceOCRRequest()
		request.OfdFileBase64 = common.StringPtr(utils.Base64(filename, false))
		var result *ocr.VerifyOfdVatInvoiceOCRResponse
		result, err = client.VerifyOfdVatInvoiceOCR(request)
		if err != nil {
			err = fmt.Errorf("腾讯云OCR接口异常:%s", err.Error())
			return
		}
		ex = model.TencentOfdToInvoiceEx(filename, result.Response)
	} else {
		request := ocr.NewVatInvoiceOCRRequest()
		request.ImageBase64 = common.StringPtr(utils.Base64(filename, false))
		request.IsPdf = common.BoolPtr(true)
		var result *ocr.VatInvoiceOCRResponse
		result, err = client.VatInvoiceOCR(request)
		if err != nil {
			return nil, fmt.Errorf("腾讯云OCR接口异常:%s", err.Error())
		}
		ex = model.TencentPdfToInvoiceEx(filename, result.Response)
	}
	return
}
