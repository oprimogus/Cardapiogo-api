package user

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type Create struct {
	userRepository repository.UserRepository
}

func NewCreate(repository repository.UserRepository) Create {
	return Create{
		userRepository: repository,
	}
}

func (c Create) Execute(ctx context.Context, input CreateParams) error {
	existUser, err := c.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return err
	}
	if existUser.Email == input.Email {
		return entity.ErrExistUserWithEmail
	}
	if existUser.Profile.Document == input.Profile.Document {
		return entity.ErrExistUserWithDocument
	}
	if existUser.Profile.Phone == input.Profile.Phone {
		return entity.ErrExistUserWithPhone
	}
	return c.userRepository.Create(ctx, input.ToEntity())
}
