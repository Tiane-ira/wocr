package service

import (
	"context"
	"fmt"
	"strings"
	"time"
	"wocr/backend/model"
	"wocr/backend/ocr"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConfigService struct {
	ctx context.Context
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (c *ConfigService) Start(ctx context.Context) {
	c.ctx = ctx
}

func (c *ConfigService) GetConfig() []model.SkConfig {
	config := new(model.SkConfig)
	configs, err := config.ListAll()
	if err != nil {
		return nil
	}
	runtime.LogDebugf(c.ctx, "读取配置:%+v", configs)
	return configs
}

func (c *ConfigService) AddConfig(config model.SkConfig) bool {
	runtime.LogDebugf(c.ctx, "保存配置:%v", config)
	config.Id = uuid.NewString()
	config.Date = time.Now().Format("2006-01-02 15:04:05")
	config.Ak = strings.TrimSpace(config.Ak)
	config.Sk = strings.TrimSpace(config.Sk)
	err := config.Create()
	if err != nil {
		runtime.LogDebugf(c.ctx, "保存配置失败:%v", err)
		return false
	}
	return true
}

func (c *ConfigService) RemoveConfig(id string) bool {
	config := model.SkConfig{Id: id}
	err := config.Delete()
	if err != nil {
		runtime.LogDebugf(c.ctx, "移除配置失败:%v", err)
		return false
	}
	return true
}

func (c *ConfigService) GetConfigCount() int64 {
	config := &model.SkConfig{}
	num, err := config.Count()
	if err != nil {
		runtime.LogDebugf(c.ctx, "count配置失败:%v", err)
		return 0
	}
	return num
}

func (c *ConfigService) ChangeDefault(id string) (err error) {
	cnf := &model.SkConfig{}
	err = cnf.ChangeDefault(id)
	if err != nil {
		runtime.LogDebugf(c.ctx, "设置默认密钥失败:%v", err)
	}
	return
}

func (c *ConfigService) CheckSk(config model.SkConfig) string {
	switch config.Type {
	case ocr.TypeList[0]:
		token, err := ocr.GetBdAccessToken(config.Ak, config.Sk)
		if err != nil || token == "" {
			return fmt.Sprintf("%s密钥验证不通过", ocr.TypeList[0])
		}
	default:
	}
	return ""
}
