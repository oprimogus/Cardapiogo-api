package user

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type FindByID struct {
	userRepository repository.UserRepository
}

func NewFindByID(repository repository.UserRepository) FindByID {
	return FindByID{
		userRepository: repository,
	}
}

func (c FindByID) Execute(ctx context.Context, id string) (entity.User, error) {
	return c.userRepository.FindByID(ctx, id)
}
