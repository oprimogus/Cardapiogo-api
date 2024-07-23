package store

import (
	"context"
	"fmt"
	"time"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type AddBusinessHour struct {
	repository repository.StoreRepository
}

func NewUpdateBusinessHour(repository repository.StoreRepository) AddBusinessHour {
	return AddBusinessHour{repository: repository}
}

func (u AddBusinessHour) Execute(ctx context.Context, params StoreBusinessHoursParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}

	isOwner, errIsOwner := u.repository.IsOwner(ctx, params.ID, userID)
	if errIsOwner != nil {
		return fmt.Errorf("fail on get store data")
	}

	if !isOwner {
		return errNotOwner
	}

	businessHours := make([]entity.BusinessHours, len(params.BusinessHours))
	for i, v := range params.BusinessHours {
		zone, errLoadTimeZone := time.LoadLocation(params.TimeZone)
		if errLoadTimeZone != nil {
			return fmt.Errorf("fail on load timezone: %w", errLoadTimeZone)
		}

		openingTimeParsed, errOpeningTime := time.ParseInLocation(entity.BusinessHourLayout, v.OpeningTime, zone)
		if errOpeningTime != nil {
			return fmt.Errorf("fail in parse openingTime: %w", errOpeningTime)
		}
		openingTime := time.Date(1970, time.January, 1, openingTimeParsed.Hour(), openingTimeParsed.Minute(), openingTimeParsed.Second(), 0, zone)

		closingTimeParsed, errClosingTime := time.ParseInLocation(entity.BusinessHourLayout, v.ClosingTime, zone)
		if errClosingTime != nil {
			return fmt.Errorf("fail in parse closingTime: %w", errClosingTime)
		}
		closingTime := time.Date(1970, time.January, 1, closingTimeParsed.Hour(), closingTimeParsed.Minute(), closingTimeParsed.Second(), 0, zone)

		businessHours[i] = entity.BusinessHours{
			WeekDay:     int(v.WeekDay),
			OpeningTime: openingTime,
			ClosingTime: closingTime,
			TimeZone:    params.TimeZone,
		}

	}

	err := u.repository.AddBusinessHour(ctx, params.ID, businessHours)
	if err != nil {
		return err
	}
	return nil
}
