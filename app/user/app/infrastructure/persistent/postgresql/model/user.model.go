package model

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
	AbstractModel
	Email         string   `gorm:"uniqueIndex;size:100;not null"`
	Password      string   `gorm:"size:255;not null"` // hashed
	FirstName     string   `gorm:"size:100"`
	LastName      string   `gorm:"size:100"`
	PhoneNumber   string   `gorm:"size:20"`
	Tier          UserTier `gorm:"type:varchar(20);default:'bronze'"`
	Role          UserRole `gorm:"type:varchar(20);default:'customer'"`
	IsActive      bool     `gorm:"default:true"`
	EmailVerified bool     `gorm:"default:false"`

	ShippingAddresses []Address `gorm:"foreignKey:UserID"`
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetPhoneNumber() string {
	return u.PhoneNumber
}

func (u *User) GetTier() UserTier {
	return u.Tier
}

func (u *User) GetRole() UserRole {
	return u.Role
}

func (u *User) GetIsActive() bool {
	return u.IsActive
}

func (u *User) GetEmailVerified() bool {
	return u.EmailVerified
}

func (u *User) GetShippingAddresses() []Address {
	return u.ShippingAddresses
}
