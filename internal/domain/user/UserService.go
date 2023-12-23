package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var log = logger.GetLogger("UserService")

// Service struct
type Service struct {
	repository Repository
}

// NewService Service constructor
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// CreateUser create a user in database
func (u Service) CreateUser(ctx context.Context, newUser CreateUserParams) error {
	validate := validator.New()
	err := validate.Struct(newUser)
	if err != nil {
		var errorMessages string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages += fmt.Sprintf("Error: Field '%s' Condition '%s'\n", err.Field(), err.Tag())
		}
		return errors.New(errorMessages)
	}
	errCreateUser := u.repository.CreateUser(ctx, newUser)
	return errCreateUser
}
