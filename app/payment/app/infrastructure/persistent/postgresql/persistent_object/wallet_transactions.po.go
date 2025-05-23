package persistentobject

type WalletTransaction struct {
	BasePO
	WalletID         *string `gorm:"type:uuid"`
	Type             string  `gorm:"size:10;not null"` // credit | debit
	Amount           int64   `gorm:"not null"`
	Reason           *string
	RelatedPaymentID *string `gorm:"type:uuid"`
}

func (w *WalletTransaction) TableName() string {
	return "wallet_transactions"
}
