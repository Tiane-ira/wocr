package model

import (
	"strings"

	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

func TencentPdfToInvoiceEx(filename string, r *ocr.VatInvoiceOCRResponseParams) *InvocieEx {
	infoMap := map[string]string{}
	for _, info := range r.VatInvoiceInfos {
		infoMap[*info.Name] = *info.Value
	}
	var names, types, units, nums, prices, amounts, rates, taxs []string
	for _, item := range r.Items {
		names = append(names, *item.Name)
		types = append(types, *item.Spec)
		units = append(units, *item.Unit)
		nums = append(nums, *item.Quantity)
		prices = append(prices, *item.UnitPrice)
		amounts = append(amounts, *item.AmountWithoutTax)
		rates = append(rates, *item.TaxRate)
		taxs = append(taxs, *item.TaxAmount)
	}
	invoiceEx := &InvocieEx{
		SourceFile:           filename,
		InvoiceTypeOrg:       infoMap["发票名称"],
		MachineCode:          infoMap["机器编号"],
		InvoiceCode:          infoMap["发票代码"],
		InvoiceNum:           infoMap["发票号码"],
		InvoiceDate:          infoMap["开票日期"],
		CheckCode:            infoMap["校验码"],
		PurchaserName:        infoMap["购买方名称"],
		PurchaserRegisterNum: infoMap["购买方识别号"],
		PurchaserBank:        infoMap["购买方开户行及账号"],
		PurchaserAddress:     infoMap["购买方地址、电话"],
		Password:             infoMap["密码区"],
		TotalAmount:          infoMap["合计金额"],
		TotalTax:             infoMap["合计税额"],
		AmountInWords:        infoMap["价税合计(大写)"],
		SellerName:           infoMap["销售方名称"],
		SellerRegisterNum:    infoMap["销售方识别号"],
		SellerAddress:        infoMap["销售方地址、电话"],
		SellerBank:           infoMap["销售方开户行及账号"],
		Remarks:              infoMap["备注"],
		Payee:                infoMap["收款人"],
		Checker:              infoMap["复核"],
		NoteDrawer:           infoMap["开票人"],
		CompanySeal:          infoMap["是否有公司印章"],
		Province:             infoMap["省"],
		City:                 infoMap["市"],
		Agent:                infoMap["是否代开"],
		InvoiceType:          infoMap["发票类型"],
		ServiceType:          infoMap["发票消费类型"],
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

func TencentOfdToInvoiceEx(filename string, r *ocr.VerifyOfdVatInvoiceOCRResponseParams) *InvocieEx {
	names, types, units, nums, prices, amounts, rates, taxs := getGoodsInfo(r.GoodsInfos)
	invoiceEx := &InvocieEx{
		SourceFile:           filename,
		InvoiceTypeOrg:       *r.InvoiceTitle,
		MachineCode:          *r.MachineNumber,
		InvoiceCode:          *r.InvoiceCode,
		InvoiceNum:           *r.InvoiceNumber,
		InvoiceDate:          *r.IssueDate,
		CheckCode:            *r.InvoiceCheckCode,
		PurchaserName:        *r.Buyer.Name,
		PurchaserRegisterNum: *r.Buyer.TaxId,
		PurchaserBank:        *r.Buyer.FinancialAccount,
		PurchaserAddress:     *r.Buyer.AddrTel,
		Password:             *r.TaxControlCode,
		TotalAmount:          *r.TaxInclusiveTotalAmount,
		TotalTax:             *r.TaxTotalAmount,
		AmountInFiguers:      *r.TaxInclusiveTotalAmount,
		SellerName:           *r.Seller.Name,
		SellerRegisterNum:    *r.Seller.TaxId,
		SellerAddress:        *r.Seller.AddrTel,
		SellerBank:           *r.Seller.FinancialAccount,
		Remarks:              *r.Note,
		Payee:                *r.Payee,
		Checker:              *r.Checker,
		NoteDrawer:           *r.InvoiceClerk,
		InvoiceType:          *r.Type,
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

func getGoodsInfo(vatInvoiceGoodsInfo []*ocr.VatInvoiceGoodsInfo) (names, types, units, nums, prices, amounts, rates, taxs []string) {
	for _, item := range vatInvoiceGoodsInfo {
		names = append(names, *item.Item)
		types = append(types, *item.Specification)
		units = append(units, *item.MeasurementDimension)
		nums = append(nums, *item.Quantity)
		prices = append(prices, *item.Price)
		amounts = append(amounts, *item.Amount)
		rates = append(rates, *item.TaxScheme)
		taxs = append(taxs, *item.TaxAmount)
	}
	return
}

func TencentToVinEx(filename string, code string) *VinEx {
	return &VinEx{
		SourceFile: filename,
		VinCode:    code,
	}
}
