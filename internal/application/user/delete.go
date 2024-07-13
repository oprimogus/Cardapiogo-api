package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Delete struct {
	userRepository repository.UserRepository
}

func NewDelete(repository repository.UserRepository) Delete {
	return Delete{userRepository: repository}
}

func (d Delete) Execute(ctx context.Context) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	return d.userRepository.Delete(ctx, userID)
}
