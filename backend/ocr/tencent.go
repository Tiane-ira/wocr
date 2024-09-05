package ocr

import (
	"context"
	"fmt"
	"strings"
	"wocr/backend/model"
	"wocr/backend/utils"

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

// GetInvoiceFileList implements Ocr.
func (t *Tencent) GetInvoiceFileList() ([]string, error) {
	exts := []string{pdf, ofd, jpg, jpeg, png}
	return utils.GetFileList(t.ocrPath, exts, t.recurcive)
}

// OcrInvoice implements Ocr.
func (t *Tencent) OcrInvoice(filename string) (ex *model.InvocieEx, err error) {
	client, err := t.getClient()
	if err != nil {
		return
	}
	if strings.HasSuffix(filename, ofd) {
		request := ocr.NewVerifyOfdVatInvoiceOCRRequest()
		request.OfdFileBase64 = common.StringPtr(utils.Base64(filename, false))
		var result *ocr.VerifyOfdVatInvoiceOCRResponse
		result, err = client.VerifyOfdVatInvoiceOCR(request)
		if err != nil {
			err = fmt.Errorf(tencentErrFmt, err.Error())
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
			return nil, fmt.Errorf(tencentErrFmt, err.Error())
		}
		ex = model.TencentPdfToInvoiceEx(filename, result.Response)
	}
	return
}

// OcrVin implements Ocr.
func (t *Tencent) OcrVin(filename string) (ex *model.VinEx, err error) {
	client, err := t.getClient()
	if err != nil {
		return
	}
	request := ocr.NewGeneralAccurateOCRRequest()
	request.ImageBase64 = common.StringPtr(utils.Base64(filename, false))
	var result *ocr.GeneralAccurateOCRResponse
	result, err = client.GeneralAccurateOCR(request)
	if err != nil {
		err = fmt.Errorf(tencentErrFmt, err.Error())
		return
	}
	code := ""
	for _, item := range result.Response.TextDetections {
		match, matchCode := utils.GetVincode(*item.DetectedText)
		if match {
			code = matchCode
			break
		}
	}
	if code == "" {
		err = fmt.Errorf("未识别到VIN码: %s", filename)
		return
	}
	ex = model.TencentToVinEx(filename, code)
	return
}

func (t *Tencent) getClient() (*ocr.Client, error) {
	credential := common.NewCredential(
		t.secretId,
		t.secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	return ocr.NewClient(credential, "ap-beijing", cpf)
}

// OcrItinerary implements Ocr.
func (t *Tencent) OcrItinerary(filename string) (ex []model.ItineraryEx, err error) {
	return nil, fmt.Errorf(unsupportTips)
}
