package user

import (
	"errors"
	"time"
)

var (
	ErrExistUserWithEmail    = errors.New("exist user with this email")
	ErrExistUserWithDocument = errors.New("exist user with this document")
	ErrExistUserWithPhone    = errors.New("exist user with this phone")
)

type Role string

const (
	RoleConsumer    Role = "consumer"
	RoleOwner       Role = "owner"
	RoleEmployee    Role = "employee"
	RoleDeliveryMan Role = "delivery_man"
	RoleAdmin       Role = "admin"
)

func IsValidRole(role string) bool {
	switch Role(role) {
	case RoleConsumer, RoleOwner, RoleEmployee, RoleDeliveryMan, RoleAdmin:
		return true
	default:
		return false
	}
}

type User struct {
	ID        string    `db:"id" json:"id" validate:"required,uuid"`
	Profile   Profile   `db:"profile" json:"profile" validate:"required"`
	Email     string    `db:"email" json:"email" validate:"required,email"`
	Password  string    `db:"password" json:"password" validate:"required"`
	Roles     []Role    `db:"role" json:"role" validate:"required,role"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	DeletedAt time.Time `db:"deleted_at" json:"deletedAt"`
}

type Profile struct {
	Name     string `db:"name" json:"name" validate:"required"`
	LastName string `db:"last_name" json:"lastName" validate:"required"`
	Document string `db:"document" json:"document" validate:"cpf"`
	Phone    string `db:"phone" json:"phone" validate:"required,phone"`
}
