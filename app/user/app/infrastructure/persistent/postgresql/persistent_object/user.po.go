package persistentobject

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
	BasePO
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
