package model

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
	"time"
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
	Words    string         `json:"words"`
	Location BdItemLocation `json:"location"`
}

type BdItemLocation struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (b *RespBdWords) ToVinEx(filename string, code string) *VinEx {
	vinEx := &VinEx{
		SourceFile: filename,
		VinCode:    code,
	}
	return vinEx
}

type RespBdRespTable struct {
	LogID       int64             `json:"log_id"`
	ErrMsg      string            `json:"error_msg"`
	ErrCode     int64             `json:"error_code"`
	TableNum    int64             `json:"table_num"`
	TableResult []RespTableResult `json:"tables_result"`
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
	// 发票抬头
	inTitle := regexp.MustCompile(`头：(.+?)--`)
	ex.InvocieTitle = getTextByRe(text, inTitle)
	taxNo := regexp.MustCompile(`号：(.+?)--`)
	ex.TaxNo = getTextByRe(text, taxNo)
	code := regexp.MustCompile(`发票代码：(.+?)--`)
	ex.InvocieCode = getTextByRe(text, code)
	date := regexp.MustCompile(`开票时间：(.+?)--`)
	ex.InvocieDate = clearDate(getTextByRe(text, date))
	num := regexp.MustCompile(`发票号码：(.+?)--`)
	ex.InvocieNum = getTextByRe(text, num)
	amount := regexp.MustCompile(`发票金额：(.+?)--`)
	ex.InvocieAmount = getTextByRe(text, amount)
	return ex
}

func getTextByRe(text string, reg *regexp.Regexp) string {
	matcher := reg.FindStringSubmatch(text)
	if len(matcher) > 1 {
		return strings.TrimSpace(strings.ReplaceAll(matcher[1], "--", ""))
	}
	return ""
}

func clearDate(date string) string {
	if date != "" {
		replace := regexp.MustCompile(`[-:：]`)
		match := replace.ReplaceAllString(date, "")
		if len(match) == 12 {
			t, err := time.Parse("200601021504", match)
			if err != nil {
				return ""
			}
			return t.Format("2006-01-02 15:04")
		} else if len(match) == 14 {
			t, err := time.Parse("20060102150405", match)
			if err != nil {
				return ""
			}
			return t.Format("2006-01-02 15:04:05")
		}
	}
	return ""
}

func BdExtractItineraryDetail(ex *ItineraryEx, data *RespBdRespTable) (exs []ItineraryEx) {
	tableCell := data.TableResult[0].Body
	maxCol := 0
	headers := []string{}
	for _, cell := range tableCell {
		if cell.ColStart > maxCol {
			maxCol = cell.ColStart
		}
		if cell.RowStart > 0 {
			break
		}
		headers = append(headers, cell.Words)
	}
	for i := 1; i < len(tableCell)/(maxCol+1); i++ {
		newEx := *ex
		for j := 0; j < len(headers); j++ {
			value := tableCell[i*(maxCol+1)+j].Words
			switch headers[j] {
			case "序号":
				newEx.No = value
			case "车牌号码":
				newEx.CarNo = value
			case "入口时间":
				newEx.EnterDate = clearDate(value)
			case "入口站":
				newEx.EnterStation = value
			case "出口站":
				newEx.OutStation = value
			case "出口时间":
				newEx.OutDate = clearDate(value)
			case "金额":
				newEx.TradeAmount = value
			default:
			}

		}
		if newEx.No == "" {
			newEx.No = fmt.Sprint(i)
		}
		exs = append(exs, newEx)
	}
	return
}

func MatchBdGeneral(ex *ItineraryEx, data *RespBdWords) (exs []ItineraryEx) {
	headers := map[int]string{}
	topCount := map[int]int{}
	for _, item := range data.WordsResult {
		if topCount[item.Location.Top] > 0 {
			topCount[item.Location.Top] = topCount[item.Location.Top] + 1
		} else {
			topCount[item.Location.Top] = 1
		}
	}
	// 相邻高度误差3px进行合并
	keys := make([]int, 0, len(topCount))
	for k := range topCount {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for index, key := range keys {
		if index < len(keys)-1 && keys[index+1]-key <= 3 {
			topCount[key] = topCount[key] + topCount[keys[index+1]]
		}
	}
	tops := []int{}
	for top, count := range topCount {
		if count >= 4 {
			tops = append(tops, top)
		}
	}
	sort.Ints(tops)
	for row, top := range tops {
		var newEx ItineraryEx
		if row > 0 {
			newEx = *ex
		}
		for _, item := range data.WordsResult {
			if row == 0 && topEqual(top, item.Location.Top) {
				headers[item.Location.Left] = item.Words
				if len(headers) >= 5 {
					break
				}
			} else if topEqual(top, item.Location.Top) {
				key := getHeader(item.Location.Left, headers)
				if key != "" {
					value := item.Words
					switch key {
					case "入口时间":
						newEx.EnterDate = clearDate(value)
					case "入口站":
						newEx.EnterStation = value
					case "出口站":
						newEx.OutStation = value
					case "出口时间":
						newEx.OutDate = clearDate(value)
					case "交易金额":
						newEx.TradeAmount = value
					default:
					}
				}
			}
		}
		if row > 0 {
			newEx.No = fmt.Sprint(row)
			exs = append(exs, newEx)
		}
	}
	return
}

func getHeader(left int, header map[int]string) string {
	keys := []int{left - 2, left - 1, left, left + 1, left + 2}
	for _, key := range keys {
		if header[key] != "" {
			return header[key]
		}
	}
	return ""
}

func topEqual(top1, top2 int) bool {
	return math.Abs(float64(top1-top2)) <= 3
}
