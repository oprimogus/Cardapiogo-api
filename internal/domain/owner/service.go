package owner

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListOwnersOfStore(ctx context.Context, storeID string) ([]Owner, error) {
	return s.repository.ListOwnersOfStore(ctx, storeID)
}

func (s *Service) AddOwnerOfStore(ctx context.Context, userID string, storeID string) error {
	return s.repository.AddOwner(ctx, userID, storeID)
}

func (s *Service) RemoveOwnerOfStore(ctx context.Context, userID string, storeID string) error {
	return s.repository.DeleteOwner(ctx, userID, storeID)
}

func (s *Service) IsOwner(ctx context.Context, userID string, storeID string) (bool, error) {
	value, err := s.repository.IsOwner(ctx, userID, storeID)
	return value, err
}
