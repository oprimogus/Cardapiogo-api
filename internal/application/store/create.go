package store

import (
	"context"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Create struct {
	repository repository.StoreRepository
}

func NewCreate(repository repository.StoreRepository) Create {
	return Create{repository: repository}
}

func (c Create) Execute(ctx context.Context, params CreateParams) error {
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
