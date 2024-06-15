package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type ResetPassword interface {
	Execute(ctx context.Context, id string) error
}

type resetPassword struct {
	userRepository repository.UserRepository
}

func NewResetPassword(repository repository.UserRepository) ResetPassword {
	return resetPassword{
		userRepository: repository,
	}
}

func (r resetPassword) Execute(ctx context.Context, id string) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("Invalid user.")
	}
	return r.userRepository.ResetPasswordByEmail(ctx, id)
}
