package store

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type GetByFilter struct {
	repository repository.StoreRepository
}

func NewGetByFilter(repository repository.StoreRepository) GetByFilter {
	return GetByFilter{repository: repository}
}

func (g GetByFilter) Execute(ctx context.Context, params entity.StoreFilter) (*[]GetStoreByFilterOutput, error) {
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
				OpeningTime: v.OpeningTime.Format(entity.BusinessHourLayout),
				ClosingTime: v.ClosingTime.Format(entity.BusinessHourLayout),
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
