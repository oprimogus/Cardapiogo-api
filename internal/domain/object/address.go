package object

type Address struct {
	AddressLine1 string `db:"address_line_1" json:"addressLine1" validate:"required"`
	AddressLine2 string `db:"address_line_2" json:"addressLine2" validate:"required"`
	Neighborhood string `db:"neighborhood" json:"neighborhood" validate:"required"`
	City         string `db:"city" json:"city" validate:"required"`
	State        string `db:"state" json:"state" validate:"required"`
	PostalCode   string `db:"postal_code" json:"postalCode" validate:"required"`
	Latitude     string `db:"latitude" json:"latitude" validate:""`
	Longitude    string `db:"longitude" json:"longitude" validate:""`
	Country      string `db:"country" json:"country" validate:"required"`
}
