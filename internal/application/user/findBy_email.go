package user

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type FindByEmail struct {
	userRepository repository.UserRepository
}

func NewFindByEmail(repository repository.UserRepository) FindByEmail {
	return FindByEmail{
		userRepository: repository,
	}
}

func (c FindByEmail) Execute(ctx context.Context, email string) (entity.User, error) {
	return c.userRepository.FindByEmail(ctx, email)
}
