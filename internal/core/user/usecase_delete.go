package user

import (
	"context"
	"errors"
)

type useCaseDelete struct {
	repository Repository
}

func newUseCaseDelete(repository Repository) useCaseDelete {
	return useCaseDelete{repository: repository}
}

func (d useCaseDelete) Execute(ctx context.Context) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	return d.repository.Delete(ctx, userID)
}
