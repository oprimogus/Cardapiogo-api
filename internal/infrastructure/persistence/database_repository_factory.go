package persistence

import (
	"context"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/sqlc"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/services/authentication/keycloak"
)

var keycloakService *keycloak.KeycloakService

func init() {
	keycloakInstance, err := keycloak.NewKeycloakService(context.Background())
	if err != nil {
		panic(err)
	}
	keycloakService = keycloakInstance
}

type DatabaseRepositoryFactory struct {
	db      *postgres.PostgresDatabase
	querier *sqlc.Queries
}

func NewDataBaseRepositoryFactory(db *postgres.PostgresDatabase) repository.Factory {
	return &DatabaseRepositoryFactory{db: db, querier: sqlc.New(db.GetDB())}
}

func (d *DatabaseRepositoryFactory) NewUserRepository() repository.UserRepository {
	return keycloakService
}

func (d *DatabaseRepositoryFactory) NewAuthenticationRepository() repository.AuthenticationRepository {
	return keycloakService
}
