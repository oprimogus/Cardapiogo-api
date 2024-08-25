package user_test

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/oprimogus/cardapiogo/internal/core/user"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Create(ctx context.Context, user user.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockRepository) FindByID(ctx context.Context, id string) (user.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *mockRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *mockRepository) Update(ctx context.Context, user user.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockRepository) AddRoles(ctx context.Context, userID string, roles []user.Role) error {
	args := m.Called(ctx, userID, roles)
	return args.Error(0)
}

func (m *mockRepository) ResetPasswordByEmail(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
