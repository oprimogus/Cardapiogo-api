package repository

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/model"
)

// UserRepository interface
type UserRepository interface {
	CreateUser(ctx context.Context, user model.CreateUserParams) error
	// GetUser(ctx context.Context, id string) (User, error)
	// UpdateUser(ctx context.Context, user UpdateUserParams) error
	// UpdateUserPassword(ctx context.Context, user UpdateUserPasswordParams) error
	// UpdateUserProfile(ctx context.Context, user UpdateUserProfileParams) error
}
