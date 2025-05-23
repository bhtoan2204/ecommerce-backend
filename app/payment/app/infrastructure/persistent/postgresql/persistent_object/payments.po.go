package persistentobject

import "time"

type Payment struct {
	BasePO
	OrderID       string `gorm:"type:uuid;not null"`
	BuyerID       string `gorm:"type:uuid;not null"`
	SellerID      string `gorm:"type:uuid;not null"`
	Method        string `gorm:"size:50"`
	Provider      string `gorm:"size:50;default:'internal'"`
	Amount        int64  `gorm:"not null"`
	Currency      string `gorm:"size:10;default:'VND'"`
	Status        string `gorm:"size:20;not null"`
	FailureReason *string
	PaidAt        *time.Time
	RefundedAt    *time.Time
}

func (p *Payment) TableName() string {
	return "payments"
}
