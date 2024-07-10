package user

import (
	"context"
	"fmt"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
)

type AddRoles interface {
	Execute(ctx context.Context, roles []entity.UserRole) error
}

type addRoles struct {
	repository repository.UserRepository
}

func NewAddRoles(repository repository.UserRepository) addRoles {
	return addRoles{repository: repository}
}

func (a addRoles) Execute(ctx context.Context, roles []entity.UserRole) error {
	userID := ctx.Value("userID").(string)
	err := a.repository.AddRoles(ctx, userID, roles)
	if err != nil {
		return fmt.Errorf("fail on add roles to user: %w", err)
	}
	return nil
}
