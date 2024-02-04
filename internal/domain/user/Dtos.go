package user

import (
	"github.com/oprimogus/cardapiogo/internal/domain/types"
)

// CreateUserParams DTO
type CreateUserParams struct {
	Email           string                `db:"email" json:"email" validate:"required,email"`
	Password        string                `db:"password" json:"password" validate:"required"`
	Role            types.Role            `db:"role" json:"role" validate:"required,role"`
	AccountProvider types.AccountProvider `db:"account_provider" json:"account_provider"`
}

// CreateUserWithOAuthParams DTO
type CreateUserWithOAuthParams struct {
	Email           string                `db:"email" json:"email" validate:"required,email"`
	Role            types.Role            `db:"role" json:"role" validate:"required,role"`
	AccountProvider types.AccountProvider `db:"account_provider" json:"account_provider" validate:"required,account_provider"`
}

// Login DTO
type Login struct {
	Email    string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required"`
}

// UpdateUserParams DTO
type UpdateUserParams struct {
	ID       string     `db:"id" json:"id" validate:"required,uuid"`
	Email    string     `db:"email" json:"email" validate:"required,email"`
	Password string     `db:"password" json:"password" validate:"required"`
	Role     types.Role `db:"role" json:"role" validate:"required,role"`
}

// UpdateUserPasswordParams DTO
type UpdateUserPasswordParams struct {
	ID          string `db:"id" json:"id" validate:"required"`
	Password    string `db:"password" json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

// UpdateUserProfileParams DTO
type UpdateUserProfileParams struct {
	ID        string `db:"id" json:"id"`
	ProfileID int    `db:"profile_id" json:"profile_id"`
}
