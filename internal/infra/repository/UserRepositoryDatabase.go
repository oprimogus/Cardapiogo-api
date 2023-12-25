package infradatabase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	errordatabase "github.com/oprimogus/cardapiogo/internal/errors/database"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
	"github.com/oprimogus/cardapiogo/pkg/converters"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var log = logger.GetLogger("UserRepository")

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
	err := u.q.CreateUser(ctx, u.createUserDatabase(newUser))
	log.Error(err)
	if err != nil {
		return errordatabase.MapDBError(err)
	}
	return nil
}

func (u UserRepositoryDatabase) createUserDatabase(cr user.CreateUserParams) (sqlc.CreateUserParams) {

	passwordSqlc := converters.ConvertStringToText(cr.Password)

	return sqlc.CreateUserParams{
		Email:           cr.Email,
		Password:        passwordSqlc,
		Role:            sqlc.UserRole(cr.Role),
		AccountProvider: sqlc.AccountProvider(cr.AccountProvider),
	}
}

// GetUser return a user of database. Must receive object validated through the service
func (u UserRepositoryDatabase) GetUser(ctx context.Context, id uuid.UUID) (user.User, error) {
	uuidBytes := id[:]

	getUser, err := u.q.GetUser(ctx, pgtype.UUID{Bytes: [16]byte(uuidBytes), Valid: true})
	if err != nil {
		return user.User{}, errordatabase.MapDBError(err)
	}

	return u.createUserModel(getUser)
}

// createUserModel convert sqlc.User to user.User
func (u UserRepositoryDatabase) createUserModel(su sqlc.Users) (user.User, error) {

	uuidStr, err := converters.ConvertUUIDToString(su.ID)
	if err != nil {
		return user.User{}, err
	}
	uuidValue := ""
	if uuidStr != nil {
		uuidValue = *uuidStr
	}

	profileIDPtr, err := converters.ConvertInt4ToInt(su.ProfileID)
	if err != nil {
		return user.User{}, err
	}
	var profileIDValue *int
	if profileIDPtr != nil {
		profileIDValue = profileIDPtr
	}

	passwordPtr, err := converters.ConvertTextToString(su.Password)
	if err != nil {
		return user.User{}, err
	}
	var passwordValue *string
	if passwordPtr != nil {
		passwordValue = passwordPtr
	}

	createdAtPtr, err := converters.ConvertTimestamptzToTime(su.CreatedAt)
	if err != nil {
		return user.User{}, err
	}
	var createdAtValue time.Time
	if createdAtPtr != nil {
		createdAtValue = *createdAtPtr
	}

	updatedAtPtr, err := converters.ConvertTimestamptzToTime(su.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}
	var updatedAtValue time.Time
	if updatedAtPtr != nil {
		updatedAtValue = *updatedAtPtr
	}

	user := user.User{
		ID:              uuidValue,
		ProfileID:       profileIDValue,
		Email:           su.Email,
		Password:        passwordValue,
		Role:            types.Role(su.Role),
		AccountProvider: types.AccountProvider(su.AccountProvider),
		CreatedAt:       createdAtValue,
		UpdatedAt:       updatedAtValue,
	}

	return user, nil

}
