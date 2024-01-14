package sqlc

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/pkg/converters"
)

type UserDatabase struct {
	ID        pgtype.UUID        `db:"id" json:"id"`
	ProfileID pgtype.Int4        `db:"profile_id" json:"profile_id"`
	Email     string             `db:"email" json:"email"`
	Password  pgtype.Text        `db:"password" json:"password"`
	Role      UserRole           `db:"role" json:"role"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

// FromUserDatabaseToUserModel convert sqlc.user to user.User
func FromUserDatabaseToUserModel(su UserDatabase) (user.User, error) {
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

	passwordValue, err := converters.ConvertTextToString(su.Password)
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
		ID:        uuidValue,
		ProfileID: profileIDValue,
		Email:     su.Email,
		Password:  passwordValue,
		Role:      types.Role(su.Role),
		CreatedAt: createdAtValue,
		UpdatedAt: updatedAtValue,
	}
	return user, nil
}
