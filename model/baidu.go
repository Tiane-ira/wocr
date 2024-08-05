package model

import "strings"

type RespBdToken struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int64  `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

type BdItem struct {
	Word string `json:"word"`
	Row  string `json:"row"`
}
type RespBd struct {
	LogID          int64    `json:"log_id"`
	ErrMsg         string   `json:"error_msg"`
	ErrCode        int64    `json:"error_code"`
	WordsResultNum int64    `json:"words_result_num"`
	WordsResult    BdResult `json:"words_result"`
	SourceFile     string   `json:"fileName"`
}

type BdResult struct {
	InvoiceNumDigit      string   `json:"InvoiceNumDigit"`
	ServiceType          string   `json:"ServiceType"`
	InvoiceNum           string   `json:"InvoiceNum"`
	InvoiceNumConfirm    string   `json:"InvoiceNumConfirm"`
	SellerName           string   `json:"SellerName"`
	CommodityTaxRate     []BdItem `json:"CommodityTaxRate"`
	SellerBank           string   `json:"SellerBank"`
	Checker              string   `json:"Checker"`
	TotalAmount          string   `json:"TotalAmount"`
	CommodityAmount      []BdItem `json:"CommodityAmount"`
	InvoiceDate          string   `json:"InvoiceDate"`
	CommodityTax         []BdItem `json:"CommodityTax"`
	PurchaserName        string   `json:"PurchaserName"`
	CommodityNum         []BdItem `json:"CommodityNum"`
	Province             string   `json:"Province"`
	City                 string   `json:"City"`
	SheetNum             string   `json:"SheetNum"`
	Agent                string   `json:"Agent"`
	PurchaserBank        string   `json:"PurchaserBank"`
	Remarks              string   `json:"Remarks"`
	Password             string   `json:"Password"`
	SellerAddress        string   `json:"SellerAddress"`
	PurchaserAddress     string   `json:"PurchaserAddress"`
	InvoiceCode          string   `json:"InvoiceCode"`
	InvoiceCodeConfirm   string   `json:"InvoiceCodeConfirm"`
	CommodityUnit        []BdItem `json:"CommodityUnit"`
	Payee                string   `json:"Payee"`
	PurchaserRegisterNum string   `json:"PurchaserRegisterNum"`
	CommodityPrice       []BdItem `json:"CommodityPrice"`
	NoteDrawer           string   `json:"NoteDrawer"`
	AmountInWords        string   `json:"AmountInWords"`
	AmountInFiguers      string   `json:"AmountInFiguers"`
	TotalTax             string   `json:"TotalTax"`
	InvoiceType          string   `json:"InvoiceType"`
	SellerRegisterNum    string   `json:"SellerRegisterNum"`
	CommodityName        []BdItem `json:"CommodityName"`
	CommodityType        []BdItem `json:"CommodityType"`
	CommodityPlateNum    []BdItem `json:"CommodityPlateNum"`
	CommodityVehicleType []BdItem `json:"CommodityVehicleType"`
	CommodityStartDate   []BdItem `json:"CommodityStartDate"`
	CommodityEndDate     []BdItem `json:"CommodityEndDate"`
	OnlinePay            string   `json:"OnlinePay"`
	CheckCode            string   `json:"CheckCode"`
	MachineCode          string   `json:"MachineCode"`
	InvoiceTypeOrg       string   `json:"InvoiceTypeOrg"`
	CompanySeal          string   `json:"company_seal"`
}

func (b *BdResult) ToInvoiceEx(filename string) InvocieEx {
	invocieEx := InvocieEx{
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

func getStrings(items []BdItem) string {
	ss := []string{}
	for _, item := range items {
		ss = append(ss, item.Word)
	}
	return strings.Join(ss, ",")
}
