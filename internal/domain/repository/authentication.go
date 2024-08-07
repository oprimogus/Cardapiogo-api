package repository

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/object"
)

type AuthenticationRepository interface {
	SignIn(ctx context.Context, email, password string) (object.JWT, error)
	RefreshToken(ctx context.Context, refreshToken string) (object.JWT, error)
	IsValidToken(ctx context.Context, token string) (bool, error)
	DecodeAccessToken(ctx context.Context, accessToken string) (map[string]interface{}, error)
}
