package persistence

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/sqlc"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/services/authentication/keycloak"
)

type DatabaseRepositoryFactory struct {
	db      *postgres.PostgresDatabase
	querier *sqlc.Queries
}

func NewDataBaseRepositoryFactory(db *postgres.PostgresDatabase) repository.Factory {
	return &DatabaseRepositoryFactory{db: db, querier: sqlc.New(db.GetDB())}
}

func (d *DatabaseRepositoryFactory) NewUserRepository() repository.UserRepository {
	keycloakService, err := keycloak.NewKeycloakService(context.Background())
	if err != nil {
		panic(err)
	}
	return keycloakService
}

func (d *DatabaseRepositoryFactory) NewAuthenticationRepository() repository.AuthenticationRepository {
	keycloakService, err := keycloak.NewKeycloakService(context.Background())
	if err != nil {
		panic(err)
	}
	return keycloakService
}
