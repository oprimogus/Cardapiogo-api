package store

import (
	"context"
	"fmt"
	"time"
)

type useCaseAddBusinessHour struct {
	repository Repository
}

func newUseCaseAddBusinessHour(repository Repository) useCaseAddBusinessHour {
	return useCaseAddBusinessHour{repository: repository}
}

func (a useCaseAddBusinessHour) Execute(ctx context.Context, params StoreBusinessHoursParams) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return fmt.Errorf("invalid userID: '%s'", userID)
	}

	isOwner, errIsOwner := a.repository.IsOwner(ctx, params.ID, userID)
	if errIsOwner != nil {
		return fmt.Errorf("fail on get store data")
	}

	if !isOwner {
		return errNotOwner
	}

	businessHours := make([]BusinessHours, len(params.BusinessHours))
	for i, v := range params.BusinessHours {
		zone, errLoadTimeZone := time.LoadLocation(params.TimeZone)
		if errLoadTimeZone != nil {
			return fmt.Errorf("fail on load timezone: %w", errLoadTimeZone)
		}

		openingTimeParsed, errOpeningTime := time.ParseInLocation(BusinessHourLayout, v.OpeningTime, zone)
		if errOpeningTime != nil {
			return fmt.Errorf("fail in parse openingTime: %w", errOpeningTime)
		}
		openingTime := time.Date(1970, time.January, 1, openingTimeParsed.Hour(), openingTimeParsed.Minute(), openingTimeParsed.Second(), 0, zone)

		closingTimeParsed, errClosingTime := time.ParseInLocation(BusinessHourLayout, v.ClosingTime, zone)
		if errClosingTime != nil {
			return fmt.Errorf("fail in parse closingTime: %w", errClosingTime)
		}
		closingTime := time.Date(1970, time.January, 1, closingTimeParsed.Hour(), closingTimeParsed.Minute(), closingTimeParsed.Second(), 0, zone)

		businessHours[i] = BusinessHours{
			WeekDay:     int(v.WeekDay),
			OpeningTime: openingTime,
			ClosingTime: closingTime,
			TimeZone:    params.TimeZone,
		}

	}

	err := a.repository.AddBusinessHour(ctx, params.ID, businessHours)
	if err != nil {
		return err
	}
	return nil
}
