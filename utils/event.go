package utils

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EventLog(ctx context.Context, format string, args ...interface{}) {
	runtime.EventsEmit(ctx, "ocr_log", fmt.Sprintf(format, args...))
}
