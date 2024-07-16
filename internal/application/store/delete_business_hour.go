package store

import (
	"context"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type DeleteBusinessHour struct {
	repository repository.StoreRepository
}

func NewDeleteBusinessHour(repository repository.StoreRepository) DeleteBusinessHour {
	return DeleteBusinessHour{repository: repository}
}

func (u DeleteBusinessHour) Execute(ctx context.Context, params StoreBusinessHoursParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}

	isOwner, errIsOwner := u.repository.IsOwner(ctx, params.ID, userID)
	if errIsOwner != nil {
		return fmt.Errorf("fail on get store data")
	}

	if !isOwner {
		return errNotOwner
	}

	err := u.repository.DeleteBusinessHour(ctx, params.ID, params.BusinessHours)
	if err != nil {
		return err
	}
	return nil
}