package factory

import (
	"github.com/oprimogus/cardapiogo/internal/domain/profile"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
)

// RepositoryFactory interface
type RepositoryFactory interface {
	NewUserRepository() user.Repository
	NewProfileRepository() profile.Repository
}
