package store

import (
	"context"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type UpdateBusinessHour struct {
	repository repository.StoreRepository
}

func NewUpdateBusinessHour(repository repository.StoreRepository) UpdateBusinessHour {
	return UpdateBusinessHour{repository: repository}
}

func (u UpdateBusinessHour) Execute(ctx context.Context, params StoreBusinessHoursParams) error {
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

	err := u.repository.UpsertBusinessHour(ctx, params.ID, params.BusinessHours)
	if err != nil {
		return err
	}
	return nil
}
