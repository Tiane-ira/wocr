package ocr

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"wocr/model"
	"wocr/utils"
)

type Baidu struct {
	ctx       *context.Context
	ocrPath   string
	recurcive bool
	token     string
}

func NewBaidu(ctx *context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	b := &Baidu{
		ctx:       ctx,
		ocrPath:   param.OcrPath,
		recurcive: param.Recursive,
	}
	err = param.SkConfig.GetById()
	if err != nil {
		return
	}
	b.token, err = GetBdAccessToken(param.SkConfig.Ak, param.SkConfig.Sk)
	if err != nil {
		return
	}
	return b, nil
}

func GetBdAccessToken(ak, sk string) (string, error) {
	params := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     ak,
		"client_secret": sk,
	}
	token := &model.RespBdToken{}
	data, err := utils.PostWithForm(baiduAuthUrl, params, nil)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(data, token)
	if err != nil {
		return "", fmt.Errorf("json解析失败,source:%s,err:%s", string(data), err.Error())
	}
	return token.AccessToken, nil
}

// GetFileList implements Ocr.
func (b *Baidu) GetFileList() (fileList []string, err error) {
	exts := []string{pdf, ofd, jpg, jpeg, png, bmp}
	return utils.GetFileList(b.ocrPath, exts, b.recurcive)
}

// OcrInvoice implements Ocr.
func (b *Baidu) OcrInvoice(filename string) (ex *model.InvocieEx, err error) {
	params := map[string]string{
		"access_token": b.token,
	}
	result := &model.RespBd{}
	body := b.getReqBody(filename)
	data, err := utils.PostWithForm(baiduInvoiceUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		err = fmt.Errorf("json解析失败,source:%s,err:%s", string(data), err.Error())
		return
	}
	if result.ErrMsg != "" {
		err = fmt.Errorf("百度云OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		return
	}
	ex = result.WordsResult.ToInvoiceEx(filename)
	return
}

func (b *Baidu) getReqBody(filename string) map[string]string {
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
	if utils.Contains([]string{jpg, jpeg, png, bmp}, filepath.Ext(filename)) {
		body["image"] = data
	}
	return body
}
