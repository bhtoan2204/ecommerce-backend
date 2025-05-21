package model

import (
	"time"
	"user/package/xtypes"
)

type User struct {
	AbstractModel
	Email         string
	Password      string
	FirstName     string
	LastName      string
	PhoneNumber   string
	BirthDate     *time.Time
	Avatar        string
	PinCode       string
	Tier          xtypes.UserTier
	Role          xtypes.UserRole
	IsActive      bool
	EmailVerified bool

	Addresses []Address
}
