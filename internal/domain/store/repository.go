package store

import "context"

type Repository interface {
	CreateStore(ctx context.Context, params CreateStoreParams) error
	UpdateStore(ctx context.Context, params UpdateStoreParams) error
	UpdateStoreCpfCnpj(ctx context.Context, params UpdateStoreCpfCnpjParams) error
	DeleteStore(ctx context.Context, ID string) error
	GetStoreByID(ctx context.Context, ID string) (StoreDetail, error)
	GetStoreByCpfCnpj(ctx context.Context, CpfCnpj string) (StoreDetail, error)
	GetStoresListByFilter(ctx context.Context, filter StoreFilter) ([]StoreSummary, error)
}
