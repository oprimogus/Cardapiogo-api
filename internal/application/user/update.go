package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Update struct {
	userRepository repository.UserRepository
}

func NewUpdate(repository repository.UserRepository) Update {
	return Update{
		userRepository: repository,
	}
}

func (u Update) Execute(ctx context.Context, input UpdateProfileParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	user := input.ToEntity()
	user.ID = userID
	return u.userRepository.Update(ctx, user)
}
