package entities

import (
	"errors"
	"time"
)

type Status int

const (
	InActive Status = 0
	Active   Status = 1
	Deleted  Status = 2
)

type User struct {
	AbstractModel
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	BirthDate    *time.Time `json:"birth_date,omitempty"`
	Address      string     `json:"address,omitempty"`
	PasswordHash string     `json:"password_hash"`
	Avatar       string     `json:"avatar,omitempty"`
	PinCode      string     `json:"pin_code,omitempty"`
	Status       Status     `json:"status"` // Default to Active
	// Roles        []*Role    `json:"roles,omitempty"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("first name and last name cannot be empty")
	}
	if u.BirthDate == nil {
		return errors.New("birth date is required")
	}

	return nil
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetPhone() string {
	return u.Phone
}

func (u *User) GetBirthDate() *time.Time {
	return u.BirthDate
}

func (u *User) GetAddress() string {
	return u.Address
}

func (u *User) GetPasswordHash() string {
	return u.PasswordHash
}

func (u *User) GetAvatar() string {
	return u.Avatar
}

func (u *User) GetPinCode() string {
	return u.PinCode
}

func (u *User) GetStatus() Status {
	return u.Status
}
