package repository

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

type StoreRepository interface {
	Create(ctx context.Context, params entity.Store) error
	Update(ctx context.Context, userID string, params entity.Store) error
	FindByID(ctx context.Context, id string) (entity.Store, error)
	GetByFilter(ctx context.Context, params entity.StoreFilter) (*[]entity.Store, error)
	Delete(ctx context.Context, id string) error
	Enable(ctx context.Context, id string) error
	IsOwner(ctx context.Context, id, userID string) (bool, error)
}
