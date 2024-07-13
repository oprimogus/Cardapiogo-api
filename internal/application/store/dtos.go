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

type GetStoreByIdOutput struct {
	ID                 string          `db:"id" json:"id" validate:"required,uuid"`
	Name               string          `db:"name" json:"name" validate:"required,lte=25"`
	Phone              string          `db:"phone" json:"phone" validate:"required,phone"`
	Score              int             `db:"score" json:"score" validate:"required,number"`
	Address            AddressOutput  `db:"address" json:"address" validate:"required"`
	Type               entity.ShopType        `db:"type" json:"type" validate:"required,shopType"`
	BusinessHours      []entity.BusinessHours `db:"business_hour" json:"businessHours" validate:"required"`
	PaymentMethodEnums []entity.PaymentMethod `db:"payment_method" json:"paymentMethod" validate:"required"`
}

type AddressOutput struct {
	AddressLine1 string `db:"address_line_1" json:"addressLine1" validate:"required,lte=40"`
	AddressLine2 string `db:"address_line_2" json:"addressLine2" validate:"required,lte=20"`
	Neighborhood string `db:"neighborhood" json:"neighborhood" validate:"required,lte=25"`
	City         string `db:"city" json:"city" validate:"required,lte=25"`
	State        string `db:"state" json:"state" validate:"required,lte=15"`
	Country      string `db:"country" json:"country" validate:"required,lte=15"`
}

func (c CreateParams) Entity(userID string) entity.Store {
	return entity.NewStore(userID, c.Name, c.CpfCnpj, c.Phone, c.Address, c.Type)
}
