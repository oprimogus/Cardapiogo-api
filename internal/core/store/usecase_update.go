package store

import (
	"context"
	"fmt"
)

type useCaseUpdate struct {
	repository Repository
}

func newUseCaseUpdate(repository Repository) useCaseUpdate {
	return useCaseUpdate{
		repository: repository,
	}
}

func (c useCaseUpdate) Execute(ctx context.Context, params UpdateParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}

	isOwner, errIsOwner := c.repository.IsOwner(ctx, params.ID, userID)
	if errIsOwner != nil {
		return errIsOwner
	}

	if !isOwner {
		return errNotOwner
	}

	UpdatedStore := Store{
		ID:                 params.ID,
		Name:               params.Name,
		Phone:              params.Phone,
		Address:            params.Address,
		Type:               params.Type,
		PaymentMethodEnums: params.PaymentMethodEnums,
	}

	err := c.repository.Update(ctx, userID, UpdatedStore)
	if err != nil {
		return err
	}
	return nil
}
