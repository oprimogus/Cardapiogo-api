package user

import (
	"context"
)

// Repository interface
type Repository interface {
	CreateUser(ctx context.Context, params CreateUserParams) error
	CreateUserWithOAuth(ctx context.Context, params CreateUserWithOAuthParams) error
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsersList(ctx context.Context, items int, page int) ([]*User, error)
	UpdateUser(ctx context.Context, params UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, params UpdateUserPasswordParams) error
	UpdateUserProfile(ctx context.Context, params UpdateUserProfileParams) error
}
