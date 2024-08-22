package entity

import (
	"errors"
	"time"
)

var ErrExistUserWithEmail = errors.New("exist user with this email")
var ErrExistUserWithDocument = errors.New("exist user with this document")
var ErrExistUserWithPhone = errors.New("exist user with this phone")

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
	ID        string     `db:"id" json:"id" validate:"required,uuid"`
	Profile   Profile    `db:"profile" json:"profile" validate:"required"`
	Email     string     `db:"email" json:"email" validate:"required,email"`
	Password  string     `db:"password" json:"password" validate:"required"`
	Roles     []UserRole `db:"role" json:"role" validate:"required,role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Profile struct {
	Name     string `db:"name" json:"name" validate:"required"`
	LastName string `db:"last_name" json:"lastName" validate:"required"`
	Document string `db:"document" json:"document" validate:"cpf"`
	Phone    string `db:"phone" json:"phone" validate:"required,phone"`
}
