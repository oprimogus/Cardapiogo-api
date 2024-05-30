package infradatabase

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/store"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
)

type StoreRepositoryDatabase struct {
	db *postgres.PostgresDatabase
	q  sqlc.Querier
}

func NewStoreRepositoryDatabase(
	db *postgres.PostgresDatabase,
	querier sqlc.Querier,
) store.Repository {
	return &StoreRepositoryDatabase{db: db, q: querier}
}

func (s *StoreRepositoryDatabase) CreateStore(
	ctx context.Context,
	params store.CreateStoreParams,
) error {
	// tx, err := s.db.GetDB().Begin(ctx)
	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback(ctx)
	//
	// storeID := s.q.CreateStore(ctx, sqlc.CreateStoreParams{
	// 	Name: params.Name,
	// 	CpfCnpj: params.CpfCnpj,
	// 	Phone: params.Phone,
	// 	Type: sqlc.ShopType(params.StoreType),
	// 	Latitude: "0",
	// 	Longitude: "0",
	// 	Street: params.Street,
	// 	Number: ,
	//
	// })

	return errors.New("fail")
}

func (s *StoreRepositoryDatabase) UpdateStore(
	ctx context.Context,
	params store.UpdateStoreParams,
) error {
	return errors.New("fail")
}

func (s *StoreRepositoryDatabase) UpdateStoreCpfCnpj(
	ctx context.Context,
	params store.UpdateStoreCpfCnpjParams,
) error {
	return errors.New("fail")
}

func (s *StoreRepositoryDatabase) DeleteStore(ctx context.Context, ID string) error {
	return errors.New("fail")
}

func (s *StoreRepositoryDatabase) GetStoreByID(
	ctx context.Context,
	ID string,
) (store.StoreDetail, error) {
	return store.StoreDetail{}, nil
}

func (s *StoreRepositoryDatabase) GetStoreByCpfCnpj(
	ctx context.Context,
	CpfCnpj string,
) (store.StoreDetail, error) {
	return store.StoreDetail{}, nil
}

func (s *StoreRepositoryDatabase) GetStoresListByFilter(
	ctx context.Context,
	filter store.StoreFilter,
) ([]store.StoreSummary, error) {
	return []store.StoreSummary{}, nil
}
