package authentication

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/object"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type SignIn struct {
	authenticationRepository repository.AuthenticationRepository
}

func NewSignIn(repository repository.AuthenticationRepository) SignIn {
	return SignIn{
		authenticationRepository: repository,
	}
}

func (s SignIn) Execute(ctx context.Context, email, password string) (object.JWT, error) {
	return s.authenticationRepository.SignIn(ctx, email, password)
}
