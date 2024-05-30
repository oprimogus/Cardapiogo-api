package owner

import "time"

type Owner struct {
	ID        int       `db:"id"         json:"id"`
	ProfileID int       `db:"profile_id" json:"profile_id"`
	StoreID   int       `db:"store_id"   json:"store_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}
