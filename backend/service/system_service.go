package service

import (
	"context"
	"os/exec"
	"wocr/backend/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SystemService struct {
	ctx context.Context
}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (s *SystemService) Start(ctx context.Context) {
	s.ctx = ctx
}

// SelectDir 选择需要目录
func (s *SystemService) SelectDir() string {
	runtime.LogInfo(appCtx, "Starting SelectDir...")
	defer func() {
		if r := recover(); r != nil {
			runtime.LogErrorf(appCtx, "Application crashed: %v", r)
		}
	}()

	selection, err := runtime.OpenDirectoryDialog(s.ctx, runtime.OpenDialogOptions{
		Title:                "选择目录",
		CanCreateDirectories: true,
	})
	if err != nil {
		runtime.LogErrorf(s.ctx, "打开目录选择弹窗失败: %s", err.Error())
	}
	return selection
}

func (s *SystemService) OpenFile(filepath string) {
	runtime.LogDebug(s.ctx, filepath)
	osType := utils.GetOS()
	runtime.LogDebug(s.ctx, osType)
	switch osType {
	case "Windows":
		// exec.Command("start", filepath).Run()
		exec.Command("rundll32", "url.dll,FileProtocolHandler", filepath).Run()
	case "macOS":
		exec.Command("open", filepath).Run()
	default:
	}
}
