package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
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
		State:        params.Address.PostalCode,
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

func (s *StoreRepository) Update(ctx context.Context, params entity.Store) error {
	return nil
}

func (s *StoreRepository) GetByID(ctx context.Context, id string) (entity.Store, error) {
	return entity.Store{}, nil
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

func (s *StoreRepository) IsOwner(ctx context.Context, userID string) (bool, error) {
	return false, nil
}
