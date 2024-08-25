package store

import "context"

type Repository interface {
	Create(ctx context.Context, params Store) (id string, err error)
	Update(ctx context.Context, userID string, params Store) error
	FindByID(ctx context.Context, id string) (Store, error)
	FindByFilter(ctx context.Context, params StoreFilter) (*[]Store, error)
	Delete(ctx context.Context, id string) error
	Enable(ctx context.Context, id string) error
	IsOwner(ctx context.Context, id, userID string) (bool, error)
	AddBusinessHour(ctx context.Context, storeID string, params []BusinessHours) error
	DeleteBusinessHour(ctx context.Context, storeID string, params []BusinessHours) error
}
