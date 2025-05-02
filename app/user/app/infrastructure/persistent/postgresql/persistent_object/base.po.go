package persistentobject

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasePO struct {
	ID        string `json:"id" gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (b *BasePO) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
