package ocr

const (
	baiduAuthUrl    = "https://aip.baidubce.com/oauth/2.0/token"
	baiduInvoiceUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"
	pdf             = ".pdf"
	ofd             = ".ofd"
	jpg             = ".jpg"
	jpeg            = ".jpeg"
	png             = ".png"
	bmp             = ".bmp"
	gif             = ".gif"
	tiff            = ".tiff"
	webp            = ".webp"
)

var (
	TypeList = []string{"百度云", "腾讯云", "阿里云"}
)
