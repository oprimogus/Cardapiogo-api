package user

import (
	"context"
	"errors"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

const (
	ExistUserWithEmail    = "Exist user with this email."
	ExistUserWithDocument = "Exist user with this document."
	ExistUserWithPhone    = "Exist user with this phone."
)

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
		return errors.New(ExistUserWithEmail)
	}
	if existUser.Profile.Document == input.Profile.Document {
		return errors.New(ExistUserWithDocument)
	}
	if existUser.Profile.Phone == input.Profile.Phone {
		return errors.New(ExistUserWithPhone)
	}
	return c.userRepository.Create(ctx, input.ToEntity())
}
