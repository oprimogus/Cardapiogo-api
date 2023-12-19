package service

import (
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

// import "github.com/oprimogus/cardapiogo/internal/domain/factory"

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repositoryFactory factory.RepositoryFactory) UserService {
	return UserService{ userRepository: repositoryFactory.NewUserRepository() }
}

func CreateUser() {
	
}