// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddOwner(ctx context.Context, arg AddOwnerParams) error
	CreateProfileAndReturnID(ctx context.Context, arg CreateProfileAndReturnIDParams) (int32, error)
	CreateStore(ctx context.Context, arg CreateStoreParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) error
	CreateUserWithOAuth(ctx context.Context, arg CreateUserWithOAuthParams) error
	GetProfileByID(ctx context.Context, id int32) (GetProfileByIDRow, error)
	GetProfileByUserID(ctx context.Context, id pgtype.UUID) (GetProfileByUserIDRow, error)
	GetStoresListByFilter(ctx context.Context, arg GetStoresListByFilterParams) ([]GetStoresListByFilterRow, error)
	GetUser(ctx context.Context, arg GetUserParams) ([]GetUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
	GetUserById(ctx context.Context, id pgtype.UUID) (GetUserByIdRow, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) error
	UpdateProfileCpf(ctx context.Context, arg UpdateProfileCpfParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error
}

var _ Querier = (*Queries)(nil)
