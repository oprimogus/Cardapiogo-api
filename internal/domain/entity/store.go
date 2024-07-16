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
	WeekDay     int    `json:"weekDay" validate:"weekDay"`
	OpeningTime string `json:"openingTime" validate:"required,businessHour"`
	ClosingTime string `json:"closingTime" validate:"required,businessHour"`
}

func IsValidBusinessHourSlice(slice []BusinessHours) bool {
	businessHours := slice
	if len(businessHours) > 0 {
		mapFrequency := map[int]int{}
		for _, v := range businessHours {
			mapFrequency[v.WeekDay] += 1
		}
		for i := range businessHours {
			if mapFrequency[i] > 1 {
				return false
			}
		}
	}
	return true
}

type StoreFilter struct {
	Range int
	Name  string
	Score int
	Type  ShopType
}

type Store struct {
	ID                 string          `json:"id" validate:"required,uuid"`
	CpfCnpj            string          `json:"cpfCnpj" validate:"required,cpfCnpj"`
	OwnerID            string          `json:"owner_id" validate:"required,uuid"`
	Name               string          `json:"name" validate:"required,lte=25"`
	Active             bool            `json:"active" validate:"required,boolean"`
	Phone              string          `json:"phone" validate:"required,phone"`
	Score              int             `json:"score" validate:"required,number"`
	Address            object.Address  `json:"address" validate:"required"`
	Type               ShopType        `json:"type" validate:"required,shopType"`
	BusinessHours      []BusinessHours `json:"businessHours" validate:"dive"`
	PaymentMethodEnums []PaymentMethod `json:"paymentMethod" validate:"dive"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	DeletedAt          time.Time       `json:"deleted_at"`
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
