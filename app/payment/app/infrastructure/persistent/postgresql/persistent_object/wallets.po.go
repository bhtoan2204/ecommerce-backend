package persistentobject

type Wallet struct {
	BasePO
	OwnerType string `gorm:"size:20"` // 'seller', 'platform'
	OwnerID   string `gorm:"type:uuid;not null"`
	Balance   int64  `gorm:"default:0"`
	Currency  string `gorm:"size:10;default:'VND'"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}
