package store

import (
	"fmt"
	"time"

	"github.com/oprimogus/cardapiogo/internal/core/address"
)

type CreateParams struct {
	CpfCnpj string          `json:"cpfCnpj" validate:"required,cpfCnpj"`
	Name    string          `json:"name" validate:"required,lte=25"`
	Phone   string          `json:"phone" validate:"required,phone"`
	Address address.Address `json:"address" validate:"required"`
	Type    ShopType        `json:"type" validate:"required,shopType"`
}

type UpdateParams struct {
	ID                 string          `json:"id" validate:"required"`
	Name               string          `json:"name" validate:"required,lte=25"`
	Phone              string          `json:"phone" validate:"required,phone"`
	Address            address.Address `json:"address" validate:"required"`
	Type               ShopType        `json:"type" validate:"required,shopType"`
	PaymentMethodEnums []PaymentMethod `json:"paymentMethod" validate:"dive"`
}

type StoreBusinessHoursParams struct {
	ID            string                `json:"id" validate:"required"`
	TimeZone      string                `json:"timeZone" validate:"required,timezone"`
	BusinessHours []BusinessHoursParams `json:"businessHours" validate:"dive"`
}

type BusinessHoursParams struct {
	WeekDay     int    `json:"weekDay" validate:"min=0,max=6"`
	OpeningTime string `json:"openingTime" validate:"required,businessHour"`
	ClosingTime string `json:"closingTime" validate:"required,businessHour"`
}

func (b *BusinessHoursParams) Entity(timeZone string) (BusinessHours, error) {

	zone, errLoadTimeZone := time.LoadLocation(timeZone)
	if errLoadTimeZone != nil {
		return BusinessHours{}, fmt.Errorf("fail on load timezone: %w", errLoadTimeZone)
	}

	openingTimeParsed, errOpeningTime := time.Parse(BusinessHourLayout, b.OpeningTime)
	if errOpeningTime != nil {
		return BusinessHours{}, fmt.Errorf("fail in parse openingTime: %w", errOpeningTime)
	}

	openingTime := time.Date(1970, time.January, 1, openingTimeParsed.Hour(), openingTimeParsed.Minute(), openingTimeParsed.Second(), 0, time.UTC)

	closingTimeParsed, errClosingTime := time.Parse(BusinessHourLayout, b.ClosingTime)
	if errClosingTime != nil {
		return BusinessHours{}, fmt.Errorf("fail in parse closingTime: %w", errClosingTime)
	}
	closingTime := time.Date(1970, time.January, 1, closingTimeParsed.Hour(), closingTimeParsed.Minute(), closingTimeParsed.Second(), 0, time.UTC)

	return BusinessHours{
		WeekDay:     b.WeekDay,
		OpeningTime: openingTime,
		ClosingTime: closingTime,
		TimeZone:    zone.String(),
	}, nil
}

type GetStoreByIdOutput struct {
	ID                 string                `json:"id"`
	Name               string                `json:"name"`
	Phone              string                `json:"phone"`
	Score              int                   `json:"score"`
	Address            AddressOutput         `json:"address"`
	Type               ShopType              `json:"type"`
	BusinessHours      []BusinessHoursParams `json:"businessHours"`
	PaymentMethodEnums []PaymentMethod       `json:"paymentMethod"`
}

type GetStoreByFilterOutput struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	Score         int                   `json:"score"`
	Type          ShopType              `json:"type"`
	Neighborhood  string                `json:"neighborhood"`
	BusinessHours []BusinessHoursParams `json:"businessHours"`
}

type AddressOutput struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

func (c CreateParams) Entity(userID string) Store {
	return NewStore(userID, c.Name, c.CpfCnpj, c.Phone, c.Address, c.Type)
}
