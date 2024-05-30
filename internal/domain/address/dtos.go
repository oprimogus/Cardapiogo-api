package address

type CreateAddressParams struct {
	Street     string `db:"street"     json:"street"`
	Number     string `db:"number"     json:"number"`
	Complement string `db:"complement" json:"complement"`
	District   string `db:"district"   json:"district"`
	ZipCode    string `db:"zip_code"   json:"zip_code"`
	City       string `db:"city"       json:"city"`
	State      string `db:"state"      json:"state"`
}
