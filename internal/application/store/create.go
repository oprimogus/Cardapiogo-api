package store

import (
	"context"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Create interface {
	Execute(ctx context.Context, params CreateParams) error
}

type create struct {
	repository repository.StoreRepository
}

func NewCreate(repository repository.StoreRepository) Create {
	return create{repository: repository}
}

func (c create) Execute(ctx context.Context, params CreateParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}
	err := c.repository.Create(ctx, params.Entity(userID))
	if err != nil {
		return fmt.Errorf("could not create a store for this user: %w", err)
	}
	return nil
}
