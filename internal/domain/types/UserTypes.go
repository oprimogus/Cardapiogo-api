package types

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
