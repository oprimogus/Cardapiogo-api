package user

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type FindByEmail interface {
	Execute(ctx context.Context, email string) (entity.User, error)
}

type findByEmail struct {
	userRepository repository.UserRepository
}

func NewFindByEmail(repository repository.UserRepository) FindByEmail {
	return findByEmail{
		userRepository: repository,
	}
}

func (c findByEmail) Execute(ctx context.Context, email string) (entity.User, error) {
	return c.userRepository.FindByEmail(ctx, email)
}
