// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: batch.go

package sqlc

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrBatchAlreadyClosed = errors.New("batch already closed")
)

const addBusinessHours = `-- name: AddBusinessHours :batchexec
INSERT INTO business_hour(store_id, week_day, opening_time, closing_time, timezone)
VALUES ($1, $2, $3, $4, $5)
`

type AddBusinessHoursBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type AddBusinessHoursParams struct {
	StoreID     pgtype.UUID `db:"store_id" json:"store_id"`
	WeekDay     int32       `db:"week_day" json:"week_day"`
	OpeningTime pgtype.Time `db:"opening_time" json:"opening_time"`
	ClosingTime pgtype.Time `db:"closing_time" json:"closing_time"`
	Timezone    string      `db:"timezone" json:"timezone"`
}

// AddBusinessHours
//
//	INSERT INTO business_hour(store_id, week_day, opening_time, closing_time, timezone)
//	VALUES ($1, $2, $3, $4, $5)
func (q *Queries) AddBusinessHours(ctx context.Context, arg []AddBusinessHoursParams) *AddBusinessHoursBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.StoreID,
			a.WeekDay,
			a.OpeningTime,
			a.ClosingTime,
			a.Timezone,
		}
		batch.Queue(addBusinessHours, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &AddBusinessHoursBatchResults{br, len(arg), false}
}

func (b *AddBusinessHoursBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, ErrBatchAlreadyClosed)
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *AddBusinessHoursBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const deleteBusinessHours = `-- name: DeleteBusinessHours :batchexec
DELETE FROM business_hour
WHERE store_id = $1
  AND week_day = $2
  AND opening_time = $3
  AND closing_time = $4
`

type DeleteBusinessHoursBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type DeleteBusinessHoursParams struct {
	StoreID     pgtype.UUID `db:"store_id" json:"store_id"`
	WeekDay     int32       `db:"week_day" json:"week_day"`
	OpeningTime pgtype.Time `db:"opening_time" json:"opening_time"`
	ClosingTime pgtype.Time `db:"closing_time" json:"closing_time"`
}

// DeleteBusinessHours
//
//	DELETE FROM business_hour
//	WHERE store_id = $1
//	  AND week_day = $2
//	  AND opening_time = $3
//	  AND closing_time = $4
func (q *Queries) DeleteBusinessHours(ctx context.Context, arg []DeleteBusinessHoursParams) *DeleteBusinessHoursBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.StoreID,
			a.WeekDay,
			a.OpeningTime,
			a.ClosingTime,
		}
		batch.Queue(deleteBusinessHours, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &DeleteBusinessHoursBatchResults{br, len(arg), false}
}

func (b *DeleteBusinessHoursBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, ErrBatchAlreadyClosed)
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *DeleteBusinessHoursBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}
