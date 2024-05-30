package store

import (
	"time"

	"github.com/oprimogus/cardapiogo/internal/domain/address"
	"github.com/oprimogus/cardapiogo/internal/domain/types"
)

type CreateStoreParams struct {
	Name      string                      `db:"name"       json:"name"       validate:"required,alpha,max=25"`
	CpfCnpj   string                      `db:"cpf_cnpj"   json:"cpf_cnpj"   validate:"required,cpf"`
	Phone     string                      `db:"phone"      json:"phone"      validate:"required,numeric,max=11"`
	Address   address.CreateAddressParams `db:"address"    json:"address"    validate:"required"`
	StoreType types.ShopType              `db:"store_type" json:"store_type" validate:"required,shopType"`
}

type UpdateStoreParams struct {
	ExternalID string                      `db:"external_id" json:"external_id"`
	Name       string                      `db:"name"        json:"name"`
	CpfCnpj    string                      `db:"cpf_cnpj"    json:"cpf_cnpj"`
	Phone      string                      `db:"phone"       json:"phone"`
	Address    address.CreateAddressParams `db:"address"     json:"address"`
	StoreType  types.ShopType              `db:"store_type"  json:"store_type"`
}

type UpdateStoreCpfCnpjParams struct {
	ExternalID string `db:"external_id" json:"external_id"`
	CpfCnpj    string `db:"cpf_cnpj"    json:"cpf_cnpj"`
}

type StoreSummary struct {
	ExternalID string `db:"external_id" json:"external_id"`
	Name       string `db:"name"        json:"name"`
	Score      uint8  `db:"score"       json:"score"`
	District   string `db:"district"    json:"district"`
}

type StoreDetail struct {
	Name      string                      `db:"name"       json:"name"`
	Phone     string                      `db:"phone"      json:"phone"`
	StoreType types.ShopType              `db:"store_type" json:"store_type"`
	Score     uint8                       `db:"score"      json:"score"`
	Address   address.CreateAddressParams `db:"address"    json:"address"`
	CreatedAt time.Time                   `db:"created_at" json:"created_at"`
}

type StoreFilter struct {
	Page         int
	Items        int
	NameLike     string
	StoreType    types.ShopType
	CousinesType []types.CousineType
}
