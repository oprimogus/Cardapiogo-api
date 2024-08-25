package store

import (
	"context"
)

type useCaseGetByFilter struct {
	repository Repository
}

func newUseCaseGetByFilter(repository Repository) useCaseGetByFilter {
	return useCaseGetByFilter{
		repository: repository,
	}
}

func (g useCaseGetByFilter) Execute(ctx context.Context, params StoreFilter) (*[]GetStoreByFilterOutput, error) {
	stores, err := g.repository.FindByFilter(ctx, params)
	if err != nil {
		return nil, err
	}

	storesFiltered := make([]GetStoreByFilterOutput, len(*stores))
	for i, v := range *stores {
		businesHours := make([]BusinessHoursParams, len(v.BusinessHours))
		for i, v := range v.BusinessHours {
			businesHours[i] = BusinessHoursParams{
				WeekDay:     v.WeekDay,
				OpeningTime: v.OpeningTime.Format(BusinessHourLayout),
				ClosingTime: v.ClosingTime.Format(BusinessHourLayout),
			}
		}
		storesFiltered[i] = GetStoreByFilterOutput{
			ID:            v.ID,
			Name:          v.Name,
			Score:         v.Score,
			Type:          v.Type,
			Neighborhood:  v.Address.Neighborhood,
			BusinessHours: businesHours,
		}
	}
	return &storesFiltered, nil
}
