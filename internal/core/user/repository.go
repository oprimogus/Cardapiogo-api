package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user User) error
	FindByID(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	// GetList(ctx context.Context, items int, page int) ([]*User, error)
	Update(ctx context.Context, user User) error
	AddRoles(ctx context.Context, userID string, roles []Role) error
	ResetPasswordByEmail(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}
