package ocr

import (
	"context"
	"errors"
	"fmt"
	"wocr/backend/event"
	"wocr/backend/excel"
	"wocr/backend/model"
	"wocr/backend/utils"

	"golang.org/x/time/rate"
)

var (
	ModeList = []string{"发票", "VIN码", "行程单", "车牌"}
	LimitMap = map[string]int{"发票": 2, "VIN码": 2, "行程单": 2, "车牌": 2}
	TypeList = []string{"百度云", "腾讯云", "阿里云"}
)

type Ocr interface {
	GetInvoiceFileList() ([]string, error)
	OcrInvoice(filename string) (*model.InvocieEx, error)
	OcrVin(filename string) (*model.VinEx, error)
	OcrItinerary(filename string) ([]model.ItineraryEx, error)
	OcrCarNo(filename string) (*model.CarNoEx, error)
}

func NewOcr(ctx *context.Context, param *model.OcrParam) (ocr Ocr, err error) {
	switch param.Type {
	case "百度云":
		ocr, err = NewBaidu(ctx, param)
	case "腾讯云":
		ocr, err = NewTencent(ctx, param)
	case "阿里云":
		ocr, err = NewAli(ctx, param)
	case "本地":
		ocr, err = NewLocalOcr(ctx, param)
	default:
		err = fmt.Errorf("暂不支持此厂商:%s", param.Type)
	}
	return
}

type OcrInstance struct {
	ctx     *context.Context
	ocr     Ocr
	limiter *rate.Limiter
	*model.OcrParam
}

func NewOcrInstance(ctx *context.Context, param *model.OcrParam) (instance *OcrInstance, err error) {
	instance = &OcrInstance{
		ctx:      ctx,
		OcrParam: param,
		limiter:  rate.NewLimiter(rate.Limit(LimitMap[param.Mode]), 1),
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
	case ModeList[3]:
		return o.OcrCarNo()
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
	if len(fileList) < 1 {
		err = errors.New(pathNoFilesTips)
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, matchFileLogFmt, count)
	data.Total = count
	// 文件识别
	dataList := make([]model.InvocieEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, startOcrLogFmt, file)
		var invoice *model.InvocieEx
		// 限流器阻塞
		o.limiter.Wait(context.Background())
		invoice, err = o.ocr.OcrInvoice(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, errOcrLogFmt, err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, *invoice)
		}
		event.EventLog(o.ctx, doneOcrLogFmt, file)
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
		err = fmt.Errorf(errOcrLogFmt, err.Error())
		return
	}
	data.SavePath = savePath
	return
}
func (o *OcrInstance) GetImgList() ([]string, error) {
	return utils.GetFileList(o.OcrPath, []string{jpg}, o.Recursive)
}
func (o *OcrInstance) OcrVin() (data *model.OcrResult, err error) {
	data = &model.OcrResult{}
	// 获取扫描文件列表
	fileList, err := o.GetImgList()
	if err != nil {
		return
	}
	if len(fileList) < 1 {
		err = errors.New(pathNoFilesTips)
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, matchFileLogFmt, count)
	data.Total = count
	// 文件识别
	dataList := make([]model.VinEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, startOcrLogFmt, file)
		var vin *model.VinEx
		// 限流器阻塞
		o.limiter.Wait(context.Background())
		vin, err = o.ocr.OcrVin(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, errOcrLogFmt, err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, *vin)
		}
		event.EventLog(o.ctx, doneOcrLogFmt, file)
	}
	// 导出excel
	savePath := utils.GetSavePath(o.SavePath)
	err = excel.Export(savePath, []string{"SourceFile", "VinCode"}, dataList)
	if err != nil {
		err = fmt.Errorf(exportErrLogFmt, err.Error())
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
	if len(fileList) < 1 {
		err = errors.New(pathNoFilesTips)
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, matchFileLogFmt, count)
	data.Total = count
	// 文件识别
	dataList := make([]model.ItineraryEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, startOcrLogFmt, file)
		// 限流器阻塞
		o.limiter.Wait(context.Background())
		exs, err := o.ocr.OcrItinerary(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, errOcrLogFmt, err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, exs...)
		}
		event.EventLog(o.ctx, doneOcrLogFmt, file)
	}
	if len(dataList) > 0 {
		// 导出excel
		savePath := utils.GetSavePath(o.SavePath)
		err = excel.Export(savePath, utils.GetFieldNames(&dataList[0]), dataList)
		if err != nil {
			err = fmt.Errorf(exportErrLogFmt, err.Error())
			return
		}
		data.SavePath = savePath
	}
	return
}

func (o *OcrInstance) OcrCarNo() (data *model.OcrResult, err error) {
	data = &model.OcrResult{}
	// 获取扫描文件列表
	fileList, err := o.GetImgList()
	if err != nil {
		return
	}
	if len(fileList) < 1 {
		err = errors.New(pathNoFilesTips)
		return
	}
	count := len(fileList)
	event.EventLog(o.ctx, matchFileLogFmt, count)
	data.Total = count
	// 文件识别
	dataList := make([]model.CarNoEx, 0)
	for _, file := range fileList {
		event.EventLog(o.ctx, startOcrLogFmt, file)
		var vin *model.CarNoEx
		// 限流器阻塞
		o.limiter.Wait(context.Background())
		vin, err = o.ocr.OcrCarNo(file)
		if err != nil {
			data.Failed++
			data.FailedList = append(data.FailedList, file)
			event.EventErrLog(o.ctx, errOcrLogFmt, err.Error())
			continue
		} else {
			data.Success++
			dataList = append(dataList, *vin)
		}
		event.EventLog(o.ctx, doneOcrLogFmt, file)
	}
	// 导出excel
	savePath := utils.GetSavePath(o.SavePath)
	err = excel.Export(savePath, []string{"SourceFile", "CarNo"}, dataList)
	if err != nil {
		err = fmt.Errorf(exportErrLogFmt, err.Error())
		return
	}
	data.SavePath = savePath
	return
}
