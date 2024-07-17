package store

import (
	"context"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type AddBusinessHour struct {
	repository repository.StoreRepository
}

func NewUpdateBusinessHour(repository repository.StoreRepository) AddBusinessHour {
	return AddBusinessHour{repository: repository}
}

func (u AddBusinessHour) Execute(ctx context.Context, params StoreBusinessHoursParams) error {
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

	err := u.repository.AddBusinessHour(ctx, params.ID, params.BusinessHours)
	if err != nil {
		return err
	}
	return nil
}
