package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	"github.com/oprimogus/cardapiogo/internal/domain/model"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var log = logger.GetLogger("UserService")

// UserService struct
type UserService struct {
	userRepository repository.UserRepository
}
// NewUserService UserService constructor
func NewUserService(repositoryFactory factory.RepositoryFactory) UserService {
	return UserService{userRepository: repositoryFactory.NewUserRepository()}
}

// CreateUser create a user in database
func (u UserService) CreateUser(newUser model.CreateUserParams) error {
	validate := validator.New()
	err := validate.Struct(newUser)
	if err != nil {
        // Cria uma mensagem de erro com detalhes dos erros de validação.
        var errorMessages string
        for _, err := range err.(validator.ValidationErrors) {
            errorMessages += fmt.Sprintf("Error: Field '%s' Condition '%s'\n", err.Field(), err.Tag())
        }
        // Retorna um erro com todas as mensagens de erro de validação.
        return errors.New(errorMessages)
    }
	errCreateUser := u.userRepository.CreateUser(context.Background(), newUser)
	return errCreateUser
}
