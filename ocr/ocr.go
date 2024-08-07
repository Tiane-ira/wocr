package ocr

import (
	"context"
	"fmt"
	"wocr/model"
	"wocr/utils"
)

type Ocr interface {
	GetFileList() ([]string, error)
	OcrInvoice(filename string) (*model.InvocieEx, error)
}

func NewOcr(ctx *context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	switch param.Type {
	case "百度云":
		ocr, err = NewBaidu(ctx, param)
	case "腾讯云":
		ocr, err = NewTencent(ctx, param)
	case "阿里云":
		ocr, err = NewAli(ctx, param)
	default:
		err = fmt.Errorf("暂不支持此厂商:%s", param.Type)
	}
	return
}

type OcrInstance struct {
	ctx      *context.Context
	ocr      Ocr
	savePath string
}

func NewOcrInstance(ctx *context.Context, param *model.OcrParam) (instance *OcrInstance, err error) {
	instance = &OcrInstance{
		ctx:      ctx,
		savePath: param.SavePath,
	}
	instance.ocr, err = NewOcr(ctx, param)
	return
}

// OcrInvoice 识别发票
// 模板方法
func (o *OcrInstance) OcrInvoice() (data *model.OcrResult, err error) {
	data = &model.OcrResult{}
	// 获取扫描文件列表
	fileList, err := o.ocr.GetFileList()
	if err != nil {
		return
	}
	utils.EventLog(o.ctx, "开始进行OCR扫描......")
	if len(fileList) < 1 {
		err = fmt.Errorf("扫描路径未识别到有效文件")
		return
	}
	count := len(fileList)
	utils.EventLog(o.ctx, "扫描路径匹配到文件数量: %d", count)
	data.Total = count
	// 文件识别
	dataList := make([]model.InvocieEx, 0)
	for _, file := range fileList {
		utils.EventLog(o.ctx, "开始扫描文件: %s", file)
		var invoice *model.InvocieEx
		invoice, err = o.ocr.OcrInvoice(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			utils.EventLog(o.ctx, "扫描文件异常: %s", err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, *invoice)
		}
		utils.EventLog(o.ctx, "完成扫描文件: %s", file)
	}
	// 导出excel
	savePath := utils.GetSavePath(o.savePath)
	field := &model.ExportField{}
	fieldNames, err := field.GetExports()
	if err != nil {
		return
	}
	err = utils.Export(savePath, fieldNames, dataList)
	if err != nil {
		err = fmt.Errorf("导出文件异常: %s", err.Error())
		return
	}
	data.SavePath = savePath
	return
}
