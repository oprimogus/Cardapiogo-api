package profile

import "time"

type Profile struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	LastName  string    `db:"last_name" json:"last_name"`
	CPF       string    `db:"cpf" json:"cpf"`
	Phone     string    `db:"phone" json:"phone"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
