package model

type ExportField struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	FieldName string `json:"fieldName"`
	Export    bool   `json:"export"`
}

func (e *ExportField) Create() error {
	return Db.Create(e).Error
}

func (e *ExportField) Update(idList []int64) (err error) {
	err = Db.Model(e).Where("id in ?", idList).Update("export", true).Error
	if err != nil {
		return
	}
	err = Db.Model(e).Where("id not in ?", idList).Update("export", false).Error
	return
}

func (e *ExportField) ListAll() (fields []ExportField, err error) {
	err = Db.Model(e).Find(&fields).Error
	return
}

func (e *ExportField) GetExports() (fieldnames []string, err error) {
	fields, err := e.ListAll()
	if err != nil {
		return
	}
	for _, field := range fields {
		if field.Export {
			fieldnames = append(fieldnames, field.FieldName)
		}
	}
	return
}
