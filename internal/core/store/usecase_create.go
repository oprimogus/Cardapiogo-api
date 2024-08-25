package store

import (
	"context"
	"fmt"
)

type useCaseCreate struct {
	repository Repository
}

func newUseCaseCreate(repository Repository) useCaseCreate {
	return useCaseCreate{
		repository: repository,
	}
}

func (c useCaseCreate) Execute(ctx context.Context, params CreateParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}
	err := c.repository.Create(ctx, params.Entity(userID))
	if err != nil {
		return fmt.Errorf("could not Create a store for this user: %w", err)
	}
	return nil
}