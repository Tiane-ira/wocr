package service

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"wocr/model"
	"wocr/utils"
)

var (
	authUrl    = "https://aip.baidubce.com/oauth/2.0/token"             // 认证连接
	invoiceUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice" // 发票
	pdf        = ".pdf"
	ofd        = ".ofd"
	imgs       = []string{".jpg", ".jpeg", ".png", ".bmp"}
)

type BdiduOcr struct {
	ctx       context.Context
	token     string // 默认有效期30天, 不刷新了
	ocrPath   string
	savePath  string
	recurcive bool
}

func NewBaiduOcr(ctx context.Context, param *model.OcrParam) (*BdiduOcr, error) {
	token, err := GetAccessToken(param.Ak, param.Sk)
	if err != nil {
		return nil, err
	}
	return &BdiduOcr{
		ctx:       ctx,
		token:     token,
		ocrPath:   param.OcrPath,
		savePath:  param.SavePath,
		recurcive: param.Recursive,
	}, nil
}

// OcrInvoice implements Ocr.
func (b *BdiduOcr) OcrInvoice() (string, error) {
	exts := append([]string{pdf, ofd}, imgs...)
	fileList := utils.GetFileList(b.ocrPath, exts, b.recurcive)
	if len(fileList) < 1 {
		return "", fmt.Errorf("扫描路径未识别到有效文件")
	}
	utils.EventLog(b.ctx, "读取到文件数量:%d", len(fileList))
	dataList := []model.InvocieEx{}
	for _, filename := range fileList {
		utils.EventLog(b.ctx, "开始处理文件: %s", filename)
		result, err := b.getInvoiceResult(filename)
		if err != nil {
			return "", err
		}
		if result.ErrMsg != "" {
			return "", fmt.Errorf("百度云OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		}
		ex := result.WordsResult.ToInvoiceEx(filename)
		dataList = append(dataList, ex)
		utils.EventLog(b.ctx, "结束处理文件: %s", filename)
	}
	savePath := utils.GetSavePath(b.savePath)
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

func (b *BdiduOcr) getInvoiceResult(filename string) (result *model.RespBd, err error) {
	params := map[string]string{
		"access_token": b.token,
	}
	result = &model.RespBd{}
	body := b.getReqBody(filename)
	data, err := utils.PostWithForm(invoiceUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("json解析失败,source:%s,err:%s", string(data), err.Error())
	}
	return
}
func (b *BdiduOcr) getReqBody(filename string) map[string]string {
	body := map[string]string{
		"seal_tag": "true",
	}
	data := utils.Base64(filename, true)
	if strings.HasSuffix(filename, pdf) {
		body["pdf_file"] = data
	}
	if strings.HasSuffix(filename, ofd) {
		body["ofd_file"] = data
	}
	if utils.Contains(imgs, filepath.Ext(filename)) {
		body["image"] = data
	}
	return body
}

func GetAccessToken(ak, sk string) (string, error) {
	params := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     ak,
		"client_secret": sk,
	}
	token := &model.RespBdToken{}
	data, err := utils.PostWithForm(authUrl, params, nil)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(data, token)
	if err != nil {
		return "", fmt.Errorf("json解析失败,source:%s,err:%s", string(data), err.Error())
	}
	return token.AccessToken, nil
}
