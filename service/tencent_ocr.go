package service

import (
	"context"
	"fmt"
	"wocr/model"
	"wocr/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

var (
	tencentImgs = []string{".jpg", ".jpeg", ".png"}
)

type TencentOcr struct {
	ctx       context.Context
	secretId  string
	secretKey string
	ocrPath   string
	savePath  string
	recurcive bool
}

func NewTencentOcr(ctx context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	return &TencentOcr{
		ctx:       ctx,
		secretId:  param.Ak,
		secretKey: param.Sk,
		ocrPath:   param.OcrPath,
		savePath:  param.SavePath,
		recurcive: param.Recursive,
	}, nil
}

func (t *TencentOcr) OcrInvoice() (string, error) {
	exts := append([]string{pdf}, tencentImgs...)
	dataList := []model.InvocieEx{}
	images := utils.GetFileList(t.ocrPath, exts, t.recurcive)
	for _, filename := range images {
		result, err := t.getInvociePdf(filename)
		if err != nil {
			return "", err
		}
		data := model.PdfToInvoiceEx(filename, result.Response)
		dataList = append(dataList, data)
	}
	ofds := utils.GetFileList(t.ocrPath, []string{ofd}, t.recurcive)
	if (len(images) + len(ofds)) < 1 {
		return "", fmt.Errorf("扫描路径未识别到有效文件")
	}
	for _, filename := range ofds {
		result, err := t.getInvocieOfd(filename)
		if err != nil {
			return "", err
		}
		data := model.OfdToInvoiceEx(filename, result.Response)
		dataList = append(dataList, data)
	}
	savePath := utils.GetSavePath(t.savePath)
	field := &model.ExportField{}
	fieldNames, err := field.GetExports()
	if err != nil {
		return "", err
	}
	err = utils.Export(savePath, fieldNames, dataList)
	if err != nil {
		return "", err
	}
	return savePath, nil
}

func (t *TencentOcr) getInvociePdf(filename string) (result *ocr.VatInvoiceOCRResponse, err error) {
	credential := common.NewCredential(
		t.secretId,
		t.secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-beijing", cpf)
	request := ocr.NewVatInvoiceOCRRequest()
	request.ImageBase64 = common.StringPtr(utils.Base64(filename, false))
	request.IsPdf = common.BoolPtr(true)
	result, err = client.VatInvoiceOCR(request)
	if err != nil {
		return nil, fmt.Errorf("腾讯云OCR接口异常:%s", err.Error())
	}
	return
}

func (t *TencentOcr) getInvocieOfd(filename string) (result *ocr.VerifyOfdVatInvoiceOCRResponse, err error) {
	credential := common.NewCredential(
		t.secretId,
		t.secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-beijing", cpf)
	request := ocr.NewVerifyOfdVatInvoiceOCRRequest()
	request.OfdFileBase64 = common.StringPtr(utils.Base64(filename, false))
	result, err = client.VerifyOfdVatInvoiceOCR(request)
	if err != nil {
		return nil, fmt.Errorf("腾讯云OCR接口异常:%s", err.Error())
	}
	return
}
