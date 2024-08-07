#arch := $(shell go env GOOS)

win-amd64:
	wails build -clean -nsis -platform windows/amd64

win-arm64:
	wails build -clean -nsis -platform windows/arm64

mac-amd64:
	wails build -clean -platform darwin/amd64

mac-arm64:
	wails build -clean -platform darwin/arm64

