package model

type OcrParam struct {
	OcrPath   string `json:"ocrPath"`
	SavePath  string `json:"savePath"`
	Recursive bool   `json:"recursive"`
	SkConfig
}

type OcrResult struct {
	Total      int
	Success    int
	Failed     int
	SavePath   string
	FailedList []string
}
