package main

import (
	"context"
	"embed"
	"wocr/backend/model"
	"wocr/backend/service"
	"wocr/backend/utils"

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

const appName = "wocr"

func main() {
	// Create an instance of the app structure
	configService := service.NewConfigService()
	fieldsService := service.NewFieldsService()
	ocrService := service.NewOcrService()
	systemService := service.NewSystemService()
	model.Init()
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
		},
		OnDomReady: func(ctx context.Context) {
			// runtime2.WindowSetPosition(ctx, x, y)
			// runtime2.WindowShow(ctx)
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			runtime.LogErrorf(ctx, "程序关闭，%s", ctx.Err().Error())
			return false
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
