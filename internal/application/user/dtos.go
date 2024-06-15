package user

import "github.com/oprimogus/cardapiogo/internal/domain/entity"

func Roles(rolesJson []string) []entity.UserRole {
	roles := make([]entity.UserRole, len(rolesJson))
	for i, v := range rolesJson {
		roles[i] = entity.UserRole(v)
	}
	return roles
}

type CreateProfileParams struct {
	Name     string `db:"name" json:"name" validate:"required"`
	LastName string `db:"last_name" json:"last_name" validate:"required"`
	Document string `db:"document" json:"document" validate:"required, cpf"`
	Phone    string `db:"phone" json:"phone" validate:"required, phone"`
}

func (d CreateProfileParams) ToEntity() entity.Profile {
	return entity.Profile{
		Name:     d.Name,
		LastName: d.LastName,
		Document: d.Document,
		Phone:    d.Phone,
	}
}

type CreateParams struct {
	Email    string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required"`
	Profile  CreateProfileParams
	Roles    []string `db:"role" json:"role" validate:"required,role"`
}

func (d CreateParams) ToEntity() entity.User {
	roles := Roles(d.Roles)
	return entity.User{
		Profile:  d.Profile.ToEntity(),
		Email:    d.Email,
		Password: d.Password,
		Roles:    roles,
	}
}

type UpdateProfileParams struct {
	Name     string `db:"name" json:"name" validate:"required"`
	LastName string `db:"last_name" json:"last_name" validate:"required"`
	Phone    string `db:"phone" json:"phone" validate:"required, phone"`
}

func (d UpdateProfileParams) ToEntity() entity.User {
	return entity.User{
		Profile: entity.Profile{
			Name:     d.Name,
			LastName: d.LastName,
			Phone:    d.Phone,
		},
	}
}

type Login struct {
	Email    string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required"`
}

type UpdatePasswordParams struct {
	Password    string `db:"password" json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
