package authentication

import (
	"context"
)

type useCaseRefresh struct {
	repository Repository
}

func newUseCaseRefresh(repository Repository) useCaseRefresh {
	return useCaseRefresh{
		repository: repository,
	}
}

func (s useCaseRefresh) Execute(ctx context.Context, refreshToken string) (JWT, error) {
	return s.repository.RefreshToken(ctx, refreshToken)
}
