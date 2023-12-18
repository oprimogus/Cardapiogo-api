package user

type UserRole string

const (
	UserRoleConsumer    UserRole = "Consumer"
	UserRoleOwner       UserRole = "Owner"
	UserRoleEmployee    UserRole = "Employee"
	UserRoleDeliveryMan UserRole = "DeliveryMan"
	UserRoleAdmin       UserRole = "Admin"
)

type AccountProvider string

const (
	AccountProviderGoogle AccountProvider = "Google"
	AccountProviderApple  AccountProvider = "Apple"
	AccountProviderMeta   AccountProvider = "Meta"
)

type User struct {
	ID              string             `db:"id" json:"id"`
	ProfileID       int                `db:"profile_id" json:"profile_id"`
	Email           string             `db:"email" json:"email"`
	Password        string             `db:"password" json:"password"`
	Role            UserRole           `db:"role" json:"role"`
	AccountProvider AccountProvider    `db:"account_provider" json:"account_provider"`
	CreatedAt       pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}
