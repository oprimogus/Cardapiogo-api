package persistence

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/core"
	"github.com/oprimogus/cardapiogo/internal/core/authentication"
	"github.com/oprimogus/cardapiogo/internal/core/store"
	"github.com/oprimogus/cardapiogo/internal/core/user"
	"github.com/oprimogus/cardapiogo/internal/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/database/sqlc"
	"github.com/oprimogus/cardapiogo/internal/services/authentication/keycloak"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	log = logger.NewLogger("RepositoryFactory")
)

type DatabaseRepositoryFactory struct {
	db      *postgres.PostgresDatabase
	querier *sqlc.Queries
}

func NewDataBaseRepositoryFactory(db *postgres.PostgresDatabase) core.RepositoryFactory {
	return &DatabaseRepositoryFactory{db: db, querier: sqlc.New(db.GetDB())}
}

func (d *DatabaseRepositoryFactory) NewUserRepository() user.Repository {
	return keycloak.GetInstance(context.Background())
}

func (d *DatabaseRepositoryFactory) NewAuthenticationRepository() authentication.Repository {
	return keycloak.GetInstance(context.Background())
}

func (d *DatabaseRepositoryFactory) NewStoreRepository() store.Repository {
	return NewStoreRepository(d.db, d.querier)
}
