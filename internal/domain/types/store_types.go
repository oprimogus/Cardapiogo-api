package types

type ShopType string

const (
	StoreShopRestaurant  ShopType = "RESTAURANT"
	StoreShopPharmacy    ShopType = "PHARMACY"
	StoreShopTobbaco     ShopType = "TOBBACO"
	StoreShopMarket      ShopType = "MARKET"
	StoreShopConvenience ShopType = "CONVENIENCE"
	StoreShopPub         ShopType = "PUB"
)

type CousineType string

const (
	CousineItalian   CousineType = "ITALIAN"
	CousineJaponese  CousineType = "JAPANESE"
	CousineMexican   CousineType = "MEXICAN"
	CousineArabic    CousineType = "ARABIC"
	CousineBrazilian CousineType = "BRAZILIAN"
	CousineThai      CousineType = "THAI"
)

type Weekday int

const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

type PaymentForm string

const (
	Credit PaymentForm = "CREDIT_CARD"
	Debit  PaymentForm = "DEBIT_CARD"
	Pix    PaymentForm = "PIX"
	Cash   PaymentForm = "CASH"
)
