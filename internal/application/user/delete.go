package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Delete interface {
	Execute(ctx context.Context) error
}

type delete struct {
	userRepository repository.UserRepository
}

func NewDelete(repository repository.UserRepository) Delete {
	return delete{userRepository: repository}
}

func (d delete) Execute(ctx context.Context) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	return d.userRepository.Delete(ctx, userID)
}
