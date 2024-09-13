package ocr

const (
	baiduAuthUrl     = "https://aip.baidubce.com/oauth/2.0/token"
	baiduInvoiceUrl  = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"
	baiduAccurateUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
	baiduGenteUrl    = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	baiduTableUrl    = "https://aip.baidubce.com/rest/2.0/ocr/v1/table"
	baiduCarNoUrl    = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	pdf              = ".pdf"
	ofd              = ".ofd"
	jpg              = ".jpg"
	jpeg             = ".jpeg"
	png              = ".png"
	bmp              = ".bmp"
	gif              = ".gif"
	tiff             = ".tiff"
	webp             = ".webp"

	unsupportTips   = "暂不支持此功能"
	pathNoFilesTips = "扫描路径未识别到有效文件"
	matchFileLogFmt = "扫描路径匹配到文件数量: %d"
	startOcrLogFmt  = "开始扫描文件: %s"
	errOcrLogFmt    = "扫描文件异常: %s"
	doneOcrLogFmt   = "完成扫描文件: %s"
	exportErrLogFmt = "导出文件异常: %s"
	tencentErrFmt   = "腾讯云OCR接口异常:%s"
)
