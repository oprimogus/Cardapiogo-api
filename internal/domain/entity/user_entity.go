package entity

import (
	"time"
)

type UserRole string

const (
	UserRoleConsumer    UserRole = "Consumer"
	UserRoleOwner       UserRole = "Owner"
	UserRoleEmployee    UserRole = "Employee"
	UserRoleDeliveryMan UserRole = "Delivery_Man"
	UserRoleAdmin       UserRole = "Admin"
)

type User struct {
	ID         int
	ExternalId string
	Profile    Profile
	Email      string
	Password   string
	Roles      []UserRole
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type Profile struct {
	Name     string
	LastName string
	Document string
	Phone    string
}
