package persistent_object

import "time"

type Session struct {
	BasePO
	UserID       int64     `gorm:"index;not null" json:"user_id"`
	RefreshToken string    `gorm:"size:100;not null" json:"refresh_token"`
	UserAgent    string    `gorm:"size:100" json:"user_agent"`
	IPAddress    string    `gorm:"type:varchar(45)" json:"ip_address"`
	IsRevoked    bool      `gorm:"default:false" json:"is_revoked"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
}

func (Session) TableName() string {
	return "sessions"
}
