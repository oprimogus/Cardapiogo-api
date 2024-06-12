package user

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type FindByID interface {
	Execute(ctx context.Context, id string) (entity.User, error)
}

type findByID struct {
	userRepository repository.UserRepository
}

func NewFindByID(repository repository.UserRepository) FindByID {
	return findByID{
		userRepository: repository,
	}
}

func (c findByID) Execute(ctx context.Context, id string) (entity.User, error) {
	return c.userRepository.FindByID(ctx, id)
}
