package persistentobject

import (
	"time"
	"user/package/xtypes"
)

type User struct {
	BasePO
	Email         string `gorm:"uniqueIndex;size:100;not null"`
	Password      string `gorm:"size:255;not null"` // hashed
	FirstName     string `gorm:"size:100"`
	LastName      string `gorm:"size:100"`
	PhoneNumber   string `gorm:"size:20"`
	BirthDate     *time.Time
	Avatar        string          `gorm:"size:255"`
	PinCode       string          `gorm:"size:20"`
	Tier          xtypes.UserTier `gorm:"type:varchar(20);default:'bronze'"`
	Role          xtypes.UserRole `gorm:"type:varchar(20);default:'customer'"`
	IsActive      bool            `gorm:"default:true"`
	EmailVerified bool            `gorm:"default:false"`

	Addresses []Address `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}
