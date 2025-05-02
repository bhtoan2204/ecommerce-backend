package entities

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

// business logic here
