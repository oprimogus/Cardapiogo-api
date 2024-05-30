package infradatabase

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/oprimogus/cardapiogo/internal/domain/profile"
	"github.com/oprimogus/cardapiogo/internal/errors"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
	"github.com/oprimogus/cardapiogo/pkg/converters"
)

type ProfileRepositoryDatabase struct {
	db *postgres.PostgresDatabase
	q  sqlc.Querier
}

func NewProfileRepositoryDatabase(
	db *postgres.PostgresDatabase,
	querier sqlc.Querier,
) profile.Repository {
	return &ProfileRepositoryDatabase{db: db, q: querier}
}

func (p *ProfileRepositoryDatabase) CreateProfile(
	ctx context.Context,
	userID string,
	params profile.CreateProfileParams,
) error {
	tx, err := p.db.GetDB().Begin(ctx)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	defer tx.Rollback(ctx)
	sqlcParams := sqlc.CreateProfileAndReturnIDParams{
		Name:     params.Name,
		LastName: params.LastName,
		Cpf:      params.CPF,
		Phone:    params.Phone,
	}
	profileID, err := p.q.CreateProfileAndReturnID(ctx, sqlcParams)
	if err != nil {
		return errors.NewDatabaseError(err)
	}

	userUUID, err := converters.ConvertStringToUUID(userID)
	if err != nil {
		return err
	}

	updateUserProfileParams := sqlc.UpdateUserProfileParams{
		ID:        userUUID,
		ProfileID: converters.ConvertInt32ToInt4(profileID),
	}

	err = p.q.UpdateUserProfile(ctx, updateUserProfileParams)
	if err != nil {
		return errors.NewDatabaseError(err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

func (p *ProfileRepositoryDatabase) GetProfileByID(
	ctx context.Context,
	profileId int,
) (profile.Profile, error) {
	userProfile, err := p.q.GetProfileByID(ctx, int32(profileId))
	if err != nil {
		return profile.Profile{}, errors.NewDatabaseError(err)
	}
	return p.fromSqlcProfileToDomainProfile(sqlc.Profile{
		ID:        userProfile.ID,
		Name:      userProfile.Name,
		LastName:  userProfile.LastName,
		Cpf:       userProfile.Cpf,
		Phone:     userProfile.Phone,
		CreatedAt: pgtype.Timestamp{},
		UpdatedAt: pgtype.Timestamp{},
	}), nil
}

func (p *ProfileRepositoryDatabase) GetProfileByUserID(
	ctx context.Context,
	userID string,
) (profile.Profile, error) {
	uuid, err := converters.ConvertStringToUUID(userID)
	if err != nil {
		return profile.Profile{}, err
	}
	userProfile, err := p.q.GetProfileByUserID(ctx, uuid)
	if err != nil {
		return profile.Profile{}, errors.NewDatabaseError(err)
	}
	return p.fromSqlcProfileToDomainProfile(sqlc.Profile{
		ID:        userProfile.ID,
		Name:      userProfile.Name,
		LastName:  userProfile.LastName,
		Cpf:       userProfile.Cpf,
		Phone:     userProfile.Phone,
		CreatedAt: pgtype.Timestamp{},
		UpdatedAt: pgtype.Timestamp{},
	}), nil
}

func (p *ProfileRepositoryDatabase) UpdateProfile(
	ctx context.Context,
	userID string,
	params profile.UpdateProfileParams,
) error {
	uuid, err := converters.ConvertStringToUUID(userID)
	if err != nil {
		return err
	}
	convertedParams := sqlc.UpdateProfileParams{
		ID:       uuid,
		Name:     params.Name,
		LastName: params.LastName,
		Phone:    params.Phone,
	}
	err = p.q.UpdateProfile(ctx, convertedParams)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

func (p *ProfileRepositoryDatabase) fromSqlcProfileToDomainProfile(
	userProfile sqlc.Profile,
) profile.Profile {
	return profile.Profile{
		ID:        int(userProfile.ID),
		Name:      userProfile.Name,
		LastName:  userProfile.LastName,
		CPF:       userProfile.Cpf,
		Phone:     userProfile.Phone,
		CreatedAt: userProfile.CreatedAt.Time,
		UpdatedAt: userProfile.UpdatedAt.Time,
	}
}
