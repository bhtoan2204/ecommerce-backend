package model

type Address struct {
	AbstractModel
	UserID    uint   `gorm:"index;not null"`
	Line1     string `gorm:"size:255;not null"`
	Line2     string `gorm:"size:255"`
	City      string `gorm:"size:100;not null"`
	State     string `gorm:"size:100"`
	ZipCode   string `gorm:"size:20;not null"`
	Country   string `gorm:"size:100;not null"`
	IsDefault bool   `gorm:"default:false"`
}

func (a *Address) GetLine1() string {
	return a.Line1
}

func (a *Address) GetLine2() string {
	return a.Line2
}

func (a *Address) GetCity() string {
	return a.City
}

func (a *Address) GetState() string {
	return a.State
}

func (a *Address) GetZipCode() string {
	return a.ZipCode
}

func (a *Address) GetCountry() string {
	return a.Country
}

func (a *Address) GetIsDefault() bool {
	return a.IsDefault
}
