package authentication

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/object"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Refresh struct {
	authenticationRepository repository.AuthenticationRepository
}

func NewRefresh(repository repository.AuthenticationRepository) Refresh {
	return Refresh{
		authenticationRepository: repository,
	}
}

func (r Refresh) Execute(ctx context.Context, refreshToken string) (object.JWT, error) {
	return r.authenticationRepository.RefreshToken(ctx, refreshToken)
}
