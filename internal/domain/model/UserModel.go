package model

import "time"

// Role enum
type Role string

const (
	// UserRoleConsumer Enum
	UserRoleConsumer Role = "Consumer"
	// UserRoleOwner Enum
	UserRoleOwner Role = "Owner"
	// UserRoleEmployee Enum
	UserRoleEmployee Role = "Employee"
	// UserRoleDeliveryMan Enum
	UserRoleDeliveryMan Role = "DeliveryMan"
	// UserRoleAdmin Enum
	UserRoleAdmin Role = "Admin"
)

// AccountProvider enum
type AccountProvider string

const (
	// AccountProviderGoogle Enum
	AccountProviderGoogle AccountProvider = "Google"
	// AccountProviderApple Enum
	AccountProviderApple AccountProvider = "Apple"
	// AccountProviderMeta Enum
	AccountProviderMeta AccountProvider = "Meta"
)

// User Model
type User struct {
	ID              string          `db:"id" json:"id"`
	ProfileID       int             `db:"profile_id" json:"profile_id"`
	Email           string          `db:"email" json:"email"`
	Password        string          `db:"password" json:"password"`
	Role            Role            `db:"role" json:"role"`
	AccountProvider AccountProvider `db:"account_provider" json:"account_provider"`
	CreatedAt       time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at" json:"updated_at"`
}
