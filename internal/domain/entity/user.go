package entity

import (
	"time"
)

type UserRole string

const (
	UserRoleConsumer    UserRole = "consumer"
	UserRoleOwner       UserRole = "owner"
	UserRoleEmployee    UserRole = "employee"
	UserRoleDeliveryMan UserRole = "delivery_man"
	UserRoleAdmin       UserRole = "admin"
)

func IsValidUserRole(role string) bool {
	switch UserRole(role) {
	case UserRoleConsumer, UserRoleOwner, UserRoleEmployee, UserRoleDeliveryMan, UserRoleAdmin:
		return true
	default:
		return false
	}
}

type User struct {
	ID        string
	Profile   Profile
	Email     string
	Password  string
	Roles     []UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Profile struct {
	Name     string
	LastName string
	Document string
	Phone    string
}
