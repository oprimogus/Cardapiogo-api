package address

import "time"

type Address struct {
	ID         int       `db:"id"         json:"id"`
	ProfileID  int       `db:"profile_id" json:"profile_id"`
	Street     string    `db:"street"     json:"street"`
	Number     string    `db:"number"     json:"number"`
	Complement string    `db:"complement" json:"complement"`
	District   string    `db:"district"   json:"district"`
	ZipCode    string    `db:"zip_code"   json:"zip_code"`
	City       string    `db:"city"       json:"city"`
	State      string    `db:"state"      json:"state"`
	Latitude   string    `db:"latitude"   json:"latitude"`
	Longitude  string    `db:"longitude"  json:"longitude"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at" json:"deleted_at"`
}
