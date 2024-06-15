package repository

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) error
	FindByID(ctx context.Context, id string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	// GetList(ctx context.Context, items int, page int) ([]*User, error)
	Update(ctx context.Context, user entity.User) error
	ResetPasswordByEmail(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}
