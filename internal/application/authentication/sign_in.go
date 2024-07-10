package authentication

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/object"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type SignIn interface {
	Execute(ctx context.Context, email, password string) (object.JWT, error)
}

type signIn struct {
	authenticationRepository repository.AuthenticationRepository
}

func NewSignIn(repository repository.AuthenticationRepository) SignIn {
	return signIn{
		authenticationRepository: repository,
	}
}

func (s signIn) Execute(ctx context.Context, email, password string) (object.JWT, error) {
	return s.authenticationRepository.SignIn(ctx, email, password)
}
