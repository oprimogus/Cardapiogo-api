package entity

import (
	"regexp"
	"time"

	"github.com/oprimogus/cardapiogo/internal/domain/object"
)

const defaultScoreStore = 500

type ShopType string

const (
	StoreShopRestaurant  ShopType = "restaurant"
	StoreShopPharmacy    ShopType = "pharmacy"
	StoreShopTobbaco     ShopType = "tobbaco"
	StoreShopMarket      ShopType = "market"
	StoreShopConvenience ShopType = "convenience"
	StoreShopPub         ShopType = "pub"
)

func IsValidShopType(shopType string) bool {
	switch ShopType(shopType) {
	case StoreShopRestaurant,
		StoreShopPharmacy,
		StoreShopTobbaco,
		StoreShopMarket,
		StoreShopConvenience,
		StoreShopPub:
		return true
	default:
		return false
	}
}

type PaymentMethod string

const (
	Credit PaymentMethod = "credit"
	Debit  PaymentMethod = "debit"
	Pix    PaymentMethod = "pix"
	Cash   PaymentMethod = "cash"
)

func IsValidPaymentMethod(PaymentMethodEnum string) bool {
	switch PaymentMethod(PaymentMethodEnum) {
	case Credit, Debit, Pix, Cash:
		return true
	default:
		return false
	}
}

func IsBusinessHour(value string) bool {
	re := regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d$`)
	return re.MatchString(value)
}

type BusinessHours struct {
	WeekDay     time.Weekday `db:"week_day" json:"weekDay" validate:"required,number,gte=0,lte=6"`
	OpeningTime string       `db:"opening_time" json:"openingTime" validate:"required,businessHour"`
	ClosingTime string       `db:"closing_time" json:"closingTime" validate:"required,businessHour"`
}

type StoreFilter struct {
	Range int
	Name  string
	Score int
	Type  ShopType
}

type Store struct {
	ID                 string          `db:"id" json:"id" validate:"required,uuid"`
	CpfCnpj            string          `db:"cpf_cnpj" json:"cpfCnpj" validate:"required"`
	OwnerID            string          `db:"owner_id" json:"owner_id" validate:"required,uuid"`
	Name               string          `db:"name" json:"name" validate:"required"`
	Active             bool            `db:"active" json:"active" validate:"required,boolean"`
	Phone              string          `db:"phone" json:"phone" validate:"required,phone"`
	Score              int             `db:"score" json:"score" validate:"required,number"`
	Address            object.Address  `db:"address" json:"address" validate:"required"`
	Type               ShopType        `db:"type" json:"type" validate:"required,shopType"`
	BusinessHours      []BusinessHours `db:"business_hour" json:"businessHours" validate:"required"`
	PaymentMethodEnums []PaymentMethod `db:"payment_method" json:"paymentMethod" validate:"required"`
	CreatedAt          time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time       `db:"updated_at" json:"updated_at"`
	DeletedAt          time.Time       `db:"deleted_at" json:"deleted_at"`
}

func NewStore(ownerID, name, cpfCnpj, phone string, address object.Address, shopType ShopType) Store {
	store := Store{
		Name:               name,
		CpfCnpj:            cpfCnpj,
		OwnerID:            ownerID,
		Phone:              phone,
		Address:            address,
		Type:               shopType,
		Score:              defaultScoreStore,
		BusinessHours:      []BusinessHours{},
		PaymentMethodEnums: []PaymentMethod{},
	}
	return store
}
