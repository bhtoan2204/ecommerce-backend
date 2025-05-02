package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AbstractModel struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (u *AbstractModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
