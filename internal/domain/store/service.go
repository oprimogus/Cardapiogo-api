package store

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/owner"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

const (
	ExistRegisteredStore = "Already exist this store."
	NotRegisteredStore   = "This store doesn't exist."
	IsNotOwner           = "You aren't the owner."
)

type Service struct {
	repository   Repository
	ownerService owner.Service
}

func NewService(repository Repository, ownerService owner.Service) *Service {
	return &Service{repository: repository, ownerService: ownerService}
}

func StoreWithIDExist(ctx context.Context, s *Service, ID string) (bool, error) {
	_, err := s.repository.GetStoreByID(ctx, ID)
	if err != nil {
		if err.Error() == errors.NOT_FOUND_RECORD {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func StoreWithCpfCnpjExist(ctx context.Context, s *Service, cpfCnpj string) (bool, error) {
	_, err := s.repository.GetStoreByCpfCnpj(ctx, cpfCnpj)
	if err != nil {
		if err.Error() == errors.NOT_FOUND_RECORD {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (s *Service) VerifyIsOwner(ctx context.Context, userID string, storeID string) (bool, error) {
	return s.ownerService.IsOwner(ctx, userID, storeID)
}

func (s *Service) CreateStore(ctx context.Context, params CreateStoreParams) error {
	existStore, err := StoreWithCpfCnpjExist(ctx, s, params.CpfCnpj)
	if err != nil {
		return err
	}
	if existStore {
		return errors.ConflictError(ExistRegisteredStore)
	}

	return s.repository.CreateStore(ctx, params)
}

func (s *Service) UpdateStore(ctx context.Context, params UpdateStoreParams, userID string) error {
	existStore, err := StoreWithIDExist(ctx, s, params.ExternalID)
	if err != nil {
		return err
	}
	if !existStore {
		return errors.NotFound(NotRegisteredStore)
	}
	isOwner, err := s.ownerService.IsOwner(ctx, userID, params.ExternalID)
	if err != nil {
		return err
	}
	if isOwner {
		return s.repository.UpdateStore(ctx, params)
	}
	return errors.Forbidden(IsNotOwner)
}

func (s *Service) UpdateStoreCpfCnpj(
	ctx context.Context,
	params UpdateStoreCpfCnpjParams,
	userID string,
) error {
	existStore, err := StoreWithIDExist(ctx, s, params.ExternalID)
	if err != nil {
		return err
	}
	if !existStore {
		return errors.NotFound(NotRegisteredStore)
	}
	isOwner, err := s.ownerService.IsOwner(ctx, userID, params.ExternalID)
	if err != nil {
		return err
	}
	if isOwner {
		return s.repository.UpdateStoreCpfCnpj(ctx, params)
	}
	return errors.Forbidden(IsNotOwner)
}

func (s *Service) DeleteStore(ctx context.Context, ID string, userID string) error {
	existStore, err := StoreWithIDExist(ctx, s, ID)
	if err != nil {
		return err
	}
	if !existStore {
		return errors.NotFound(NotRegisteredStore)
	}
	isOwner, err := s.ownerService.IsOwner(ctx, userID, ID)
	if err != nil {
		return err
	}
	if isOwner {
		return s.repository.DeleteStore(ctx, ID)
	}
	return errors.Forbidden(IsNotOwner)
}

func (s *Service) GetStoreByID(ctx context.Context, ID string) (StoreDetail, error) {
	return s.repository.GetStoreByID(ctx, ID)
}

func (s *Service) GetStoresListByFilter(
	ctx context.Context,
	filter StoreFilter,
) ([]StoreSummary, error) {
	stores, err := s.repository.GetStoresListByFilter(ctx, filter)
	if err != nil {
		return []StoreSummary{}, err
	}
	storeSummaryList := make([]StoreSummary, len(stores))
	for i, v := range stores {
		storeSummaryList[i] = StoreSummary{
			ExternalID: v.ExternalID,
			Name:       v.Name,
			Score:      v.Score,
			District:   v.District,
		}
	}
	return storeSummaryList, nil
}
