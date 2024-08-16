package model

import (
	"context"
	"fmt"
	"wocr/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func Init(ctx context.Context) {
	path := utils.GetProjectPath()
	dbPath := fmt.Sprintf("%s/wocr.db", path)
	Db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&SkConfig{}, &ExportField{})

	fields := []*ExportField{
		{Id: 1, Name: "文件名", FieldName: "SourceFile", Export: true},
		{Id: 2, Name: "发票名称", FieldName: "InvoiceTypeOrg", Export: true},
		{Id: 3, Name: "机器编号", FieldName: "MachineCode", Export: true},
		{Id: 4, Name: "发票代码", FieldName: "InvoiceCode", Export: true},
		{Id: 5, Name: "发票号码", FieldName: "InvoiceNum", Export: true},
		{Id: 6, Name: "开票日期", FieldName: "InvoiceDate", Export: true},
		{Id: 7, Name: "校验码", FieldName: "CheckCode", Export: true},
		{Id: 8, Name: "购方名称", FieldName: "PurchaserName", Export: true},
		{Id: 9, Name: "购方纳税人识别号", FieldName: "PurchaserRegisterNum", Export: true},
		{Id: 10, Name: "购方地址及电话", FieldName: "PurchaserAddress", Export: true},
		{Id: 11, Name: "购方开户行及账号", FieldName: "PurchaserBank", Export: true},
		{Id: 12, Name: "密码区", FieldName: "Password", Export: true},
		{Id: 13, Name: "合计金额", FieldName: "TotalAmount", Export: true},
		{Id: 14, Name: "合计税额", FieldName: "TotalTax", Export: true},
		{Id: 15, Name: "价税合计(大写)", FieldName: "AmountInWords", Export: true},
		{Id: 16, Name: "价税合计(小写)", FieldName: "AmountInFiguers", Export: true},
		{Id: 17, Name: "销售方名称", FieldName: "SellerName", Export: true},
		{Id: 18, Name: "销售方纳税人识别号", FieldName: "SellerRegisterNum", Export: true},
		{Id: 19, Name: "销售方地址及电话", FieldName: "SellerAddress", Export: true},
		{Id: 20, Name: "销售方开户行及账号", FieldName: "SellerBank", Export: true},
		{Id: 21, Name: "备注", FieldName: "Remarks", Export: true},
		{Id: 22, Name: "收款人", FieldName: "Payee", Export: true},
		{Id: 23, Name: "复核", FieldName: "Checker", Export: true},
		{Id: 24, Name: "开票人", FieldName: "NoteDrawer", Export: true},
		{Id: 25, Name: "是否有公司印章", FieldName: "CompanySeal", Export: false},
		{Id: 26, Name: "省", FieldName: "Province", Export: false},
		{Id: 27, Name: "市", FieldName: "City", Export: false},
		{Id: 28, Name: "是否代开", FieldName: "Agent", Export: false},
		{Id: 29, Name: "发票种类", FieldName: "InvoiceType", Export: true},
		{Id: 30, Name: "发票消费类型", FieldName: "ServiceType", Export: false},
		{Id: 31, Name: "货物名称", FieldName: "CommodityName", Export: true},
		{Id: 32, Name: "规格型号", FieldName: "CommodityType", Export: false},
		{Id: 33, Name: "单位", FieldName: "CommodityUnit", Export: false},
		{Id: 34, Name: "数量", FieldName: "CommodityNum", Export: false},
		{Id: 35, Name: "单价", FieldName: "CommodityPrice", Export: false},
		{Id: 36, Name: "金额", FieldName: "CommodityAmount", Export: false},
		{Id: 37, Name: "税率", FieldName: "CommodityTaxRate", Export: false},
		{Id: 38, Name: "税额", FieldName: "CommodityTax", Export: false},
	}
	Db.Create(fields)
}
