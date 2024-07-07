package entity

import (
	"time"

	"github.com/oprimogus/cardapiogo/internal/domain/object"
)

const defaultScoreStore = 500

type ShopType string

const (
	StoreShopRestaurant  ShopType = "RESTAURANT"
	StoreShopPharmacy    ShopType = "PHARMACY"
	StoreShopTobbaco     ShopType = "TOBBACO"
	StoreShopMarket      ShopType = "MARKET"
	StoreShopConvenience ShopType = "CONVENIENCE"
	StoreShopPub         ShopType = "PUB"
)

type PaymentMethod string

const (
	Credit PaymentMethod = "CREDIT_CARD"
	Debit  PaymentMethod = "DEBIT_CARD"
	Pix    PaymentMethod = "PIX"
	Cash   PaymentMethod = "CASH"
)

type BusinessHours struct {
	DayOfWeek   time.Weekday
	OpeningTime time.Time
	ClosingTime time.Time
}

type Store struct {
	ID     string
	BusinessID     string
	OwnerID        string
	Name           string
	Active         bool
	Phone          string
	score          int
	Address        object.Address
	StoreType      ShopType
	BusinessHours  []BusinessHours
	PaymentMethods []PaymentMethod
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func NewStore(ownerId, name, businessId, phone string, address object.Address, shopType ShopType) Store {
	store := Store{
		Name:           name,
		BusinessID:     businessId,
		OwnerID:        ownerId,
		Phone:          phone,
		Address:        address,
		StoreType:      shopType,
		score:          defaultScoreStore,
		BusinessHours:  []BusinessHours{},
		PaymentMethods: []PaymentMethod{},
	}
	return store
}