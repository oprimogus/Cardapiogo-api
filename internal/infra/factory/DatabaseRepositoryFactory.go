package factory

import (
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
	infradatabase "github.com/oprimogus/cardapiogo/internal/infra/repository"
)

// DatabaseRepositoryFactory struct
type DatabaseRepositoryFactory struct {
	db      *postgres.PostgresDatabase
	querier *sqlc.Queries
}

// NewDataBaseRepositoryFactory return repository factory
func NewDataBaseRepositoryFactory(db *postgres.PostgresDatabase) factory.RepositoryFactory {
	return DatabaseRepositoryFactory{db: db, querier: sqlc.New(db.GetDB())}
}

// NewUserRepository return a user repository
func (f DatabaseRepositoryFactory) NewUserRepository() user.Repository {
	return infradatabase.NewUserRepositoryDatabase(f.db, f.querier)
}
