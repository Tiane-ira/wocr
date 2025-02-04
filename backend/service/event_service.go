package service

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type EventService struct {
	ctx context.Context
}

func NewEventService() *EventService {
	return &EventService{}
}

func (e *EventService) Start(ctx context.Context) {
	e.ctx = ctx
	runtime.EventsOn(ctx, "window-close", func(optionalData ...interface{}) {
		runtime.LogDebug(ctx, "Window is closing, minimizing to tray...")

		// 隐藏窗口（最小化到托盘）
		runtime.WindowHide(ctx)
	})
}
