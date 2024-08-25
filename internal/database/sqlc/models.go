// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentMethodEnum string

const (
	PaymentMethodEnumCredit  PaymentMethodEnum = "credit"
	PaymentMethodEnumDebit   PaymentMethodEnum = "debit"
	PaymentMethodEnumPix     PaymentMethodEnum = "pix"
	PaymentMethodEnumCash    PaymentMethodEnum = "cash"
	PaymentMethodEnumBitcoin PaymentMethodEnum = "bitcoin"
)

func (e *PaymentMethodEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentMethodEnum(s)
	case string:
		*e = PaymentMethodEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentMethodEnum: %T", src)
	}
	return nil
}

type NullPaymentMethodEnum struct {
	PaymentMethodEnum PaymentMethodEnum `json:"payment_method_enum"`
	Valid             bool              `json:"valid"` // Valid is true if PaymentMethodEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentMethodEnum) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentMethodEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentMethodEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentMethodEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PaymentMethodEnum), nil
}

type ShopType string

const (
	ShopTypeRestaurant  ShopType = "restaurant"
	ShopTypePharmacy    ShopType = "pharmacy"
	ShopTypeTobbaco     ShopType = "tobbaco"
	ShopTypeMarket      ShopType = "market"
	ShopTypeConvenience ShopType = "convenience"
	ShopTypePub         ShopType = "pub"
)

func (e *ShopType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ShopType(s)
	case string:
		*e = ShopType(s)
	default:
		return fmt.Errorf("unsupported scan type for ShopType: %T", src)
	}
	return nil
}

type NullShopType struct {
	ShopType ShopType `json:"shop_type"`
	Valid    bool     `json:"valid"` // Valid is true if ShopType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullShopType) Scan(value interface{}) error {
	if value == nil {
		ns.ShopType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ShopType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullShopType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ShopType), nil
}

type Address struct {
	ID           int32            `db:"id" json:"id"`
	UserID       pgtype.UUID      `db:"user_id" json:"user_id"`
	AddressLine1 string           `db:"address_line_1" json:"address_line_1"`
	AddressLine2 string           `db:"address_line_2" json:"address_line_2"`
	Neighborhood string           `db:"neighborhood" json:"neighborhood"`
	City         string           `db:"city" json:"city"`
	State        string           `db:"state" json:"state"`
	PostalCode   string           `db:"postal_code" json:"postal_code"`
	Latitude     pgtype.Text      `db:"latitude" json:"latitude"`
	Longitude    string           `db:"longitude" json:"longitude"`
	Country      string           `db:"country" json:"country"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
}

type BusinessHour struct {
	StoreID     pgtype.UUID `db:"store_id" json:"store_id"`
	WeekDay     int32       `db:"week_day" json:"week_day"`
	OpeningTime pgtype.Time `db:"opening_time" json:"opening_time"`
	ClosingTime pgtype.Time `db:"closing_time" json:"closing_time"`
	Timezone    string      `db:"timezone" json:"timezone"`
}

type PaymentMethod struct {
	ID     int32                 `db:"id" json:"id"`
	Method NullPaymentMethodEnum `db:"method" json:"method"`
}

type Store struct {
	ID      pgtype.UUID `db:"id" json:"id"`
	CpfCnpj string      `db:"cpf_cnpj" json:"cpf_cnpj"`
	// user ID of store owner
	OwnerID      pgtype.UUID      `db:"owner_id" json:"owner_id"`
	Name         string           `db:"name" json:"name"`
	Active       bool             `db:"active" json:"active"`
	Phone        string           `db:"phone" json:"phone"`
	Score        int32            `db:"score" json:"score"`
	Type         ShopType         `db:"type" json:"type"`
	AddressLine1 string           `db:"address_line_1" json:"address_line_1"`
	AddressLine2 string           `db:"address_line_2" json:"address_line_2"`
	Neighborhood string           `db:"neighborhood" json:"neighborhood"`
	City         string           `db:"city" json:"city"`
	State        string           `db:"state" json:"state"`
	PostalCode   string           `db:"postal_code" json:"postal_code"`
	Latitude     pgtype.Text      `db:"latitude" json:"latitude"`
	Longitude    pgtype.Text      `db:"longitude" json:"longitude"`
	Country      string           `db:"country" json:"country"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type StorePaymentMethod struct {
	StoreID         pgtype.UUID `db:"store_id" json:"store_id"`
	PaymentMethodID pgtype.Int4 `db:"payment_method_id" json:"payment_method_id"`
}