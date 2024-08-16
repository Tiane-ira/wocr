package ocr

import (
	"context"
	"fmt"
	"wocr/event"
	"wocr/excel"
	"wocr/model"
	"wocr/utils"
)

type Ocr interface {
	GetInvoiceFileList() ([]string, error)
	OcrInvoice(filename string) (*model.InvocieEx, error)
	OcrVin(filename string) (*model.VinEx, error)
	OcrItinerary(filename string) ([]model.ItineraryEx, error)
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
	ctx *context.Context
	ocr Ocr
	*model.OcrParam
}

func NewOcrInstance(ctx *context.Context, param *model.OcrParam) (instance *OcrInstance, err error) {
	instance = &OcrInstance{
		ctx:      ctx,
		OcrParam: param,
	}
	instance.ocr, err = NewOcr(ctx, param)
	return
}

func (o *OcrInstance) DoOcr() (data *model.OcrResult, err error) {
	switch o.Mode {
	case ModeList[0]:
		return o.OcrInvoice()
	case ModeList[1]:
		return o.OcrVin()
	case ModeList[2]:
		return o.OcrItinerary()
	default:
		err = fmt.Errorf("暂不支持模式:%s", o.Mode)
	}
	return
}

// OcrInvoice 识别发票
func (o *OcrInstance) OcrInvoice() (data *model.OcrResult, err error) {
	data = &model.OcrResult{}
	// 获取扫描文件列表
	fileList, err := o.ocr.GetInvoiceFileList()
	if err != nil {
		return
	}
	event.EventLog(o.ctx, "开始进行OCR扫描......")
	if len(fileList) < 1 {
		err = fmt.Errorf("扫描路径未识别到有效文件")
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, "扫描路径匹配到文件数量: %d", count)
	data.Total = count
	// 文件识别
	dataList := make([]model.InvocieEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, "开始扫描文件: %s", file)
		var invoice *model.InvocieEx
		invoice, err = o.ocr.OcrInvoice(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, "扫描文件异常: %s", err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, *invoice)
		}
		event.EventLog(o.ctx, "完成扫描文件: %s", file)
	}
	// 导出excel
	savePath := utils.GetSavePath(o.SavePath)
	field := &model.ExportField{}
	fieldNames, err := field.GetExports()
	if err != nil {
		return
	}
	err = excel.Export(savePath, fieldNames, dataList)
	if err != nil {
		err = fmt.Errorf("导出文件异常: %s", err.Error())
		return
	}
	data.SavePath = savePath
	return
}
func (o *OcrInstance) GetVinFileList() ([]string, error) {
	return utils.GetFileList(o.OcrPath, []string{jpg}, o.Recursive)
}
func (o *OcrInstance) OcrVin() (data *model.OcrResult, err error) {
	data = &model.OcrResult{}
	// 获取扫描文件列表
	fileList, err := o.GetVinFileList()
	if err != nil {
		return
	}
	event.EventLog(o.ctx, "开始进行OCR扫描......")
	if len(fileList) < 1 {
		err = fmt.Errorf("扫描路径未识别到有效文件")
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, "扫描路径匹配到文件数量: %d", count)
	data.Total = count
	// 文件识别
	dataList := make([]model.VinEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, "开始扫描文件: %s", file)
		var vin *model.VinEx
		vin, err = o.ocr.OcrVin(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, "扫描文件异常: %s", err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, *vin)
		}
		event.EventLog(o.ctx, "完成扫描文件: %s", file)
	}
	// 导出excel
	savePath := utils.GetSavePath(o.SavePath)
	err = excel.Export(savePath, []string{"SourceFile", "VinCode"}, dataList)
	if err != nil {
		err = fmt.Errorf("导出文件异常: %s", err.Error())
		return
	}
	data.SavePath = savePath
	return
}
func (o *OcrInstance) GetItineraryFileList() ([]string, error) {
	return utils.GetFileList(o.OcrPath, []string{pdf}, o.Recursive)
}

func (o *OcrInstance) OcrItinerary() (data *model.OcrResult, err error) {
	data = &model.OcrResult{}
	// 获取扫描文件列表
	fileList, err := o.GetItineraryFileList()
	if err != nil {
		return
	}
	event.EventLog(o.ctx, "开始进行OCR扫描......")
	if len(fileList) < 1 {
		err = fmt.Errorf("扫描路径未识别到有效文件")
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, "扫描路径匹配到文件数量: %d", count)
	data.Total = count
	// 文件识别
	dataList := make([]model.ItineraryEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, "开始扫描文件: %s", file)
		exs, err := o.ocr.OcrItinerary(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, "扫描文件异常: %s", err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, exs...)
		}
		event.EventLog(o.ctx, "完成扫描文件: %s", file)
	}
	if len(dataList) > 0 {
		// 导出excel
		savePath := utils.GetSavePath(o.SavePath)
		err = excel.Export(savePath, utils.GetFieldNames(&dataList[0]), dataList)
		if err != nil {
			err = fmt.Errorf("导出文件异常: %s", err.Error())
			return
		}
		data.SavePath = savePath
	}
	return
}
