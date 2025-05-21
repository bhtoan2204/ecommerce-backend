package model

import (
	"time"
)

type Session struct {
	AbstractModel
	UserID       int64
	RefreshToken string
	UserAgent    string
	IPAddress    string
	IsRevoked    bool
	ExpiresAt    time.Time
}
