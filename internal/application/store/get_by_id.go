package store

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type GetByID struct {
	repository repository.StoreRepository
}

func NewGetByID(repository repository.StoreRepository) GetByID {
	return GetByID{repository: repository}
}

func (g GetByID) Execute(ctx context.Context, id string) (GetStoreByIdOutput, error) {
	storeInstance, err := g.repository.FindByID(ctx, id)
	if err != nil {
		return GetStoreByIdOutput{}, err
	}
	businessHours := make([]BusinessHoursParams, len(storeInstance.BusinessHours))
	for i, v := range storeInstance.BusinessHours {
		businessHours[i] = BusinessHoursParams{
			WeekDay:     v.WeekDay,
			OpeningTime: v.OpeningTime.Format(entity.BusinessHourLayout),
			ClosingTime: v.ClosingTime.Format(entity.BusinessHourLayout),
		}
	}
	return GetStoreByIdOutput{
		ID:    storeInstance.ID,
		Name:  storeInstance.Name,
		Phone: storeInstance.Phone,
		Score: storeInstance.Score,
		Address: AddressOutput{
			AddressLine1: storeInstance.Address.AddressLine1,
			AddressLine2: storeInstance.Address.AddressLine2,
			Neighborhood: storeInstance.Address.Neighborhood,
			City:         storeInstance.Address.City,
			State:        storeInstance.Address.State,
			Country:      storeInstance.Address.Country,
		},
		Type:               storeInstance.Type,
		BusinessHours:      businessHours,
		PaymentMethodEnums: storeInstance.PaymentMethodEnums,
	}, nil
}
