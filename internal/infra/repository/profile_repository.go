package infradatabase

import (
	"context"

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

func NewProfileRepositoryDatabase(db *postgres.PostgresDatabase, querier sqlc.Querier) profile.Repository {
	return &ProfileRepositoryDatabase{db: db, q: querier}
}

func (p *ProfileRepositoryDatabase) CreateProfile(ctx context.Context, userID string, params profile.CreateProfileParams) error {
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
