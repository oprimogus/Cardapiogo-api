package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type ResetPassword struct {
	userRepository repository.UserRepository
}

func NewResetPassword(repository repository.UserRepository) ResetPassword {
	return ResetPassword{
		userRepository: repository,
	}
}

func (r ResetPassword) Execute(ctx context.Context, id string) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	return r.userRepository.ResetPasswordByEmail(ctx, id)
}
