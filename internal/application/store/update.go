package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

var errNotOwner = errors.New("only owner can do this action")

type Update interface {
	Execute(ctx context.Context, params UpdateParams) error
}

type update struct {
	repository repository.StoreRepository
}

func NewUpdate(repository repository.StoreRepository) Update {
	return update{repository: repository}
}

func (u update) Execute(ctx context.Context, params UpdateParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}

	isOwner, errIsOwner := u.repository.IsOwner(ctx, userID)
	if errIsOwner != nil {
		return errIsOwner
	}

	if !isOwner {
		return errNotOwner
	}

	updatedStore := entity.Store{
		Name:               params.Name,
		Phone:              params.Phone,
		Address:            params.Address,
		Type:               params.Type,
		BusinessHours:      params.BusinessHours,
		PaymentMethodEnums: params.PaymentMethodEnums,
	}

	err := u.repository.Update(ctx, updatedStore)
	if err != nil {
		return err
	}
	return nil
}
