package user

import "github.com/oprimogus/cardapiogo/internal/domain/entity"

func Roles(rolesJSON []string) []entity.UserRole {
	roles := make([]entity.UserRole, len(rolesJSON))
	for i, v := range rolesJSON {
		roles[i] = entity.UserRole(v)
	}
	return roles
}

type CreateProfileParams struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Document string `json:"document" validate:"required,cpf"`
	Phone    string `json:"phone" validate:"required"`
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Profile  CreateProfileParams
}

func (d CreateParams) ToEntity() entity.User {
	return entity.User{
		Profile:  d.Profile.ToEntity(),
		Email:    d.Email,
		Password: d.Password,
	}
}

type UpdateProfileParams struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type AddRolesParams struct {
	Roles []entity.UserRole `json:"roles" validate:"required,role"`
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordParams struct {
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
