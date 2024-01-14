package infradatabase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
	"github.com/oprimogus/cardapiogo/pkg/converters"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

// UserRepositoryDatabase struct
type UserRepositoryDatabase struct {
	db *postgres.PostgresDatabase
	q  sqlc.Querier
}

// NewUserRepositoryDatabase return repository of database
func NewUserRepositoryDatabase(db *postgres.PostgresDatabase, querier sqlc.Querier) user.Repository {
	return &UserRepositoryDatabase{db: db, q: querier}
}

// CreateUser create a user in database. Must receive object validated through the service
func (u UserRepositoryDatabase) CreateUser(ctx context.Context, newUser user.CreateUserParams) error {
	err := u.q.CreateUser(ctx, u.createUserDatabase(newUser))
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

// GetUserByID return a user of database. Must receive object validated through the service
func (u UserRepositoryDatabase) GetUserByID(ctx context.Context, id string) (user.User, error) {
	log := logger.GetLogger("UserRepository", ctx)
	uuid, err := converters.ConvertStringToUUID(id)
	if err != nil {
		return user.User{}, err
	}

	getUser, err := u.q.GetUserById(ctx, pgtype.UUID{Bytes: uuid.Bytes, Valid: true})
	if err != nil {
		log.Error(err)
		return user.User{}, errors.NewDatabaseError(err)
	}

	return u.fromUserByIDRowToUserModel(getUser)
}

// GetUsersList return a list of users
func (u UserRepositoryDatabase) GetUsersList(ctx context.Context, items int, page int) ([]user.User, error) {
	log := logger.GetLogger("UserRepository", ctx)
	offset := items * (page - 1)

	searchParams := sqlc.GetUserParams{
		Limit:  int32(items),
		Offset: int32(offset),
	}

	listUsers, err := u.q.GetUser(ctx, searchParams)
	if err != nil {
		log.Error(err)
		return nil, errors.NewDatabaseError(err)
	}

	listUsersModel := make([]user.User, len(listUsers))
	for i, v := range listUsers {
		user, err := u.fromUserRowToUserModel(v)
		log.Info(user.Email)
		if err != nil {
			return nil, err
		}
		listUsersModel[i] = user
	}
	return listUsersModel, nil
}

// UpdateUserPassword update hash of password
func (u UserRepositoryDatabase) UpdateUserPassword(ctx context.Context, updateParams user.UpdateUserPasswordParams) error {
	log := logger.GetLogger("UserRepository", ctx)
	updateUserPasswordParams, err := u.createUpdateUserPasswordsParams(updateParams)
	if err != nil {
		return err
	}
	err = u.q.UpdateUserPassword(ctx, updateUserPasswordParams)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// UpdateUser update data user with the exception of password
func (u UserRepositoryDatabase) UpdateUser(ctx context.Context, updateParams user.UpdateUserParams) error {
	log := logger.GetLogger("UserRepository", ctx)
	params, err := u.createUpdateUserParams(updateParams)
	if err != nil {
		log.Error(err)
		return err
	}
	return u.q.UpdateUser(ctx, params)
}

// createUserDatabase format DTO to CreateUserParams of sqlc
func (u UserRepositoryDatabase) createUserDatabase(cr user.CreateUserParams) sqlc.CreateUserParams {

	passwordSqlc := converters.ConvertStringToText(cr.Password)

	return sqlc.CreateUserParams{
		Email:           cr.Email,
		Password:        passwordSqlc,
		Role:            sqlc.UserRole(cr.Role),
		AccountProvider: sqlc.AccountProvider(cr.AccountProvider),
	}
}

// createUpdateUserPasswordsParams format DTO to UpdateUserPasswordParams of sqlc
func (u UserRepositoryDatabase) createUpdateUserPasswordsParams(params user.UpdateUserPasswordParams) (sqlc.UpdateUserPasswordParams, error) {
	convertedID, err := converters.ConvertStringToUUID(params.ID)
	if err != nil {
		return sqlc.UpdateUserPasswordParams{}, err
	}
	return sqlc.UpdateUserPasswordParams{
		ID:       convertedID,
		Password: converters.ConvertStringToText(params.Password),
	}, nil
}

// createUpdateUserParams format DTO to UpdateUserPasswordParams of sqlc
func (u UserRepositoryDatabase) createUpdateUserParams(params user.UpdateUserParams) (sqlc.UpdateUserParams, error) {
	convertedID, err := converters.ConvertStringToUUID(params.ID)
	if err != nil {
		return sqlc.UpdateUserParams{}, err
	}
	return sqlc.UpdateUserParams{
		ID:    convertedID,
		Email: params.Email,
		Role:  sqlc.UserRole(params.Role),
	}, nil
}

// fromUserByIDRowToUserModel convert sqlc.GetUserByIdRow to user.User
func (u UserRepositoryDatabase) fromUserByIDRowToUserModel(su sqlc.GetUserByIdRow) (user.User, error) {

	uuidStr, err := converters.ConvertUUIDToString(su.ID)
	if err != nil {
		return user.User{}, err
	}
	uuidValue := ""
	if uuidStr != nil {
		uuidValue = *uuidStr
	}

	passwordValue, err := converters.ConvertTextToString(su.Password)
	if err != nil {
		return user.User{}, err
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
		ID:        uuidValue,
		ProfileID: profileIDValue,
		Email:     su.Email,
		Password:  *passwordValue,
		Role:      types.Role(su.Role),
		CreatedAt: createdAtValue,
		UpdatedAt: updatedAtValue,
	}

	return user, nil
}

// fromUserRowToUserModel convert sqlc.getUserRow to user.User
func (u UserRepositoryDatabase) fromUserRowToUserModel(su sqlc.GetUserRow) (user.User, error) {
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
		ID:        uuidValue,
		ProfileID: profileIDValue,
		Email:     su.Email,
		Role:      types.Role(su.Role),
		CreatedAt: createdAtValue,
		UpdatedAt: updatedAtValue,
	}
	return user, nil
}
