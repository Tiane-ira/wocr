package model

type LocalOcr struct {
	Result string `json:"result"`
	Err    string `json:"err"`
}

func (l *LocalOcr) ToVinEx(filename string, code string) *VinEx {
	return &VinEx{
		SourceFile: filename,
		VinCode:    code,
	}
}
