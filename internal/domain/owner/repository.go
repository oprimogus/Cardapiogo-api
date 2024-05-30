package owner

import "context"

type Repository interface {
	AddOwner(ctx context.Context, userID string, storeID string) error
	DeleteOwner(ctx context.Context, userID string, storeID string) error
	ListOwnersOfStore(ctx context.Context, storeID string) ([]Owner, error)
	IsOwner(ctx context.Context, userID string, storeID string) (bool, error)
}
