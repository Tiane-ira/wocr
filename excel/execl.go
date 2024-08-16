package excel

import (
	"strings"
)

// 测试结构体
type Test struct {
	Id       string `excel:"name:用户账号;"`
	Name     string `excel:"name:用户姓名;style:Alignment{Horizontal:left},Font{Color:#ff0000,Size:16.0,Family:黑体},Fill{Color:#f9cb9c,Type:pattern,Pattern:1};"`
	Email    string `excel:"name:用户邮箱;width:25;"`
	Com      string `excel:"name:所属公司;"`
	Dept     bool   `excel:"name:所在部门;replace:false_超级管理员,true_普通用户;"`
	RoleName string `excel:"name:角色名称;replace:1_超级管理员,2_普通用户;"`
	Remark   int    `excel:"name:备注;replace:1_超级管理员,2_普通用户;width:40;"`
}

func Export(savePath string, fields []string, dataList interface{}) error {
	exportField := ""
	if len(fields) > 0 {
		exportField = strings.Join(fields, ",") + ","
	}
	f, err := NormalDynamicExport("Sheet1", "", exportField, false, false, dataList, nil)
	if err != nil {
		return err
	}
	f.Path = savePath
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
