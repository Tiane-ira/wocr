package model

import "strings"

type RespAliInvoice struct {
	Data *RespAliInvoiceData `json:"data"`
}

type RespAliInvoiceData struct {
	SellerBankAccountInfo    string                 `json:"sellerBankAccountInfo"`
	PurchaserBankAccountInfo string                 `json:"purchaserBankAccountInfo"`
	SellerName               string                 `json:"sellerName"`
	InvoiceTax               string                 `json:"invoiceTax"`
	PasswordArea             string                 `json:"passwordArea"`
	Title                    string                 `json:"title"`
	PrintedInvoiceNumber     string                 `json:"printedInvoiceNumber"`
	TotalAmountInWords       string                 `json:"totalAmountInWords"`
	InvoiceNumber            string                 `json:"invoiceNumber"`
	InvoiceType              string                 `json:"invoiceType"`
	InvoiceDetails           []RespAliInvoiceDetail `json:"invoiceDetails"`
	PurchaserContactInfo     string                 `json:"purchaserContactInfo"`
	FormType                 string                 `json:"formType"`
	MachineCode              string                 `json:"machineCode"`
	SpecialTag               string                 `json:"specialTag"`
	PrintedInvoiceCode       string                 `json:"printedInvoiceCode"`
	Drawer                   string                 `json:"drawer"`
	Reviewer                 string                 `json:"reviewer"`
	InvoiceDate              string                 `json:"invoiceDate"`
	PurchaserTaxNumber       string                 `json:"purchaserTaxNumber"`
	InvoiceCode              string                 `json:"invoiceCode"`
	PurchaserName            string                 `json:"purchaserName"`
	CheckCode                string                 `json:"checkCode"`
	TotalAmount              string                 `json:"totalAmount"`
	SellerContactInfo        string                 `json:"sellerContactInfo"`
	InvoiceAmountPreTax      string                 `json:"invoiceAmountPreTax"`
	Recipient                string                 `json:"recipient"`
	SellerTaxNumber          string                 `json:"sellerTaxNumber"`
	Remarks                  string                 `json:"remarks"`
}

type RespAliInvoiceDetail struct {
	UnitPrice     string `json:"unitPrice"`
	TaxRate       string `json:"taxRate"`
	ItemName      string `json:"itemName"`
	Unit          string `json:"unit"`
	Amount        string `json:"amount"`
	Quantity      string `json:"quantity"`
	Specification string `json:"specification"`
	Tax           string `json:"tax"`
}

func (r *RespAliInvoiceData) AliToInvoiceEx(filename string) *InvocieEx {
	names, types, units, nums, prices, amounts, rates, taxs := getDetails(r.InvoiceDetails)
	invoiceEx := &InvocieEx{
		SourceFile:           filename,
		InvoiceTypeOrg:       r.Title,
		MachineCode:          r.MachineCode,
		InvoiceCode:          r.InvoiceCode,
		InvoiceNum:           r.InvoiceNumber,
		InvoiceDate:          r.InvoiceDate,
		CheckCode:            r.CheckCode,
		PurchaserName:        r.PurchaserName,
		PurchaserRegisterNum: r.PurchaserTaxNumber,
		PurchaserBank:        r.PurchaserBankAccountInfo,
		PurchaserAddress:     r.PurchaserContactInfo,
		Password:             r.PasswordArea,
		TotalAmount:          r.TotalAmount,
		TotalTax:             r.InvoiceTax,
		AmountInWords:        r.TotalAmountInWords,
		SellerName:           r.SellerName,
		SellerRegisterNum:    r.SellerTaxNumber,
		SellerAddress:        r.SellerContactInfo,
		SellerBank:           r.SellerBankAccountInfo,
		Remarks:              r.Remarks,
		Payee:                r.Recipient,
		Checker:              r.Reviewer,
		NoteDrawer:           r.Drawer,
		InvoiceType:          r.InvoiceType,
		CommodityName:        strings.Join(names, ","),
		CommodityType:        strings.Join(types, ","),
		CommodityUnit:        strings.Join(units, ","),
		CommodityNum:         strings.Join(nums, ","),
		CommodityPrice:       strings.Join(prices, ","),
		CommodityAmount:      strings.Join(amounts, ","),
		CommodityTaxRate:     strings.Join(rates, ","),
		CommodityTax:         strings.Join(taxs, ","),
	}

	return invoiceEx
}

func getDetails(vatInvoiceGoodsInfo []RespAliInvoiceDetail) (names, types, units, nums, prices, amounts, rates, taxs []string) {
	for _, item := range vatInvoiceGoodsInfo {
		names = append(names, item.ItemName)
		types = append(types, item.Specification)
		units = append(units, item.Unit)
		nums = append(nums, item.Quantity)
		prices = append(prices, item.UnitPrice)
		amounts = append(amounts, item.Amount)
		rates = append(rates, item.TaxRate)
		taxs = append(taxs, item.Tax)
	}
	return
}
