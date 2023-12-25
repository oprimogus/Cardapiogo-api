package infradatabase

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
)

// UserRepositoryDatabase struct
type UserRepositoryDatabase struct {
	db *postgres.PostgresDatabase
	q  sqlc.Querier
}

// NewUserRepositoryDatabase return repository of database
func NewUserRepositoryDatabase(db *postgres.PostgresDatabase, querier sqlc.Querier) user.Repository {
	return UserRepositoryDatabase{db: db, q: querier}
}

// CreateUser create a user in database. Must receive object validated through the service
func (u UserRepositoryDatabase) CreateUser(ctx context.Context, newUser user.CreateUserParams) error {

	user := sqlc.CreateUserParams{
		Email:           newUser.Email,
		Password:        pgtype.Text{String: newUser.Password, Valid: true},
		Role:            sqlc.UserRole(newUser.Role),
		AccountProvider: sqlc.AccountProvider(newUser.AccountProvider),
	}

	err := u.q.CreateUser(ctx, user)
	if err != nil {
		return postgres.MapDBError(err)
	}
	return nil

}
