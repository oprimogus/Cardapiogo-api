package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/oprimogus/cardapiogo/internal/core/address"
	"github.com/oprimogus/cardapiogo/internal/core/store"
	"github.com/oprimogus/cardapiogo/internal/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/database/sqlc"
	"github.com/oprimogus/cardapiogo/pkg/converters"
)

type StoreRepository struct {
	db      *postgres.PostgresDatabase
	querier *sqlc.Queries
}

func NewStoreRepository(db *postgres.PostgresDatabase, querier *sqlc.Queries) *StoreRepository {
	return &StoreRepository{db: db, querier: querier}
}

func (s *StoreRepository) Create(ctx context.Context, params store.Store) error {

	ownerIDUUID, err := converters.ConvertStringToUUID(params.OwnerID)
	if err != nil {
		return fmt.Errorf("fail in uuid convert: %w", err)
	}

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("fail in generate uuidv7: %w", err)
	}
	convertedUUIDV7, err := converters.ConvertStringToUUID(id.String())
	if err != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", err)
	}

	args := sqlc.CreateStoreParams{
		ID:           convertedUUIDV7,
		CpfCnpj:      params.CpfCnpj,
		OwnerID:      ownerIDUUID,
		Name:         params.Name,
		Active:       params.Active,
		Phone:        params.Phone,
		Score:        int32(params.Score),
		Type:         sqlc.ShopType(params.Type),
		AddressLine1: params.Address.AddressLine1,
		AddressLine2: params.Address.AddressLine2,
		Neighborhood: params.Address.Neighborhood,
		City:         params.Address.City,
		State:        params.Address.State,
		PostalCode:   params.Address.PostalCode,
		Latitude:     converters.ConvertStringToText(params.Address.Latitude),
		Longitude:    converters.ConvertStringToText(params.Address.Longitude),
		Country:      params.Address.Country,
	}
	err = s.querier.CreateStore(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreRepository) Update(ctx context.Context, userID string, params store.Store) error {
	convertedUserID, errUserID := converters.ConvertStringToUUID(userID)
	if errUserID != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", errUserID)
	}

	convertedStoreID, errStoreId := converters.ConvertStringToUUID(params.ID)
	if errStoreId != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}

	errUpdateStore := s.querier.UpdateStore(ctx, sqlc.UpdateStoreParams{
		ID:           convertedStoreID,
		OwnerID:      convertedUserID,
		Name:         params.Name,
		Phone:        params.Phone,
		Type:         sqlc.ShopType(params.Type),
		AddressLine1: params.Address.AddressLine1,
		AddressLine2: params.Address.AddressLine2,
		Neighborhood: params.Address.Neighborhood,
		City:         params.Address.City,
		State:        params.Address.State,
		PostalCode:   params.Address.PostalCode,
		Country:      params.Address.Country,
	})
	if errUpdateStore != nil {
		return errUpdateStore
	}
	return nil
}

func (s *StoreRepository) AddBusinessHour(ctx context.Context, storeID string, params []store.BusinessHours) error {
	convertedStoreID, errStoreId := converters.ConvertStringToUUID(storeID)
	if errStoreId != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}

	argsSlice := make([]sqlc.AddBusinessHoursParams, len(params))
	for i, v := range params {
		argsSlice[i] = sqlc.AddBusinessHoursParams{
			StoreID:     convertedStoreID,
			WeekDay:     int32(v.WeekDay),
			OpeningTime: converters.PgtypeTime(v.OpeningTime),
			ClosingTime: converters.PgtypeTime(v.ClosingTime),
			Timezone:    v.TimeZone,
		}
	}

	batchAddBusinessHour := s.querier.AddBusinessHours(ctx, argsSlice)
	batchAddBusinessHour.Exec(nil)
	errBatchAddBusinessHour := batchAddBusinessHour.Close()
	if errBatchAddBusinessHour != nil {
		return errBatchAddBusinessHour
	}
	return nil
}

func (s *StoreRepository) DeleteBusinessHour(ctx context.Context, storeID string, params []store.BusinessHours) error {
	convertedStoreID, errStoreId := converters.ConvertStringToUUID(storeID)
	if errStoreId != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}

	argsSlice := make([]sqlc.DeleteBusinessHoursParams, len(params))
	for i, v := range params {
		argsSlice[i] = sqlc.DeleteBusinessHoursParams{
			StoreID:     convertedStoreID,
			WeekDay:     int32(v.WeekDay),
			OpeningTime: converters.PgtypeTime(v.OpeningTime),
			ClosingTime: converters.PgtypeTime(v.ClosingTime),
		}
	}
	batchDeleteBusinessHours := s.querier.DeleteBusinessHours(ctx, argsSlice)
	batchDeleteBusinessHours.Exec(nil)
	return nil
}

