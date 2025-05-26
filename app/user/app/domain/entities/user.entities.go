package entities

import (
	"errors"
	"strings"
	"time"
	"user/app/domain/value_object"
	"user/package/xtypes"
)

type User struct {
	id            int64
	email         string
	password      value_object.Password
	passwordHash  string
	firstName     string
	lastName      string
	phoneNumber   string
	birthDate     *time.Time
	addresses     []*Address
	avatar        string
	pinCode       string
	tier          xtypes.UserTier
	role          xtypes.UserRole
	isActive      bool
	emailVerified bool
}

func DefaultUser() *User {
	return &User{}
}

func NewUser(
	email string,
	passwordHash string,
	firstName string,
	lastName string,
	phoneNumber string,
	birthDate *time.Time,
	addresses []*Address,
	avatar string,
	pinCode string,
	role xtypes.UserRole,
) *User {
	user := &User{
		email:         email,
		passwordHash:  passwordHash,
		firstName:     firstName,
		lastName:      lastName,
		phoneNumber:   phoneNumber,
		birthDate:     birthDate,
		addresses:     addresses,
		avatar:        avatar,
		pinCode:       pinCode,
		tier:          xtypes.TierBronze,
		role:          role,
		isActive:      false,
		emailVerified: false,
	}

	return user
}

func NewUserWithID(
	id int64,
	email string,
	passwordHash string,
	firstName string,
	lastName string,
	phoneNumber string,
	birthDate *time.Time,
	addresses []*Address,
	avatar string,
	pinCode string,
	role xtypes.UserRole,
) *User {
	user := &User{
		id:            id,
		email:         email,
		passwordHash:  passwordHash,
		firstName:     firstName,
		lastName:      lastName,
		phoneNumber:   phoneNumber,
		birthDate:     birthDate,
		addresses:     addresses,
		avatar:        avatar,
		pinCode:       pinCode,
		tier:          xtypes.TierBronze,
		role:          role,
		isActive:      false,
		emailVerified: false,
	}

	return user
}

// Getters
func (u *User) ID() int64 {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() value_object.Password {
	return u.password
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) PhoneNumber() string {
	return u.phoneNumber
}

func (u *User) BirthDate() *time.Time {
	return u.birthDate
}

func (u *User) Address() []*Address {
	address := make([]*Address, len(u.addresses))
	copy(address, u.addresses)
	return address
}

func (u *User) Avatar() string {
	return u.avatar
}

func (u *User) PinCode() string {
	return u.pinCode
}

func (u *User) Tier() xtypes.UserTier {
	return u.tier
}

func (u *User) Role() xtypes.UserRole {
	return u.role
}

func (u *User) IsActive() bool {
	return u.isActive
}

func (u *User) IsEmailVerified() bool {
	return u.emailVerified
}

func (u *User) DefaultAddress() *Address {
	for _, addr := range u.addresses {
		if addr.IsDefault() {
			return addr
		}
	}
	return nil
}

// Setters
func (u *User) SetEmail(email string) error {
	if email == "" || !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	u.email = email
	return nil
}

func (u *User) SetPasswordHash(hash string) {
	u.passwordHash = hash
}

func (u *User) SetPassword(password value_object.Password) {
	u.password = password
}

func (u *User) SetFirstName(name string) {
	u.firstName = name
}

func (u *User) SetLastName(name string) {
	u.lastName = name
}

func (u *User) SetPhoneNumber(phone string) {
	u.phoneNumber = phone
}

func (u *User) SetBirthDate(birth *time.Time) {
	u.birthDate = birth
}

func (u *User) SetAddress(addr []*Address) {
	u.addresses = addr
}

func (u *User) SetAvatar(avatar string) {
	u.avatar = avatar
}

func (u *User) SetPinCode(pin string) {
	u.pinCode = pin
}

func (u *User) SetTier(tier xtypes.UserTier) {
	u.tier = tier
}

func (u *User) SetRole(role xtypes.UserRole) {
	u.role = role
}

func (u *User) Validate() error {
	if u.email == "" || !strings.Contains(u.email, "@") {
		return errors.New("invalid email")
	}
	if u.passwordHash == "" {
		return errors.New("password hash is required")
	}
	if u.firstName == "" || u.lastName == "" {
		return errors.New("first name and last name cannot be empty")
	}
	if u.birthDate == nil {
		return errors.New("birth date is required")
	}
	return nil
}

func (u *User) VerifyEmail() {
	u.emailVerified = true
}

func (u *User) Activate() {
	u.isActive = true
}

func (u *User) Deactivate() {
	u.isActive = false
}

func (u *User) UpgradeTier() {
	switch u.tier {
	case xtypes.TierBronze:
		u.tier = xtypes.TierSilver
	case xtypes.TierSilver:
		u.tier = xtypes.TierGold
	case xtypes.TierGold:
		u.tier = xtypes.TierPlatinum
	}
}

func (u *User) IsAdmin() bool {
	return u.role == xtypes.RoleAdmin
}

func (u *User) IsSeller() bool {
	return u.role == xtypes.RoleSeller
}

func (u *User) IsCustomer() bool {
	return u.role == xtypes.RoleCustomer
}
