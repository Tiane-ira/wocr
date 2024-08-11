package utils

import (
	"context"
	"fmt"
	"wocr/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EventLog(ctx *context.Context, format string, args ...interface{}) {
	runtime.EventsEmit(*ctx, "ocr_log", model.OcrLog{IsError: false, Msg: fmt.Sprintf(format, args...)})
}

func EventErrLog(ctx *context.Context, format string, args ...interface{}) {
	runtime.EventsEmit(*ctx, "ocr_log", model.OcrLog{IsError: true, Msg: fmt.Sprintf(format, args...)})
}
