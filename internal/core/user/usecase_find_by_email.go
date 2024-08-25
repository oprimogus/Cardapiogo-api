package user

import (
	"context"
)

type useCaseFindByEmail struct {
	repository Repository
}

func newUseCaseFindByEmail(repository Repository) useCaseFindByEmail {
	return useCaseFindByEmail{
		repository: repository,
	}
}

func (c useCaseFindByEmail) Execute(ctx context.Context, email string) (User, error) {
	return c.repository.FindByEmail(ctx, email)
}
