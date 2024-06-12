package entity

import "time"

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
	ID             int
	ExternalID     string
	BusinessID     string
	OwnerID        string
	Name           string
	active         bool
	Phone          string
	score          int
	Address        Address
	StoreType      ShopType
	BusinessHours  []BusinessHours
	PaymentMethods []PaymentMethod
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func NewStore(ownerId, name, businessId, phone string, address Address, shopType ShopType) Store {
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
	store.hasPlanActive()
	return store
}

// hasPlanActive Should verify if store owner has a plan active
// TODO: Introduce logic here
func (s *Store) hasPlanActive() {
	s.active = false
}
