package repository

import (
	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"golang.org/x/net/context"
)

type AuthenticationRepository interface {
	SignIn(ctx context.Context, email, password string) (entity.JWT, error)
}
