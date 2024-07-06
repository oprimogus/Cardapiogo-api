package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

var ErrExistUserWithEmail = errors.New("exist user with this email")
var ErrExistUserWithDocument = errors.New("exist user with this document")
var ErrExistUserWithPhone = errors.New("exist user with this phone")

type Create interface {
	Execute(ctx context.Context, input CreateParams) error
}

type create struct {
	userRepository repository.UserRepository
}

func NewCreate(repository repository.UserRepository) Create {
	return create{
		userRepository: repository,
	}
}

func (c create) Execute(ctx context.Context, input CreateParams) error {
	existUser, err := c.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return err
	}
	if existUser.Email == input.Email {
		return ErrExistUserWithEmail
	}
	if existUser.Profile.Document == input.Profile.Document {
		return ErrExistUserWithDocument
	}
	if existUser.Profile.Phone == input.Profile.Phone {
		return ErrExistUserWithPhone
	}
	return c.userRepository.Create(ctx, input.ToEntity())
}
