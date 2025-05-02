package entities

import (
	"errors"
	"strings"
	"time"
)

type UserTier string

const (
	TierBronze   UserTier = "bronze"
	TierSilver   UserTier = "silver"
	TierGold     UserTier = "gold"
	TierPlatinum UserTier = "platinum"
)

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleSeller   UserRole = "seller"
	RoleAdmin    UserRole = "admin"
)

type User struct {
	ID            string
	Email         string
	PasswordHash  string
	FirstName     string
	LastName      string
	PhoneNumber   string
	BirthDate     *time.Time
	Address       string
	Avatar        string
	PinCode       string
	Tier          UserTier
	Role          UserRole
	IsActive      bool
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u *User) Validate() error {
	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return errors.New("invalid email")
	}
	if u.PasswordHash == "" {
		return errors.New("password hash is required")
	}
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("first name and last name cannot be empty")
	}
	if u.BirthDate == nil {
		return errors.New("birth date is required")
	}
	return nil
}

func (u *User) VerifyEmail() {
	u.EmailVerified = true
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

func (u *User) UpgradeTier() {
	switch u.Tier {
	case TierBronze:
		u.Tier = TierSilver
	case TierSilver:
		u.Tier = TierGold
	case TierGold:
		u.Tier = TierPlatinum
	}
	u.UpdatedAt = time.Now()
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsSeller() bool {
	return u.Role == RoleSeller
}

func (u *User) IsCustomer() bool {
	return u.Role == RoleCustomer
}
