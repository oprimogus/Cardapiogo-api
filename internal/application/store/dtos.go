package store

import (
	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/object"
)

type CreateParams struct {
	CpfCnpj string          `json:"cpfCnpj" validate:"required,cpfCnpj"`
	Name    string          `json:"name" validate:"required,lte=25"`
	Phone   string          `json:"phone" validate:"required,phone"`
	Address object.Address  `json:"address" validate:"required"`
	Type    entity.ShopType `json:"type" validate:"required,shopType"`
}

type UpdateParams struct {
	ID                 string                 `json:"id" validate:"required"`
	Name               string                 `json:"name" validate:"required,lte=25"`
	Phone              string                 `json:"phone" validate:"required,phone"`
	Address            object.Address         `json:"address" validate:"required"`
	Type               entity.ShopType        `json:"type" validate:"required,shopType"`
	BusinessHours      []entity.BusinessHours `json:"businessHours"`
	PaymentMethodEnums []entity.PaymentMethod
}

func (c CreateParams) Entity(userID string) entity.Store {
	return entity.NewStore(userID, c.Name, c.CpfCnpj, c.Phone, c.Address, c.Type)
}
