package model

type Address struct {
	AbstractModel
	UserID    int64
	Line1     string
	Line2     string
	City      string
	State     string
	ZipCode   string
	Country   string
	IsDefault bool
}
