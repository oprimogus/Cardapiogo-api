package entity

type ItemType string

const (
	ItemTypeRetail    ItemType = "retail"
	ItemTypeMadeOrder ItemType = "made_to_order"
)

type StoreItem struct {
	ID          string            `json:"id" validate:"required,UUID"`
	StoreID     string            `json:"store_id" validate:"required,UUID"`
	Name        string            `json:"name" validate:"required,gte=2,lte=25"`
	Description string            `json:"description" validate:"required,gte=10,lte=50"`
	Price       int               `json:"price"`
	Cost        int               `json:"cost"`
	Stock       int               `json:"stock"`
	Type        ItemType          `json:"type"`
	Attributes  map[string]string `json:"attributes"`
}
