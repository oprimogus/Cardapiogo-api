package persistence

import (
	"context"
	"time"

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

func NewDataBaseRepositoryFactory(db *postgres.PostgresDatabase) core.RepositoryFactory {
	return &DatabaseRepositoryFactory{db: db, querier: sqlc.New(db.GetDB())}
}

func (d *DatabaseRepositoryFactory) NewUserRepository() user.Repository {
	return keycloakService
}

func (d *DatabaseRepositoryFactory) NewAuthenticationRepository() authentication.Repository {
	return keycloakService
}

func (d *DatabaseRepositoryFactory) NewStoreRepository() store.Repository {
	return NewStoreRepository(d.db, d.querier)
}
