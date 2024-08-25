package user

import (
	"context"
	"errors"
)

type useCaseResetPassword struct {
	repository Repository
}

func newResetPassword(repository Repository) useCaseResetPassword {
	return useCaseResetPassword{
		repository: repository,
	}
}

func (r useCaseResetPassword) Execute(ctx context.Context, id string) error {
	userID := ctx.Value("userID").(string)
	if userID == "" {
		return errors.New("invalid user")
	}
	return r.repository.ResetPasswordByEmail(ctx, id)
}
