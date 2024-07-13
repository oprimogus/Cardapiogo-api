package persistence

import (
	"context"
	"time"

	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/database/sqlc"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/services/authentication/keycloak"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	keycloakService *keycloak.KeycloakService
	log             = logger.NewLogger("RepositoryFactory")
)

func init() {
	hasKeycloak := false
	for !hasKeycloak {
		keycloakInstance, err := keycloak.NewKeycloakService(context.Background())
		if err != nil {
			log.Errorf("Fail in generate keycloak instance: %s", err)
			time.Sleep(5 * time.Second)
		} else {
			hasKeycloak = true
			keycloakService = keycloakInstance
		}
	}
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

func (d *DatabaseRepositoryFactory) NewStoreRepository() repository.StoreRepository {
	return NewStoreRepository(d.db, d.querier)
}
