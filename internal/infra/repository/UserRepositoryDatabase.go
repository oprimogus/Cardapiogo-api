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

// GetUserByID return a user of database. Must receive object validated through the service
func (u UserRepositoryDatabase) GetUserByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	uuidBytes := id[:]

	getUser, err := u.q.GetUserById(ctx, pgtype.UUID{Bytes: [16]byte(uuidBytes), Valid: true})
	if err != nil {
		return user.User{}, errordatabase.MapDBError(err)
	}

	return u.getUserByIDRowToUserModel(getUser)
}

// GetUsersList return a list of users
func (u UserRepositoryDatabase) GetUsersList (ctx context.Context, items int, page int) ([]user.User, error) {

	offset := items * (page - 1)

	searchParams := sqlc.GetUserParams {
		Limit: int32(items),
		Offset: int32(offset),
	}

	listUsers, err := u.q.GetUser(ctx, searchParams)
	if err != nil {
		return nil, errordatabase.MapDBError(err)
	}

	listUsersModel := make([]user.User, len(listUsers))
	for i, v := range listUsers {
		log.Info(v)
		user, err := u.getUserRowToUserModel(v)
		if err != nil {
			return nil, err
		}
		listUsersModel[i] = user
	}
	return listUsersModel, nil
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

// getUserByIDRowToUserModel convert sqlc.GetUserByIdRow to user.User
func (u UserRepositoryDatabase) getUserByIDRowToUserModel(su sqlc.GetUserByIdRow) (user.User, error) {

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
		Role:            types.Role(su.Role),
		CreatedAt:       createdAtValue,
		UpdatedAt:       updatedAtValue,
	}

	return user, nil
}

// getUserRowToUserModel convert sqlc.getUserRow to user.User
func (u UserRepositoryDatabase) getUserRowToUserModel(su sqlc.GetUserRow) (user.User, error) {
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
		Role:            types.Role(su.Role),
		CreatedAt:       createdAtValue,
		UpdatedAt:       updatedAtValue,
	}
	return user, nil
}