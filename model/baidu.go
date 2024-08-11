package model

import (
	"regexp"
	"strings"
)

type RespBdToken struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int64  `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

type BdInvoiceItem struct {
	Word string `json:"word"`
	Row  string `json:"row"`
}
type RespBdInvoice struct {
	LogID          int64           `json:"log_id"`
	ErrMsg         string          `json:"error_msg"`
	ErrCode        int64           `json:"error_code"`
	WordsResultNum int64           `json:"words_result_num"`
	WordsResult    BdInvoiceResult `json:"words_result"`
}

type BdInvoiceResult struct {
	InvoiceNumDigit      string          `json:"InvoiceNumDigit"`
	ServiceType          string          `json:"ServiceType"`
	InvoiceNum           string          `json:"InvoiceNum"`
	InvoiceNumConfirm    string          `json:"InvoiceNumConfirm"`
	SellerName           string          `json:"SellerName"`
	CommodityTaxRate     []BdInvoiceItem `json:"CommodityTaxRate"`
	SellerBank           string          `json:"SellerBank"`
	Checker              string          `json:"Checker"`
	TotalAmount          string          `json:"TotalAmount"`
	CommodityAmount      []BdInvoiceItem `json:"CommodityAmount"`
	InvoiceDate          string          `json:"InvoiceDate"`
	CommodityTax         []BdInvoiceItem `json:"CommodityTax"`
	PurchaserName        string          `json:"PurchaserName"`
	CommodityNum         []BdInvoiceItem `json:"CommodityNum"`
	Province             string          `json:"Province"`
	City                 string          `json:"City"`
	SheetNum             string          `json:"SheetNum"`
	Agent                string          `json:"Agent"`
	PurchaserBank        string          `json:"PurchaserBank"`
	Remarks              string          `json:"Remarks"`
	Password             string          `json:"Password"`
	SellerAddress        string          `json:"SellerAddress"`
	PurchaserAddress     string          `json:"PurchaserAddress"`
	InvoiceCode          string          `json:"InvoiceCode"`
	InvoiceCodeConfirm   string          `json:"InvoiceCodeConfirm"`
	CommodityUnit        []BdInvoiceItem `json:"CommodityUnit"`
	Payee                string          `json:"Payee"`
	PurchaserRegisterNum string          `json:"PurchaserRegisterNum"`
	CommodityPrice       []BdInvoiceItem `json:"CommodityPrice"`
	NoteDrawer           string          `json:"NoteDrawer"`
	AmountInWords        string          `json:"AmountInWords"`
	AmountInFiguers      string          `json:"AmountInFiguers"`
	TotalTax             string          `json:"TotalTax"`
	InvoiceType          string          `json:"InvoiceType"`
	SellerRegisterNum    string          `json:"SellerRegisterNum"`
	CommodityName        []BdInvoiceItem `json:"CommodityName"`
	CommodityType        []BdInvoiceItem `json:"CommodityType"`
	CommodityPlateNum    []BdInvoiceItem `json:"CommodityPlateNum"`
	CommodityVehicleType []BdInvoiceItem `json:"CommodityVehicleType"`
	CommodityStartDate   []BdInvoiceItem `json:"CommodityStartDate"`
	CommodityEndDate     []BdInvoiceItem `json:"CommodityEndDate"`
	OnlinePay            string          `json:"OnlinePay"`
	CheckCode            string          `json:"CheckCode"`
	MachineCode          string          `json:"MachineCode"`
	InvoiceTypeOrg       string          `json:"InvoiceTypeOrg"`
	CompanySeal          string          `json:"company_seal"`
}

func (b *BdInvoiceResult) ToInvoiceEx(filename string) *InvocieEx {
	invocieEx := &InvocieEx{
		SourceFile:           filename,
		InvoiceTypeOrg:       b.InvoiceTypeOrg,
		MachineCode:          b.MachineCode,
		InvoiceNum:           b.InvoiceNum,
		InvoiceCode:          b.InvoiceCode,
		InvoiceDate:          b.InvoiceDate,
		CheckCode:            b.CheckCode,
		PurchaserName:        b.PurchaserName,
		PurchaserRegisterNum: b.PurchaserRegisterNum,
		PurchaserAddress:     b.PurchaserAddress,
		PurchaserBank:        b.PurchaserBank,
		Password:             b.Password,
		TotalAmount:          b.TotalAmount,
		TotalTax:             b.TotalTax,
		AmountInWords:        b.AmountInWords,
		AmountInFiguers:      b.AmountInFiguers,
		SellerName:           b.SellerName,
		SellerRegisterNum:    b.SellerRegisterNum,
		SellerAddress:        b.SellerAddress,
		SellerBank:           b.SellerBank,
		Remarks:              b.Remarks,
		Payee:                b.Payee,
		Checker:              b.Checker,
		NoteDrawer:           b.NoteDrawer,
		CompanySeal:          b.CompanySeal,
		Province:             b.Province,
		City:                 b.City,
		Agent:                b.Agent,
		InvoiceType:          b.InvoiceType,
		ServiceType:          b.ServiceType,
		CommodityName:        getStrings(b.CommodityName),
		CommodityType:        getStrings(b.CommodityType),
		CommodityUnit:        getStrings(b.CommodityUnit),
		CommodityNum:         getStrings(b.CommodityNum),
		CommodityPrice:       getStrings(b.CommodityPrice),
		CommodityAmount:      getStrings(b.CommodityAmount),
		CommodityTaxRate:     getStrings(b.CommodityTaxRate),
		CommodityTax:         getStrings(b.CommodityTax),
	}
	return invocieEx
}

func getStrings(items []BdInvoiceItem) string {
	ss := []string{}
	for _, item := range items {
		ss = append(ss, item.Word)
	}
	return strings.Join(ss, ",")
}

type RespBdWords struct {
	LogID          int64         `json:"log_id"`
	ErrMsg         string        `json:"error_msg"`
	ErrCode        int64         `json:"error_code"`
	WordsResultNum int64         `json:"words_result_num"`
	WordsResult    []BdWordsItem `json:"words_result"`
	SourceFile     string        `json:"fileName"`
}

type BdWordsItem struct {
	Words string `json:"words"`
}

func (b *RespBdWords) ToVinEx(filename string, code string) *VinEx {
	vinEx := &VinEx{
		SourceFile: filename,
		VinCode:    code,
	}
	return vinEx
}

type RespBdRespTable struct {
	LogID       int64           `json:"log_id"`
	ErrMsg      string          `json:"error_msg"`
	ErrCode     int64           `json:"error_code"`
	TableNum    int64           `json:"table_num"`
	TableResult RespTableResult `json:"tables_result"`
}

type RespTableResult struct {
	Body []RespTableCell `json:"body"`
}

type RespTableCell struct {
	ColStart int    `json:"col_start"`
	RowStart int    `json:"row_start"`
	Words    string `json:"words"`
}

func BdExtractItinerary(filename string, r *RespBdWords) *ItineraryEx {
	ex := &ItineraryEx{SourceFile: filename}
	words := []string{}
	for _, item := range r.WordsResult {
		words = append(words, item.Words)
	}
	text := strings.Join(words, "--")
	inTitle := regexp.MustCompile(`发票抬头：(.+)--`)
	ex.InvocieTitle = getTextByRe(text, inTitle)
	taxNo := regexp.MustCompile(`税号：(.+)--`)
	ex.TaxNo = getTextByRe(text, taxNo)
	code := regexp.MustCompile(`发票代码：(.+)--`)
	ex.InvocieCode = getTextByRe(text, code)
	date := regexp.MustCompile(`开票时间：(.+)--`)
	ex.InvocieDate = getTextByRe(text, date)
	num := regexp.MustCompile(`发票号码：(.+)--`)
	ex.InvocieNum = getTextByRe(text, num)
	amount := regexp.MustCompile(`发票金额：(.+)--`)
	ex.InvocieAmount = getTextByRe(text, amount)
	return ex
}

func getTextByRe(text string, reg *regexp.Regexp) string {
	matcher := reg.FindAllString(text, -1)
	if len(matcher) > 0 {
		return matcher[0]
	}
	return ""
}
