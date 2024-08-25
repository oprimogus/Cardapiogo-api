package user

import (
	"context"
	"errors"
)

type useCaseUpdate struct {
	repository Repository
}

func newUseCaseUpdate(repository Repository) useCaseUpdate {
	return useCaseUpdate{
		repository: repository,
	}
}

func (c useCaseUpdate) Execute(ctx context.Context, input UpdateProfileParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	user := input.ToEntity()
	user.ID = userID
	return c.repository.Update(ctx, user)
}
