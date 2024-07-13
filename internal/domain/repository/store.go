package repository

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

type StoreRepository interface {
	Create(ctx context.Context, params entity.Store) error
	Update(ctx context.Context, params entity.Store) error
	GetByID(ctx context.Context, id string) (entity.Store, error)
	GetByFilter(ctx context.Context, params entity.StoreFilter) (*[]entity.Store, error)
	Delete(ctx context.Context, id string) error
	Enable(ctx context.Context, id string) error
	IsOwner(ctx context.Context, userID string) (bool, error)
}
