package authentication

import (
	"context"
)

type useCaseSignIn struct {
	repository Repository
}

func newUseCaseSignIn(repository Repository) useCaseSignIn {
	return useCaseSignIn{
		repository: repository,
	}
}

func (s useCaseSignIn) Execute(ctx context.Context, email, password string) (JWT, error) {
	return s.repository.SignIn(ctx, email, password)
}
