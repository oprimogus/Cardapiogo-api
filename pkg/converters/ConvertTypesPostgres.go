package converters

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// ConvertUUIDToString converts pgtype.UUID to *string. Returns nil if UUID is not valid.
func ConvertUUIDToString(uuidVal pgtype.UUID) (*string, error) {
	if !uuidVal.Valid {
		return nil, nil
	}

	u, err := uuid.FromBytes(uuidVal.Bytes[:])
	if err != nil {
		return nil, err
	}

	str := u.String()
	return &str, nil
}

// ConvertStringToUUID converts string to pgtype.UUID.
func ConvertStringToUUID(str string) (pgtype.UUID, error) {
	u, err := uuid.Parse(str)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return pgtype.UUID{
		Bytes: u,
		Valid: true,
	}, nil
}

// ConvertInt4ToInt converts pgtype.Int4 to *int. Returns nil if Int4 is not valid.
func ConvertInt4ToInt(int4Val pgtype.Int4) (*int, error) {
	if !int4Val.Valid {
		return nil, nil
	}

	result := int(int4Val.Int32)
	return &result, nil
}

// ConvertIntToInt4 converts int to pgtype.Int4.
func ConvertIntToInt4(i int) pgtype.Int4 {
	return pgtype.Int4{
		Int32: int32(i),
		Valid: true,
	}
}

// ConvertTextToString converts pgtype.Text to *string. Returns nil if Text is not valid.
func ConvertTextToString(textVal pgtype.Text) (*string, error) {
	if !textVal.Valid {
		return nil, nil
	}

	return &textVal.String, nil
}

// ConvertStringToText converts string to pgtype.Text.
func ConvertStringToText(str string) pgtype.Text {
	return pgtype.Text{
		String: str,
		Valid:  true,
	}
}

// ConvertTimestamptzToTime converts pgtype.Timestamptz to *time.Time. Returns nil if Timestamptz is not valid.
func ConvertTimestamptzToTime(tzVal pgtype.Timestamptz) (*time.Time, error) {
	if !tzVal.Valid {
		return nil, nil
	}

	return &tzVal.Time, nil
}

// ConvertTimeToTimestamptz converts time.Time to pgtype.Timestamptz.
func ConvertTimeToTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
