package infradatabase

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/oprimogus/cardapiogo/internal/domain/model"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	"github.com/oprimogus/cardapiogo/internal/infra/database"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
)

// UserRepositoryDatabase struct
type UserRepositoryDatabase struct {
	db *database.PostgresDatabase
	q  sqlc.Querier
}

// NewUserRepositoryDatabase return repository of database
func NewUserRepositoryDatabase(db *database.PostgresDatabase, querier sqlc.Querier) repository.UserRepository {
	return UserRepositoryDatabase{db: db, q: querier}
}

// CreateUser create a user in database. Must receive object validated through the service
func (u UserRepositoryDatabase) CreateUser(ctx context.Context, newUser model.CreateUserParams) error {

	user := sqlc.CreateUserParams{
		Email:           newUser.Email,
		Password:        pgtype.Text{String: newUser.Password, Valid: true},
		Role:            sqlc.UserRole(newUser.Role),
		AccountProvider: sqlc.AccountProvider(newUser.AccountProvider),
	}

	err := u.q.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
