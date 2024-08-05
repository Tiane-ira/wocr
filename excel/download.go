// ------------------------------------------------------------------------
// -------------------           Author：符华            -------------------
// -------------------           Gitee：寒霜剑客          -------------------
// ------------------------------------------------------------------------

package excel

import (
	"github.com/xuri/excelize/v2"
	"html/template"
	"net/http"
	"net/url"
)

// ================================= 下载到浏览器 =================================

// NormalDynamicExport NormalDownLoad 导出并下载（单个sheet）
func NormalDownLoad(fileName, sheet, title string, isGhbj bool, list interface{}, res http.ResponseWriter) error {
	f, err := NormalDynamicExport(sheet, title, "", isGhbj, false, list, nil)
	if err != nil {
		return err
	}
	DownLoadExcel(fileName, res, f)
	return nil
}

// NormalDynamicExport NormalDynamicDownLoad 动态表头导出并下载（单个sheet）
func NormalDynamicDownLoad(fileName, sheet, title, fields string, isGhbj, isIgnore bool, list interface{}, changeHead map[string]string, res http.ResponseWriter) error {
	f, err := NormalDynamicExport(sheet, title, fields, isGhbj, isIgnore, list, changeHead)
	if err != nil {
		return err
	}
	DownLoadExcel(fileName, res, f)
	return nil
}

// CustomHeaderDownLoad 复杂表头导出并下载
func CustomHeaderDownLoad(fileName, sheet, title string, isGhbj bool, heads interface{}, list interface{}, res http.ResponseWriter) error {
	f, err := CustomHeaderExport(sheet, title, isGhbj, heads, list)
	if err != nil {
		return err
	}
	DownLoadExcel(fileName, res, f)
	return nil
}

// MapExport MapExportDownLoad 基于map导出并下载
func MapExportDownLoad(heads interface{}, list []map[string]interface{}, fileName, sheet, title string, isGhbj bool, res http.ResponseWriter) error {
	f, err := MapExport(heads, list, sheet, title, isGhbj)
	if err != nil {
		return err
	}
	DownLoadExcel(fileName, res, f)
	return nil
}

// 下载excel文件
func DownLoadExcel(fileName string, res http.ResponseWriter, file *excelize.File) {
	// 设置响应头
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", "attachment; filename="+url.PathEscape(fileName)+".xlsx")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	err := file.Write(res) // 写入Excel文件内容到响应体
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// 根据模板下载文件：templatePath 模板路径，fileName 文件名称（需要加上后缀名），data 模板数据
func DownLoadByTemplate(templatePath, fileName string, data map[string]interface{}, res http.ResponseWriter) {
	// 解析模板
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(res, "模板解析失败："+err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置响应头
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", "attachment; filename="+url.PathEscape(fileName))
	res.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	// 渲染模板并输出结果
	err = tmpl.Execute(res, data)
	if err != nil {
		http.Error(res, "模板数据渲染失败："+err.Error(), http.StatusInternalServerError)
	}
}
