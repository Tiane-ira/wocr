package model

type SkConfig struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Ak   string `json:"ak"`
	Sk   string `json:"sk"`
	Date string `json:"date"`
}

func (s *SkConfig) Create() error {
	return Db.Create(s).Error
}

func (s *SkConfig) GetById() error {
	return Db.Find(s).Error
}

func (s *SkConfig) ListAll() (configs []SkConfig, err error) {
	err = Db.Find(&configs).Error
	return
}

func (s *SkConfig) ListBy(typeName string) (configs []SkConfig, err error) {
	err = Db.Where("type = ?", typeName).Find(&configs).Error
	return
}

func (s *SkConfig) Delete() error {
	return Db.Delete(s).Error
}

func (s *SkConfig) Count() (count int64, err error) {
	err = Db.Model(s).Count(&count).Error
	return
}
