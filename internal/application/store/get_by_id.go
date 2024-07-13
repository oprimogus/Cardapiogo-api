package store

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type GetByID struct {
	repository repository.StoreRepository
}

func NewGetByID(repository repository.StoreRepository) GetByID {
	return GetByID{repository: repository}
}

func (g GetByID) Execute(ctx context.Context, id string) (GetStoreByIdOutput, error) {
	store, err := g.repository.FindByID(ctx, id)
	if err != nil {
		return GetStoreByIdOutput{}, err
	}
	return GetStoreByIdOutput{
		ID: store.ID,
		Name: store.Name,
		Phone: store.Phone,
		Score: store.Score,
		Address: AddressOutput{
			AddressLine1: store.Address.AddressLine1,
			AddressLine2: store.Address.AddressLine2,
			Neighborhood: store.Address.Neighborhood,
			City: store.Address.City,
			State: store.Address.State,
			Country: store.Address.Country,
		},
		Type: store.Type,
		BusinessHours: store.BusinessHours,
		PaymentMethodEnums: store.PaymentMethodEnums,

	}, nil
}
