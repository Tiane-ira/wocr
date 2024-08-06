package service

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

var (
	aliExt = []string{".jpg", ".png", ".jpeg", ".bmp", ".gif", ".tiff", ".webp", ".pdf", ".ofd"}
)

type AliOcr struct {
	ctx       context.Context
	client    *ocr.Client
	ocrPath   string
	savePath  string
	recurcive bool
}

func NewAliOcr(ctx context.Context, param *model.OcrParam) (ocr *AliOcr, err error) {
	client, err := CreateAliClient(param.Ak, param.Sk)
	if err != nil {
		return
	}
	ocr = &AliOcr{
		ctx:       ctx,
		client:    client,
		ocrPath:   param.OcrPath,
		savePath:  param.SavePath,
		recurcive: param.Recursive,
	}
	return
}

func (a *AliOcr) OcrInvoice() (string, error) {
	fileList := utils.GetFileList(a.ocrPath, aliExt, a.recurcive)
	if len(fileList) < 1 {
		return "", fmt.Errorf("扫描路径未识别到有效文件")
	}
	utils.EventLog(a.ctx, "读取到文件数量:%d", len(fileList))
	dataList := []model.InvocieEx{}
	for _, filename := range fileList {
		utils.EventLog(a.ctx, "开始处理文件: %s", filename)
		f, err := os.Open(filename)
		if err != nil {
			return "", err
		}
		defer f.Close()
		request := &ocr.RecognizeInvoiceRequest{Body: io.Reader(f)}
		resp, err := a.client.RecognizeInvoice(request)
		if err != nil {
			return "", err
		}
		if resp.Body.Message != nil {
			return "", fmt.Errorf("阿里云OCR接口异常: %s:%s", *resp.Body.Code, *resp.Body.Data)
		}
		aliData := &model.RespAliInvoice{}
		err = json.Unmarshal([]byte(*resp.Body.Data), aliData)
		if err != nil {
			return "", err
		}
		utils.EventLog(a.ctx, "结束处理文件: %s", filename)
		dataList = append(dataList, aliData.Data.AliToInvoiceEx(filename))
	}
	savePath := utils.GetSavePath(a.savePath)
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

func CreateAliClient(ak, sk string) (client *ocr.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(ak),
		AccessKeySecret: tea.String(sk),
	}
	config.Endpoint = tea.String("ocr-api.cn-hangzhou.aliyuncs.com")
	client = &ocr.Client{}
	client, err = ocr.NewClient(config)
	return
}
