package ocr

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"wocr/backend/model"
	"wocr/backend/utils"
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
	// err = param.SkConfig.GetById()
	// if err != nil {
	// 	return
	// }
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
		return "", fmt.Errorf("json解析失败,err:%s,source:%s", string(data), err.Error())
	}
	return token.AccessToken, nil
}

// GetInvoiceFileList implements Ocr.
func (b *Baidu) GetInvoiceFileList() (fileList []string, err error) {
	exts := []string{pdf, ofd, jpg, jpeg, png, bmp}
	return utils.GetFileList(b.ocrPath, exts, b.recurcive)
}

// OcrInvoice implements Ocr.
func (b *Baidu) OcrInvoice(filename string) (ex *model.InvocieEx, err error) {
	params := map[string]string{
		"access_token": b.token,
	}
	result := &model.RespBdInvoice{}
	body := b.getReqBody(filename)
	body["seal_tag"] = "true"
	data, err := utils.PostWithForm(baiduInvoiceUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return
	}
	if result.ErrMsg != "" {
		err = fmt.Errorf("百度云发票OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		return
	}
	ex = result.WordsResult.ToInvoiceEx(filename)
	return
}

func (b *Baidu) getReqBody(filename string) map[string]string {
	body := map[string]string{}
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

// OcrVin implements Ocr.
func (b *Baidu) OcrVin(filename string) (ex *model.VinEx, err error) {
	params := map[string]string{
		"access_token": b.token,
	}
	result := &model.RespBdWords{}
	body := b.getReqBody(filename)
	body["paragraph"] = "true"
	body["detect_direction"] = "true"
	data, err := utils.PostWithForm(baiduAccurateUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		err = fmt.Errorf("json解析失败,source:%s,err:%s", string(data), err.Error())
		return
	}
	if result.ErrMsg != "" {
		err = fmt.Errorf("百度云精确OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		return
	}
	code := ""
	for _, item := range result.WordsResult {
		match, matchCode := utils.GetVincode(item.Words)
		if match {
			code = matchCode
			break
		}
	}
	if code == "" {
		err = fmt.Errorf("未识别到VIN码: %s", filename)
		return
	}
	ex = result.ToVinEx(filename, code)
	return
}

// OcrItinerary implements Ocr.
func (b *Baidu) OcrItinerary(filename string) (exs []model.ItineraryEx, err error) {
	params := map[string]string{
		"access_token": b.token,
	}
	exs = make([]model.ItineraryEx, 0)
	result := &model.RespBdWords{}
	body := b.getReqBody(filename)
	body["paragraph"] = "true"
	// 获取头部信息
	data, err := utils.PostWithForm(baiduGenteUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return
	}
	if result.ErrMsg != "" {
		err = fmt.Errorf("百度云普通OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		return
	}
	ex := model.BdExtractItinerary(filename, result)

	// 获取明细信息
	tableResult := &model.RespBdRespTable{}
	body = b.getReqBody(filename)
	body["cell_contents"] = "false"
	body["return_excel"] = "false"
	tableData, err := utils.PostWithForm(baiduTableUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tableData, tableResult)
	if err != nil {
		return
	}
	if result.ErrMsg != "" {
		err = fmt.Errorf("百度云表格OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		return
	}
	if len(tableResult.TableResult) > 0 {
		exs = model.BdExtractItineraryDetail(ex, tableResult)
	} else {
		exs = model.MatchBdGeneral(ex, result)
	}
	return
}

// OcrCarNo implements Ocr.
func (b *Baidu) OcrCarNo(filename string) (ex *model.CarNoEx, err error) {
	params := map[string]string{
		"access_token": b.token,
	}
	result := &model.RespBdCarNo{}
	body := b.getReqBody(filename)
	data, err := utils.PostWithForm(baiduCarNoUrl, params, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		err = fmt.Errorf("json解析失败,source:%s,err:%s", string(data), err.Error())
		return
	}
	if result.ErrMsg != "" {
		err = fmt.Errorf("百度云车牌OCR接口异常: %d:%s", result.ErrCode, result.ErrMsg)
		return
	}
	ex = model.NewCarNoEx(filename, result.WordsResult.CarNo)
	return
}
