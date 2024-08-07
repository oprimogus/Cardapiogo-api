package repository

import (
	"github.com/oprimogus/cardapiogo/internal/domain/object"
	"golang.org/x/net/context"
)

type AuthenticationRepository interface {
	SignIn(ctx context.Context, email, password string) (object.JWT, error)
	IsValidToken(ctx context.Context, token string) (bool, error)
	DecodeAccessToken(ctx context.Context, accessToken string) (map[string]interface{}, error)
}
