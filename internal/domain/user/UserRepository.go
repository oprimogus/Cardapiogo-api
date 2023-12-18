package user

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/database/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user sqlc.CreateUserParams) error
	GetUser(ctx context.Context, id pgtype.UUID) (Users, error)
	UpdateUser(ctx context.Context, user UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, user UpdateUserPasswordParams) error
	UpdateUserProfile(ctx context.Context, user UpdateUserProfileParams) error
}
