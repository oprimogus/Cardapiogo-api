package user

import (
	"context"
)

// Repository interface
type Repository interface {
	CreateUser(ctx context.Context, user CreateUserParams) error
	// GetUser(ctx context.Context, id string) (User, error)
	// UpdateUser(ctx context.Context, user UpdateUserParams) error
	// UpdateUserPassword(ctx context.Context, user UpdateUserPasswordParams) error
	// UpdateUserProfile(ctx context.Context, user UpdateUserProfileParams) error
}
