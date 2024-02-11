package user

import (
	"context"
)

// Repository interface
type Repository interface {
	CreateUser(ctx context.Context, user CreateUserParams) error
	CreateUserWithOAuth(ctx context.Context, user CreateUserWithOAuthParams) error
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsersList(ctx context.Context, items int, page int) ([]*User, error)
	UpdateUser(ctx context.Context, user UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, user UpdateUserPasswordParams) error
	// UpdateUserProfile(ctx context.Context, user UpdateUserProfileParams) error
}
