package model

type OcrParam struct {
	OcrPath   string `json:"ocrPath"`
	SavePath  string `json:"savePath"`
	Recursive bool   `json:"recursive"`
	Mode      string `json:"Mode"`
	SkConfig
}

type OcrResult struct {
	Total      int
	Success    int
	Failed     int
	SavePath   string
	FailedList []string
}

type OcrLog struct {
	IsError bool   `json:"isError"`
	Msg     string `json:"msg"`
}
