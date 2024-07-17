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
	BusinessHours      []entity.BusinessHours `json:"businessHours" validate:"dive"`
	PaymentMethodEnums []entity.PaymentMethod `json:"paymentMethod" validate:"dive"`
}

type StoreBusinessHoursParams struct {
	ID            string                 `json:"id" validate:"required"`
	BusinessHours []entity.BusinessHours `json:"businessHours" validate:"dive"`
}

type GetStoreByIdOutput struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Phone              string                 `json:"phone"`
	Score              int                    `json:"score"`
	Address            AddressOutput          `json:"address"`
	Type               entity.ShopType        `json:"type"`
	BusinessHours      []entity.BusinessHours `json:"businessHours"`
	PaymentMethodEnums []entity.PaymentMethod `json:"paymentMethod"`
}

type AddressOutput struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

func (c CreateParams) Entity(userID string) entity.Store {
	return entity.NewStore(userID, c.Name, c.CpfCnpj, c.Phone, c.Address, c.Type)
}
