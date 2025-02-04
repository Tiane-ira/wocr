package service

import (
	"context"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TrayService struct {
	ctx  context.Context
	icon []byte
}

func NewTrayService(icon []byte) *TrayService {
	return &TrayService{icon: icon}
}

func (t *TrayService) Start(ctx context.Context) {
	t.ctx = ctx
	systray.Run(t.onReady, func() {})

}

func (t *TrayService) onReady() {
	systray.SetIcon(t.icon)
	systray.SetTitle("wocr")
	systray.SetTitle("发票识别工具")
	openItem := systray.AddMenuItem("打开", "打开主面板")
	quitItem := systray.AddMenuItem("退出", "退出软件")
	go func() {
		for range openItem.ClickedCh {
			t.openMainWindow()
		}
	}()

	go func() {
		for range quitItem.ClickedCh {
			systray.Quit()
			runtime.Quit(t.ctx)
			return
		}
	}()
}

func (t *TrayService) openMainWindow() {
	runtime.WindowShow(t.ctx)
}
