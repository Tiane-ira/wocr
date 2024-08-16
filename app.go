package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
	"wocr/event"
	"wocr/model"
	"wocr/ocr"
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

func (a *App) DoOcr(param model.OcrParam) (path string) {
	event.EventLog(&a.ctx, "开始进行OCR扫描......")
	config := &model.SkConfig{Id: param.Id}
	err := config.GetById()
	if err != nil {
		event.EventErrLog(&a.ctx, "匹配密钥失败:%s", config.Id)
		return ""
	}
	event.EventLog(&a.ctx, "匹配密钥成功:%s", config.Name)
	param.SkConfig = *config
	ocrInstance, err := ocr.NewOcrInstance(&a.ctx, &param)
	if err != nil {
		event.EventErrLog(&a.ctx, "匹配OCR厂商失败:%s", err.Error())
		return
	}
	event.EventLog(&a.ctx, "匹配OCR厂商成功:%s", config.Type)
	result, err := ocrInstance.DoOcr()
	if err != nil {
		event.EventErrLog(&a.ctx, "OCR扫描异常:%s", err.Error())
		return
	}
	event.EventLog(&a.ctx, "扫描完成,共:%d份,成功:%d份,失败:%d份", result.Total, result.Success, result.Failed)
	if len(result.FailedList) > 0 {
		event.EventErrLog(&a.ctx, "扫描异常文件: %d份", len(result.FailedList))
		for Index, failed := range result.FailedList {
			event.EventErrLog(&a.ctx, "%d: %s", Index+1, failed)
		}
	}
	if result.Success < 1 {
		return
	}
	path = result.SavePath
	event.EventLog(&a.ctx, "扫描结果保存路径:%s", result.SavePath)
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
	case ocr.TypeList[0]:
		token, err := ocr.GetBdAccessToken(config.Ak, config.Sk)
		if err != nil || token == "" {
			return fmt.Sprintf("%s密钥验证不通过", ocr.TypeList[0])
		}
	default:
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

func (a *App) ChangeDefault(id string) (err error) {
	c := &model.SkConfig{}
	err = c.ChangeDefault(id)
	if err != nil {
		runtime.LogDebugf(a.ctx, "设置默认密钥失败:%v", err)
	}
	return
}
