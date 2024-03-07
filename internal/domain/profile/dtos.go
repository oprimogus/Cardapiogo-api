package profile

type CreateProfileParams struct {
	Name     string `db:"name" json:"name" validate:"required,alpha,max=25"`
	LastName string `db:"last_name" json:"last_name" validate:"required,alpha,max=40"`
	CPF      string `db:"cpf" json:"cpf" validate:"required,cpf`
	Phone    string `db:"phone" json:"phone" validate:"required,numeric,max=11"`
}

type UpdateProfileParams struct {
	Name     string `db:"name" json:"name" validate:"required,alpha,max=25"`
	LastName string `db:"last_name" json:"last_name" validate:"required,alpha,max=40"`
	Phone    string `db:"phone" json:"phone" validate:"required,numeric,max=11"`
}
