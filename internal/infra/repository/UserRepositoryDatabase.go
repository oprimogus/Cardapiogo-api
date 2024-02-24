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
func (u *UserRepositoryDatabase) CreateUser(ctx context.Context, newUser user.CreateUserParams) error {
	err := u.q.CreateUser(ctx, u.createUserDatabase(newUser))
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

// CreateUserWithOAuth create a user in database. Must receive object validated through the service
func (u *UserRepositoryDatabase) CreateUserWithOAuth(ctx context.Context, newUser user.CreateUserWithOAuthParams) error {
	err := u.q.CreateUserWithOAuth(ctx, u.createUserOAuthDatabase(newUser))
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

// GetUserByID return a user of database. Must receive object validated through the service
func (u *UserRepositoryDatabase) GetUserByID(ctx context.Context, id string) (user.User, error) {
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

	return u.fromSqlcUserToUserModel(getUser)
}

// GetUserByEmail return a user of database.
func (u *UserRepositoryDatabase) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	log := logger.GetLogger("UserRepository", ctx)

	getUser, err := u.q.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(err)
		return user.User{}, errors.NewDatabaseError(err)
	}

	return u.fromSqlcUserToUserModel(getUser)
}

// GetUsersList return a list of users
func (u *UserRepositoryDatabase) GetUsersList(ctx context.Context, items int, page int) ([]*user.User, error) {
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

	listUsersModel := make([]*user.User, len(listUsers))
	for i, v := range listUsers {
		user, err := u.fromSqlcUserToUserModel(u.fromGetUserRowToSqlcUser(v))
		if err != nil {
			return nil, err
		}
		listUsersModel[i] = &user
	}
	return listUsersModel, nil
}

// UpdateUserPassword update hash of password
func (u *UserRepositoryDatabase) UpdateUserPassword(ctx context.Context, updateParams user.UpdateUserPasswordParams) error {
	updateUserPasswordParams, err := u.createUpdateUserPasswordsParams(updateParams)
	if err != nil {
		return err
	}
	err = u.q.UpdateUserPassword(ctx, updateUserPasswordParams)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser update data user with the exception of password
func (u *UserRepositoryDatabase) UpdateUser(ctx context.Context, updateParams user.UpdateUserParams) error {
	params, err := u.createUpdateUserParams(updateParams)
	if err != nil {
		return err
	}
	return u.q.UpdateUser(ctx, params)
}

// createUserDatabase format DTO to CreateUserParams of sqlc
func (u *UserRepositoryDatabase) createUserDatabase(cr user.CreateUserParams) sqlc.CreateUserParams {

	passwordSqlc := converters.ConvertStringToText(cr.Password)

	return sqlc.CreateUserParams{
		Email:           cr.Email,
		Password:        passwordSqlc,
		Role:            sqlc.UserRole(cr.Role),
		AccountProvider: sqlc.AccountProvider(cr.AccountProvider),
	}
}

// createUserOAuthDatabase format DTO to CreateUserWithOAuthParams of sqlc
func (u *UserRepositoryDatabase) createUserOAuthDatabase(cr user.CreateUserWithOAuthParams) sqlc.CreateUserWithOAuthParams {

	return sqlc.CreateUserWithOAuthParams{
		Email:           cr.Email,
		Role:            sqlc.UserRole(cr.Role),
		AccountProvider: sqlc.AccountProvider(cr.AccountProvider),
	}
}

// createUpdateUserPasswordsParams format DTO to UpdateUserPasswordParams of sqlc
func (u *UserRepositoryDatabase) createUpdateUserPasswordsParams(params user.UpdateUserPasswordParams) (sqlc.UpdateUserPasswordParams, error) {
	convertedID, err := converters.ConvertStringToUUID(params.ID)
	if err != nil {
		return sqlc.UpdateUserPasswordParams{}, err
	}
	return sqlc.UpdateUserPasswordParams{
		ID:       convertedID,
		Password: converters.ConvertStringToText(params.NewPassword),
	}, nil
}

// createUpdateUserParams format DTO to UpdateUserPasswordParams of sqlc
func (u *UserRepositoryDatabase) createUpdateUserParams(params user.UpdateUserParams) (sqlc.UpdateUserParams, error) {
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

// fromSqlcUserToUserModel convert sqlc.User to user.User
func (u *UserRepositoryDatabase) fromSqlcUserToUserModel(su sqlc.Users) (user.User, error) {

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

	profileIDValue, err := converters.ConvertInt4ToInt(su.ProfileID)
	if err != nil {
		return user.User{}, err
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

// fromGetUserRowToSqlcUser convert sqlc.GetUserRow to user.Users
func (u *UserRepositoryDatabase) fromGetUserRowToSqlcUser(su sqlc.GetUserRow) sqlc.Users {
	return sqlc.Users{
		ID:        su.ID,
		ProfileID: su.ProfileID,
		Email:     su.Email,
		Password:  pgtype.Text{Valid: false, String: ""},
		Role:      su.Role,
		AccountProvider: su.AccountProvider,
		CreatedAt: su.CreatedAt,
		UpdatedAt: su.UpdatedAt,
	}
}
