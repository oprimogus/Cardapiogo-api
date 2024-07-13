package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/object"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/sqlc"
	"github.com/oprimogus/cardapiogo/pkg/converters"
)

type StoreRepository struct {
	db      *postgres.PostgresDatabase
	querier *sqlc.Queries
}

func NewStoreRepository(db *postgres.PostgresDatabase, querier *sqlc.Queries) *StoreRepository {
	return &StoreRepository{db: db, querier: querier}
}

func (s *StoreRepository) Create(ctx context.Context, params entity.Store) error {

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

func (s *StoreRepository) Update(ctx context.Context, userID string, params entity.Store) error {
	convertedUserID, errUserID := converters.ConvertStringToUUID(userID)
	if errUserID != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", errUserID)
	}

	convertedStoreID, errStoreId := converters.ConvertStringToUUID(params.ID)
	if errStoreId != nil {
		return fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}
	err := s.querier.UpdateStore(ctx, sqlc.UpdateStoreParams{
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
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreRepository) FindByID(ctx context.Context, id string) (entity.Store, error) {
	convertedStoreID, errStoreId := converters.ConvertStringToUUID(id)
	if errStoreId != nil {
		return entity.Store{}, fmt.Errorf("fail in convert uuidv7: %w", errStoreId)
	}
	store, err := s.querier.GetStoreByID(ctx, convertedStoreID)
	if err != nil {
		return entity.Store{}, err
	}
	return entity.Store{
		ID:    id,
		Name:  store.Name,
		Phone: store.Phone,
		Score: int(store.Score),
		Type:  entity.ShopType(store.Type),
		Address: object.Address{
			AddressLine1: store.AddressLine1,
			AddressLine2: store.AddressLine2,
			Neighborhood: store.Neighborhood,
			City:         store.City,
			State:        store.State,
			Country:      store.Country,
		},
	}, nil
}

func (s *StoreRepository) GetByFilter(ctx context.Context, params entity.StoreFilter) (*[]entity.Store, error) {
	return &[]entity.Store{}, nil
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
