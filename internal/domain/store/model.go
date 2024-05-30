package store

import (
	"time"
)

type Store struct {
	ID         int       `db:"id"          json:"id"`
	ExternalID string    `db:"external_id" json:"external_id"`
	Name       string    `db:"name"        json:"name"`
	Active     bool      `db:"active"      json:"active"`
	CpfCnpj    string    `db:"cpf_cnpj"    json:"cpf_cnpj"`
	Phone      string    `db:"phone"       json:"phone"`
	Score      uint8     `db:"score"       json:"score"`
	AddressID  int       `db:"address_id"  json:"address_id"`
	CreatedAt  time.Time `db:"created_at"  json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"  json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at"  json:"deleted_at"`
}
