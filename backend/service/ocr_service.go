package service

import (
	"context"
	"wocr/backend/event"
	"wocr/backend/model"
	"wocr/backend/ocr"
)

type OcrService struct {
	ctx context.Context
}

func NewOcrService() *OcrService {
	return &OcrService{}
}

func (o *OcrService) Start(ctx context.Context) {
	o.ctx = ctx
}

func (o *OcrService) DoOcr(param model.OcrParam) (path string) {
	event.EventLog(&o.ctx, "开始进行OCR扫描......")
	config := &model.SkConfig{Id: param.Id}
	err := config.GetById()
	if err != nil {
		event.EventErrLog(&o.ctx, "匹配密钥失败:%s", config.Id)
		return ""
	}
	event.EventLog(&o.ctx, "匹配密钥成功:%s", config.Name)
	param.SkConfig = *config
	ocrInstance, err := ocr.NewOcrInstance(&o.ctx, &param)
	if err != nil {
		event.EventErrLog(&o.ctx, "匹配OCR厂商失败:%s", err.Error())
		return
	}
	event.EventLog(&o.ctx, "匹配OCR厂商成功:%s", config.Type)
	result, err := ocrInstance.DoOcr()
	if err != nil {
		event.EventErrLog(&o.ctx, "OCR扫描异常:%s", err.Error())
		return
	}
	event.EventLog(&o.ctx, "扫描完成,共:%d份,成功:%d份,失败:%d份", result.Total, result.Success, result.Failed)
	if len(result.FailedList) > 0 {
		event.EventErrLog(&o.ctx, "扫描异常文件: %d份", len(result.FailedList))
		for Index, failed := range result.FailedList {
			event.EventErrLog(&o.ctx, "%d: %s", Index+1, failed)
		}
	}
	if result.Success < 1 {
		return
	}
	path = result.SavePath
	event.EventLog(&o.ctx, "扫描结果保存路径:%s", result.SavePath)
	return
}
