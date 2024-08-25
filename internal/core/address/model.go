package address

type Address struct {
	AddressLine1 string `db:"address_line_1" json:"addressLine1" validate:"required,lte=40"`
	AddressLine2 string `db:"address_line_2" json:"addressLine2" validate:"required,lte=20"`
	Neighborhood string `db:"neighborhood" json:"neighborhood" validate:"required,lte=25"`
	City         string `db:"city" json:"city" validate:"required,lte=25"`
	State        string `db:"state" json:"state" validate:"required,lte=15"`
	PostalCode   string `db:"postal_code" json:"postalCode" validate:"required,lte=15"`
	Latitude     string `db:"latitude" json:"latitude"`
	Longitude    string `db:"longitude" json:"longitude"`
	Country      string `db:"country" json:"country" validate:"required,lte=15"`
}