package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
	"wocr/model"
	"wocr/service"
	"wocr/utils"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

var (
	bd = "百度云"
	tc = "腾讯云"
)

func (a *App) DoOcr(param model.OcrParam) (path string) {
	utils.EventLog(a.ctx, "开始进行OCR扫描......")
	config := &model.SkConfig{Id: param.Id}
	err := config.GetById()
	if err != nil {
		utils.EventLog(a.ctx, "匹配密钥失败:%s", config.Id)
		return ""
	}
	utils.EventLog(a.ctx, "匹配密钥成功:%s", config.Name)
	param.SkConfig = *config
	ocr, err := service.NewOcr(a.ctx, &param)
	if err != nil {
		utils.EventLog(a.ctx, "匹配OCR厂商失败:%s", err.Error())
		return
	}
	utils.EventLog(a.ctx, "匹配OCR厂商成功:%s", config.Type)
	path, err = ocr.OcrInvoice()
	if err != nil {
		utils.EventLog(a.ctx, "OCR扫描异常:%s", err.Error())
		return
	}
	utils.EventLog(a.ctx, "扫描成功,保存路径,%s", path)
	return
}

// SelectDir 选择需要目录
func (a *App) SelectDir() string {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:                "选择目录",
		CanCreateDirectories: true,
	})
	if err != nil {
		runtime.LogErrorf(a.ctx, "打开目录选择弹窗失败: %s", err.Error())
	}
	return selection
}

func (a *App) OpenFile(filepath string) {
	runtime.LogDebug(a.ctx, filepath)
	osType := utils.GetOS()
	runtime.LogDebug(a.ctx, osType)
	switch osType {
	case "Windows":
		// exec.Command("start", filepath).Run()
		exec.Command("rundll32", "url.dll,FileProtocolHandler", filepath).Run()
	case "macOS":
		exec.Command("open", filepath).Run()
	default:
	}
}

func (a *App) GetConfig() []model.SkConfig {
	config := new(model.SkConfig)
	configs, err := config.ListAll()
	if err != nil {
		return nil
	}
	runtime.LogDebugf(a.ctx, "读取配置:%+v", configs)
	return configs
}

func (a *App) AddConfig(config model.SkConfig) bool {
	runtime.LogDebugf(a.ctx, "保存配置:%v", config)
	config.Id = uuid.NewString()
	config.Date = time.Now().Format("2006-01-02 15:04:05")
	config.Ak = strings.TrimSpace(config.Ak)
	config.Sk = strings.TrimSpace(config.Sk)
	err := config.Create()
	if err != nil {
		runtime.LogDebugf(a.ctx, "保存配置失败:%v", err)
		return false
	}
	return true
}

func (a *App) RemoveConfig(id string) bool {
	config := model.SkConfig{Id: id}
	err := config.Delete()
	if err != nil {
		runtime.LogDebugf(a.ctx, "移除配置失败:%v", err)
		return false
	}
	return true
}

func (a *App) GetConfigCount() int64 {
	config := &model.SkConfig{}
	num, err := config.Count()
	if err != nil {
		runtime.LogDebugf(a.ctx, "count配置失败:%v", err)
		return 0
	}
	return num
}

func (a *App) CheckSk(config model.SkConfig) string {
	switch config.Type {
	case bd:
		token, err := service.GetAccessToken(config.Ak, config.Sk)
		if err != nil || token == "" {
			return fmt.Sprintf("%s密钥验证不通过", bd)
		}
	case tc:
	}
	return ""
}

func (a *App) GetFields() (list []model.ExportField, err error) {
	field := &model.ExportField{}
	list, err = field.ListAll()
	return
}

func (a *App) UpdateFields(idList []int64) (err error) {
	field := &model.ExportField{}
	err = field.Update(idList)
	if err != nil {
		runtime.LogDebugf(a.ctx, "更新失败:%v", err)
	}
	return
}
