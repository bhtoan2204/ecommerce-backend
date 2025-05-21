package model

import (
	"time"

	"gorm.io/gorm"
)

type AbstractModel struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
