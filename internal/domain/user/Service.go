package user

import (
	"context"

	"github.com/google/uuid"
)

// Service struct
type Service struct {
	repository Repository
}

// NewService Service constructor
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// CreateUser create a user in database
func (u Service) CreateUser(ctx context.Context, newUser CreateUserParams) error {
	return u.repository.CreateUser(ctx, newUser)
}

func (u Service) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	return 
}