func (s *StoreRepository) FindByID(ctx context.Context, id string) (store.Store, error) {
	convertedStoreID, errStoreId := converters.ConvertStringToUUID(id)
	if errStoreId != nil {
		return store.Store{}, fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}
	storeInstance, err := s.querier.GetStoreByID(ctx, convertedStoreID)
	if err != nil {
		return store.Store{}, err
	}

	sqlcStoreBusinessHours, errSqlc := s.querier.GetStoreBusinessHoursByID(ctx, convertedStoreID)
	if errSqlc != nil {
		return store.Store{}, errSqlc
	}

	businessHours := make([]store.BusinessHours, len(sqlcStoreBusinessHours))
	if len(sqlcStoreBusinessHours) > 0 {
		for i, v := range sqlcStoreBusinessHours {
			openingTime, errOpeningTime := converters.Time(v.OpeningTime)
			if errOpeningTime != nil {
				return store.Store{}, errOpeningTime
			}
			closingTime, errClosingTime := converters.Time(v.ClosingTime)
			if errClosingTime != nil {
				return store.Store{}, errClosingTime
			}

			businessHours[i] = store.BusinessHours{
				WeekDay:     int(v.WeekDay),
				OpeningTime: openingTime,
				ClosingTime: closingTime,
			}
		}
	}

	return store.Store{
		ID:    id,
		Name:  storeInstance.Name,
		Phone: storeInstance.Phone,
		Score: int(storeInstance.Score),
		Type:  store.ShopType(storeInstance.Type),
		Address: address.Address{
			AddressLine1: storeInstance.AddressLine1,
			AddressLine2: storeInstance.AddressLine2,
			Neighborhood: storeInstance.Neighborhood,
			City:         storeInstance.City,
			State:        storeInstance.State,
			Country:      storeInstance.Country,
		},
		BusinessHours: businessHours,
	}, nil
}

func (s *StoreRepository) FindByFilter(ctx context.Context, params store.StoreFilter) (*[]store.Store, error) {

	args := sqlc.GetStoreByFilterParams{
		Column1: converters.ConvertStringToText(params.Name),
		Column2: int32(params.Score),
		Column3: params.Type,
		Column4: params.City,
	}
	filteredStores, err := s.querier.GetStoreByFilter(ctx, args)
	if err != nil {
		return nil, err
	}

	uuids := make([]pgtype.UUID, len(filteredStores))
	for i, v := range filteredStores {
		uuids[i] = v.ID
	}

	businessHours, errFindBusinessHour := s.querier.FindStoreBusinessHoursByStoreId(ctx, uuids)
	if errFindBusinessHour != nil {
		return nil, errFindBusinessHour
	}

	mapBusinessHours := make(map[pgtype.UUID][]sqlc.BusinessHour)
	for _, v := range businessHours {
		mapBusinessHours[v.StoreID] = append(mapBusinessHours[v.StoreID], v)
	}

	stores := make([]store.Store, len(filteredStores))
	for i, v := range filteredStores {
		convertedID, errConvertUUID := converters.ConvertUUIDToString(v.ID)
		if errConvertUUID != nil {
			return nil, fmt.Errorf("fail on convert database UUID: %w", errConvertUUID)
		}
		businessHours := mapBusinessHours[v.ID]
		entityBusinessHour := make([]store.BusinessHours, len(businessHours))
		for i, v := range businessHours {
			openingTime, errOpeningTime := converters.Time(v.OpeningTime)
			if errOpeningTime != nil {
				return nil, fmt.Errorf("fail on convert openingTime: %w", errOpeningTime)
			}
			closingTime, errClosingTime := converters.Time(v.ClosingTime)
			if errClosingTime != nil {
				return nil, fmt.Errorf("fail on convert closingTime: %w", errClosingTime)
			}
			entityBusinessHour[i] = store.BusinessHours{
				WeekDay:     int(v.WeekDay),
				OpeningTime: openingTime,
				ClosingTime: closingTime,
			}

		}
		stores[i] = store.Store{
			ID:    *convertedID,
			Name:  v.Name,
			Score: int(v.Score),
			Type:  store.ShopType(v.Type),
			Address: address.Address{
				Neighborhood: v.Neighborhood,
				Latitude:     v.Latitude.String,
				Longitude:    v.Longitude.String,
			},
			BusinessHours: entityBusinessHour,
		}
	}

	return &stores, nil
}

func (s *StoreRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *StoreRepository) Enable(ctx context.Context, id string) error {
	return nil
}

func (s *StoreRepository) IsOwner(ctx context.Context, id, userID string) (bool, error) {
	convertedUserID, errUserID := converters.ConvertStringToUUID(userID)
	if errUserID != nil {
		return false, fmt.Errorf("fail in convert uuidv7: %w", errUserID)
	}

	convertedStoreID, errStoreId := converters.ConvertStringToUUID(id)
	if errStoreId != nil {
		return false, fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}
	isOwner, err := s.querier.IsOwner(ctx, sqlc.IsOwnerParams{ID: convertedStoreID, OwnerID: convertedUserID})
	if err != nil {
		return false, err
	}
	return isOwner, nil
}
