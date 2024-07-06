package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Update interface {
	Execute(ctx context.Context, input UpdateProfileParams) error
}

type update struct {
	userRepository repository.UserRepository
}

func NewUpdate(repository repository.UserRepository) Update {
	return update{
		userRepository: repository,
	}
}

func (u update) Execute(ctx context.Context, input UpdateProfileParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	return u.userRepository.Update(ctx, input.ToEntity())
}
