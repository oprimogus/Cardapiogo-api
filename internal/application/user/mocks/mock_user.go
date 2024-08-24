package mock_user

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByID(ctx context.Context, id string) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(ctx context.Context, user entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepositoryMock) AddRoles(ctx context.Context, userID string, roles []entity.UserRole) error {
	args := m.Called(ctx, userID, roles)
	return args.Error(0)
}

func (m *UserRepositoryMock) ResetPasswordByEmail(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *UserRepositoryMock) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

