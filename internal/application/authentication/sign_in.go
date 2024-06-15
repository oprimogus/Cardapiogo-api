package authentication

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type SignIn interface {
	Execute(ctx context.Context, email, password string) (entity.JWT, error)
}

type signIn struct {
	authenticationRepository repository.AuthenticationRepository
}

func NewSignIn(repository repository.AuthenticationRepository) SignIn {
	return signIn{
		authenticationRepository: repository,
	}
}

func (s signIn) Execute(ctx context.Context, email, password string) (entity.JWT, error) {
	return s.authenticationRepository.SignIn(ctx, email, password)
}
