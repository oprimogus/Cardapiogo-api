package factory

import (
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type RepositoryFactory interface {
	NewUserRepository() repository.UserRepository
}
