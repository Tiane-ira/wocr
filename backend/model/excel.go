package model

type InvocieEx struct {
	SourceFile           string `excel:"name:文件名;"`
	InvoiceTypeOrg       string `excel:"name:发票名称;"`
	MachineCode          string `excel:"name:机器编号;"`
	InvoiceCode          string `excel:"name:发票代码;"`
	InvoiceNum           string `excel:"name:发票号码;"`
	InvoiceDate          string `excel:"name:开票日期;"`
	CheckCode            string `excel:"name:校验码;"`
	PurchaserName        string `excel:"name:购方名称;"`
	PurchaserRegisterNum string `excel:"name:购方纳税人识别号;"`
	PurchaserAddress     string `excel:"name:购方地址及电话;"`
	PurchaserBank        string `excel:"name:购方开户行及账号;"`
	Password             string `excel:"name:密码区;"`
	TotalAmount          string `excel:"name:合计金额;"`
	TotalTax             string `excel:"name:合计税额;"`
	AmountInWords        string `excel:"name:价税合计(大写);"`
	AmountInFiguers      string `excel:"name:价税合计(小写);"`
	SellerName           string `excel:"name:销售方名称;"`
	SellerRegisterNum    string `excel:"name:销售方纳税人识别号;"`
	SellerAddress        string `excel:"name:销售方地址及电话;"`
	SellerBank           string `excel:"name:销售方开户行及账号;"`
	Remarks              string `excel:"name:备注;"`
	Payee                string `excel:"name:收款人;"`
	Checker              string `excel:"name:复核;"`
	NoteDrawer           string `excel:"name:开票人;"`
	CompanySeal          string `excel:"name:是否有公司印章;"`
	Province             string `excel:"name:省;"`
	City                 string `excel:"name:市;"`
	Agent                string `excel:"name:是否代开;"`
	InvoiceType          string `excel:"name:发票种类;"`
	ServiceType          string `excel:"name:发票消费类型;"`
	CommodityName        string `excel:"name:货物名称;"`
	CommodityType        string `excel:"name:规格型号;"`
	CommodityUnit        string `excel:"name:单位;"`
	CommodityNum         string `excel:"name:数量;"`
	CommodityPrice       string `excel:"name:单价;"`
	CommodityAmount      string `excel:"name:金额;"`
	CommodityTaxRate     string `excel:"name:税率;"`
	CommodityTax         string `excel:"name:税额;"`
}

type VinEx struct {
	SourceFile string `excel:"name:文件名;"`
	VinCode    string `excel:"name:Vin码;"`
}

func NewVinEX(filename, code string) *VinEx {
	return &VinEx{
		SourceFile: filename,
		VinCode:    code,
	}
}

type ItineraryEx struct {
	SourceFile    string `excel:"name:文件名;"`
	InvocieTitle  string `excel:"name:发票抬头;"`
	TaxNo         string `excel:"name:税号;"`
	InvocieCode   string `excel:"name:发票代码;"`
	InvocieNum    string `excel:"name:发票号码;"`
	InvocieDate   string `excel:"name:开票时间;"`
	InvocieAmount string `excel:"name:发票金额;"`
	No            string `excel:"name:序号;"`
	CarNo         string `excel:"name:车牌号码;"`
	EnterDate     string `excel:"name:入口时间;"`
	EnterStation  string `excel:"name:入口站;"`
	OutDate       string `excel:"name:出口时间;"`
	OutStation    string `excel:"name:出口站;"`
	TradeAmount   string `excel:"name:交易金额;"`
}

type CarNoEx struct {
	SourceFile string `excel:"name:文件名;"`
	CarNo      string `excel:"name:车牌号;"`
}

func NewCarNoEx(filename, carNo string) *CarNoEx {
	return &CarNoEx{
		SourceFile: filename,
		CarNo:      carNo,
	}
}
