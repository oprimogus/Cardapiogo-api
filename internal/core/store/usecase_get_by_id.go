package store

import (
	"context"
)

type useCaseGetByID struct {
	repository Repository
}

func newUseCaseGetByID(repository Repository) useCaseGetByID {
	return useCaseGetByID{
		repository: repository,
	}
}

func (g useCaseGetByID) Execute(ctx context.Context, id string) (GetStoreByIdOutput, error) {
	storeInstance, err := g.repository.FindByID(ctx, id)
	if err != nil {
		return GetStoreByIdOutput{}, err
	}
	businessHours := make([]BusinessHoursParams, len(storeInstance.BusinessHours))
	for i, v := range storeInstance.BusinessHours {
		businessHours[i] = BusinessHoursParams{
			WeekDay:     v.WeekDay,
			OpeningTime: v.OpeningTime.Format(BusinessHourLayout),
			ClosingTime: v.ClosingTime.Format(BusinessHourLayout),
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
		ProfileImage: storeInstance.ProfileImage,
		HeaderImage: storeInstance.HeaderImage,
	}, nil
}
