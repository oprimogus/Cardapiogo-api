package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type UpdatePassword interface {
	Execute(ctx context.Context, oldPassword, newPassword string) error
}

type updatePassword struct {
	userRepository repository.UserRepository
}

func NewUpdatePassword(repository repository.UserRepository) UpdatePassword {
	return updatePassword{
		userRepository: repository,
	}
}

func (u updatePassword) Execute(ctx context.Context, oldPassword, newPassword string) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("Invalid user.")
	}
	return u.userRepository.UpdatePassword(ctx, oldPassword, newPassword)
}
