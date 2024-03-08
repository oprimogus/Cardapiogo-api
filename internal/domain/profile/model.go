package profile

import "time"

type Profile struct {
	ID        int       `db:"id" json:"-"`
	Name      string    `db:"name" json:"name"`
	LastName  string    `db:"last_name" json:"last_name"`
	CPF       string    `db:"cpf" json:"cpf"`
	Phone     string    `db:"phone" json:"phone"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}
