package store

import "time"

type Store struct {
	ID         int       `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	CpfCnpj    string    `db:"cpf_cnpj" json:"cpf_cnpj"`
	Phone      string    `db:"phone" json:"phone"`
	Score      uint8     `db:"score" json:"score"`
	Street     string    `db:"street" json:"street"`
	Number     string    `db:"number" json:"number"`
	Complement string    `db:"complement" json:"complement"`
	District   string    `db:"district" json:"district"`
	ZipCode    string    `db:"zip_code" json:"zip_code"`
	City       string    `db:"city" json:"city"`
	State      string    `db:"state" json:"state"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at" json:"deleted_at"`
}
