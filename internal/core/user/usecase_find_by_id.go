package user

import (
	"context"
)

type useCaseFindByID struct {
	repository Repository
}

func newUseCaseFindByID(repository Repository) useCaseFindByID {
	return useCaseFindByID{
		repository: repository,
	}
}

func (c useCaseFindByID) Execute(ctx context.Context, email string) (User, error) {
	return c.repository.FindByID(ctx, email)
}
