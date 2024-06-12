package entity

import (
	"time"
)

type UserRole string

const (
	UserRoleConsumer    UserRole = "CONSUMER"
	UserRoleOwner       UserRole = "OWNER"
	UserRoleEmployee    UserRole = "EMPLOYEE"
	UserRoleDeliveryMan UserRole = "DELIVERY_MAN"
	UserRoleAdmin       UserRole = "ADMIN"
)

type User struct {
	ID         int
	ExternalId string
	Profile    Profile
	Email      string
	Password   string
	Role       UserRole
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
