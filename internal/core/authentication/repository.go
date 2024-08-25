package authentication

import "context"

type Repository interface {
	SignIn(ctx context.Context, email, password string) (JWT, error)
	RefreshToken(ctx context.Context, refreshToken string) (JWT, error)
	IsValidToken(ctx context.Context, token string) (bool, error)
	DecodeAccessToken(ctx context.Context, accessToken string) (map[string]interface{}, error)
}
