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

type UserRepositoryDatabase struct {
	db *postgres.PostgresDatabase
	q  sqlc.Querier
}

func NewUserRepositoryDatabase(db *postgres.PostgresDatabase, querier sqlc.Querier) user.Repository {
	return &UserRepositoryDatabase{db: db, q: querier}
}

func (u *UserRepositoryDatabase) CreateUser(ctx context.Context, params user.CreateUserParams) error {
	err := u.q.CreateUser(ctx, u.createUserDatabase(params))
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

func (u *UserRepositoryDatabase) CreateUserWithOAuth(ctx context.Context, params user.CreateUserWithOAuthParams) error {
	err := u.q.CreateUserWithOAuth(ctx, u.createUserOAuthDatabase(params))
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

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

	return u.fromSqlcUserToUserModel(u.fromGetUserByIDRowToSqlcUser(getUser))
}

func (u *UserRepositoryDatabase) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	log := logger.GetLogger("UserRepository", ctx)

	getUser, err := u.q.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(err)
		return user.User{}, errors.NewDatabaseError(err)
	}

	return u.fromSqlcUserToUserModel(u.fromGetUserByEmailRowToSqlcUser(getUser))
}

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

func (u *UserRepositoryDatabase) UpdateUserPassword(ctx context.Context, params user.UpdateUserPasswordParams) error {
	updateUserPasswordParams, err := u.createUpdateUserPasswordsParams(params)
	if err != nil {
		return err
	}
	err = u.q.UpdateUserPassword(ctx, updateUserPasswordParams)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

func (u *UserRepositoryDatabase) UpdateUser(ctx context.Context, params user.UpdateUserParams) error {
	updateParams, err := u.createUpdateUserParams(params)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return u.q.UpdateUser(ctx, updateParams)
}

func (u *UserRepositoryDatabase) UpdateUserProfile(ctx context.Context, params user.UpdateUserProfileParams) error {
	convertedParams, err := u.createUpdateUserProfileParams(params)
	if err != nil {
		return err
	}
	err = u.q.UpdateUserProfile(ctx, convertedParams)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

func (u *UserRepositoryDatabase) createUserDatabase(cr user.CreateUserParams) sqlc.CreateUserParams {

	passwordSqlc := converters.ConvertStringToText(cr.Password)

	return sqlc.CreateUserParams{
		Email:           cr.Email,
		Password:        passwordSqlc,
		Role:            sqlc.UserRole(cr.Role),
		AccountProvider: sqlc.AccountProvider(cr.AccountProvider),
	}
}

func (u *UserRepositoryDatabase) createUserOAuthDatabase(cr user.CreateUserWithOAuthParams) sqlc.CreateUserWithOAuthParams {

	return sqlc.CreateUserWithOAuthParams{
		Email:           cr.Email,
		Role:            sqlc.UserRole(cr.Role),
		AccountProvider: sqlc.AccountProvider(cr.AccountProvider),
	}
}

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

	createdAtPtr, err := converters.ConvertTimestamptToTime(su.CreatedAt)
	if err != nil {
		return user.User{}, err
	}
	var createdAtValue time.Time
	if createdAtPtr != nil {
		createdAtValue = *createdAtPtr
	}

	updatedAtPtr, err := converters.ConvertTimestamptToTime(su.UpdatedAt)
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

func (u *UserRepositoryDatabase) fromGetUserRowToSqlcUser(su sqlc.GetUserRow) sqlc.Users {
	return sqlc.Users{
		ID:              su.ID,
		ProfileID:       su.ProfileID,
		Email:           su.Email,
		Password:        pgtype.Text{Valid: false, String: ""},
		Role:            su.Role,
		AccountProvider: su.AccountProvider,
		CreatedAt:       su.CreatedAt,
		UpdatedAt:       su.UpdatedAt,
	}
}

func (u *UserRepositoryDatabase) fromGetUserByEmailRowToSqlcUser(su sqlc.GetUserByEmailRow) sqlc.Users {
	return sqlc.Users{
		ID:              su.ID,
		ProfileID:       su.ProfileID,
		Email:           su.Email,
		Password:        pgtype.Text{Valid: false, String: su.Password.String},
		Role:            su.Role,
		AccountProvider: su.AccountProvider,
		CreatedAt:       su.CreatedAt,
		UpdatedAt:       su.UpdatedAt,
	}
}

func (u *UserRepositoryDatabase) fromGetUserByIDRowToSqlcUser(su sqlc.GetUserByIdRow) sqlc.Users {
	return sqlc.Users{
		ID:              su.ID,
		ProfileID:       su.ProfileID,
		Email:           su.Email,
		Password:        pgtype.Text{Valid: false, String: su.Password.String},
		Role:            su.Role,
		AccountProvider: su.AccountProvider,
		CreatedAt:       su.CreatedAt,
		UpdatedAt:       su.UpdatedAt,
	}
}

func (u *UserRepositoryDatabase) createUpdateUserProfileParams(su user.UpdateUserProfileParams) (sqlc.UpdateUserProfileParams, error) {
	id, err := converters.ConvertStringToUUID(su.ID)
	if err != nil {
		return sqlc.UpdateUserProfileParams{}, err
	}
	profileId := converters.ConvertIntToInt4(su.ProfileID)

	return sqlc.UpdateUserProfileParams{
		ID:        id,
		ProfileID: profileId,
	}, nil
}
