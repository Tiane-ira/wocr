package main

import (
	"context"
	"embed"
	"wocr/backend/model"
	"wocr/backend/service"
	"wocr/backend/utils"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

//go:embed build/windows/icon.ico
var wicon []byte

const appName = "wocr"

func main() {

	configService := service.NewConfigService()
	fieldsService := service.NewFieldsService()
	ocrService := service.NewOcrService()
	systemService := service.NewSystemService()
	trayService := service.NewTrayService(wicon)
	eventService := service.NewEventService()
	model.Init()
	var appCtx context.Context
	defer func() {
		if r := recover(); r != nil {
			runtime.LogErrorf(appCtx, "Application crashed: %v", r)
		}
	}()
	// Create application with options
	err := wails.Run(&options.App{
		Title:     appName,
		Width:     1024,
		Height:    768,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			configService.Start(ctx)
			fieldsService.Start(ctx)
			ocrService.Start(ctx)
			systemService.Start(ctx)
			trayService.Start(ctx)
			eventService.Start(ctx)
			appCtx = ctx
		},
		OnDomReady: func(ctx context.Context) {
			// runtime2.WindowSetPosition(ctx, x, y)
			// runtime2.WindowShow(ctx)
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Type:          runtime.QuestionDialog,
				Title:         "请确认",
				Message:       "是否退出到托盘?",
				DefaultButton: "No",
			})
			if err != nil {
				return false
			}
			if dialog == "Yes" {
				runtime.WindowHide(ctx)
				return true
			}
			return false
		},
		OnShutdown: func(ctx context.Context) {
			systray.Quit()
		},

		Bind: []interface{}{
			configService,
			fieldsService,
			ocrService,
			systemService,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   appName,
				Message: "A ocr tool cross-platform desktop app.\n\nCopyright © 2024",
				Icon:    icon,
			},
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableFramelessWindowDecorations: false,
		},
		Logger:             logger.NewFileLogger(utils.GetLogPath()),
		LogLevelProduction: logger.DEBUG,
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
