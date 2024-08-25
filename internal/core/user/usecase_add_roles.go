package user

import (
	"context"
	"fmt"
)

type useCaseAddRoles struct {
	repository Repository
}

func newUseCaseAddRoles(repository Repository) useCaseAddRoles {
	return useCaseAddRoles{repository: repository}
}

func (a useCaseAddRoles) Execute(ctx context.Context, roles []Role) error {
	userID := ctx.Value("userID").(string)
	err := a.repository.AddRoles(ctx, userID, roles)
	if err != nil {
		return fmt.Errorf("fail on add roles to user: %w", err)
	}
	return nil
}
