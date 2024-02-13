package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockRepository é um mock do Repository.
type MockRepository struct {
	mock.Mock
}

// Aqui declaramos todos os métodos que queremos mockar.
func (m *MockRepository) CreateUser(ctx context.Context, user CreateUserParams) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) CreateUserWithOAuth(ctx context.Context, user CreateUserWithOAuthParams) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) GetUserByID(ctx context.Context, id string) (User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockRepository) GetUsersList(ctx context.Context, items int, page int) ([]*User, error) {
	args := m.Called(ctx, items, page)
	return args.Get(0).([]*User), args.Error(1)
}

func (m *MockRepository) UpdateUser(ctx context.Context, user UpdateUserParams) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) UpdateUserPassword(ctx context.Context, user UpdateUserPasswordParams) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// TestService test user.Service
func TestService(t *testing.T) {

	mockRepo := new(MockRepository)
	service := NewService(mockRepo)
	ctx := context.Background()
	newUser := CreateUserParams{
		Email:           "johndoe@example.com",
		Password:        "johndoe@example.com",
		Role:            "Owner",
		AccountProvider: "Google",
	}

	t.Run("Create a new user", func(t *testing.T) {
		t.Run("Should return success", func(t *testing.T) {
			mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("CreateUserParams")).Return(nil)

			err := service.CreateUser(ctx, newUser)

			mockRepo.AssertExpectations(t)
			if err != nil {
				t.Fatalf("CreateUser falhou: %v", err)
			}
		})
	})
}
