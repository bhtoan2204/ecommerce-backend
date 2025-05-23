package persistentobject

import "time"

type Settlement struct {
	BasePO
	SellerID  string `gorm:"type:uuid;not null"`
	PaymentID string `gorm:"type:uuid;not null"`
	Amount    int64  `gorm:"not null"`
	Status    string `gorm:"size:20"`
	SettledAt *time.Time
}

func (s *Settlement) TableName() string {
	return "settlements"
}
