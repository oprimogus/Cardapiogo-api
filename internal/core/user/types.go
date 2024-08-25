package user

func Roles(rolesJSON []string) []Role {
	roles := make([]Role, len(rolesJSON))
	for i, v := range rolesJSON {
		roles[i] = Role(v)
	}
	return roles
}

type CreateProfileParams struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Phone    string `json:"phone" validate:"required,phone"`
}

func (d CreateProfileParams) ToEntity() Profile {
	return Profile{
		Name:     d.Name,
		LastName: d.LastName,
		Phone:    d.Phone,
	}
}

type CreateParams struct {
	Email    string              `json:"email" validate:"required,email"`
	Password string              `json:"password" validate:"required"`
	Profile  CreateProfileParams `json:"profile" validate:"required"`
}

func (d CreateParams) ToEntity() User {
	return User{
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
	Roles []Role `json:"role" validate:"required,role"`
}

func (d UpdateProfileParams) ToEntity() User {
	return User{
		Profile: Profile{
			Name:     d.Name,
			LastName: d.LastName,
			Phone:    d.Phone,
		},
	}
}

type LoginParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordParams struct {
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
