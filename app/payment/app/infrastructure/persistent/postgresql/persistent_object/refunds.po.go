package persistentobject

import "time"

type Refund struct {
	BasePO
	PaymentID   string `gorm:"type:uuid;not null"`
	Amount      int64  `gorm:"not null"`
	Reason      *string
	Status      *string   `gorm:"size:20"`
	RequestedAt time.Time `gorm:"autoCreateTime"`
	ProcessedAt *time.Time
}

func (r *Refund) TableName() string {
	return "refunds"
}
