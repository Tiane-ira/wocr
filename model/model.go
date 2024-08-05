package model

type OcrParam struct {
	OcrPath   string `json:"ocrPath"`
	SavePath  string `json:"savePath"`
	Recursive bool   `json:"recursive"`
	SkConfig
}
