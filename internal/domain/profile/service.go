package profile

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

type Service struct {
	repository Repository
	userRepository user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *Service {
	return &Service{repository: repository, userRepository: userRepository}
}

func (u *Service) CreateProfile(ctx context.Context, userID string, params CreateProfileParams) error {
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user.ProfileID != 0 {
		return errors.ConflictError("The user already has a registered profile")
	}
	return u.repository.CreateProfile(ctx, userID, params)
}