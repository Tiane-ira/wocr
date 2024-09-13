package service

import (
	"context"
	"wocr/backend/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FieldsService struct {
	ctx context.Context
}

func NewFieldsService() *FieldsService {
	return &FieldsService{}
}

func (f *FieldsService) Start(ctx context.Context) {
	f.ctx = ctx
}

func (f *FieldsService) GetFields() (list []model.ExportField, err error) {
	field := &model.ExportField{}
	list, err = field.ListAll()
	return
}

func (f *FieldsService) UpdateFields(idList []int64) (err error) {
	field := &model.ExportField{}
	err = field.Update(idList)
	if err != nil {
		runtime.LogDebugf(f.ctx, "更新失败:%v", err)
	}
	return
}
