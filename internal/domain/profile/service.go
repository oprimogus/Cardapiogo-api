package profile

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

const (
	ExistRegisteredProfile = "The user already has a registered profile."
)

type Service struct {
	repository     Repository
	userRepository user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *Service {
	return &Service{repository: repository, userRepository: userRepository}
}

func (s *Service) CreateProfile(ctx context.Context, userID string, params CreateProfileParams) error {
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user.ProfileID != 0 {
		return errors.ConflictError(ExistRegisteredProfile)
	}
	return s.repository.CreateProfile(ctx, userID, params)
}

func (s *Service) GetProfileByID(ctx context.Context, profileID int) (Profile, error) {
	return s.repository.GetProfileByID(ctx, profileID)
}

func (s *Service) GetProfileByUserID(ctx context.Context, userID string) (Profile, error) {
	return s.repository.GetProfileByUserID(ctx, userID)
}

func (s *Service) UpdateProfile(ctx context.Context, userID string, params UpdateProfileParams) error {
	return s.repository.UpdateProfile(ctx, userID, params)
}
