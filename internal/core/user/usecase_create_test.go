package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/oprimogus/cardapiogo/internal/core/user"
)

type CreateSuite struct {
	suite.Suite
	userRepository *mockRepository
	userModule     user.UserModule
}

func (s *CreateSuite) SetupSuite() {
	s.userRepository = new(mockRepository)
	s.userModule = user.NewUserModule(s.userRepository)
}

func (s *CreateSuite) TearDownTest() {
	s.userRepository.AssertExpectations(s.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateSuite))
}

func (s *CreateSuite) TestExecute() {
	input := user.CreateParams{
		Email:    "johndoe@example.com",
		Password: "teste123",
		Profile: user.CreateProfileParams{
			Name:     "John",
			LastName: "Doe",
			Phone:    "13981142501"},
	}
	tests := []struct {
		name                 string
		input                user.CreateParams
		mockFindByEmailValue user.User
		mockFindByEmailError error
		mockFindByEmailTimes int
		mockCreateError      error
		mockCreateErrorTimes int
		expected             error
	}{
		{
			name:  "Exist user with this email",
			input: input,
			mockFindByEmailValue: user.User{
				ID: "randowm_uuid",
				Profile: user.Profile{
					Name:     input.Profile.Name,
					LastName: input.Profile.LastName,
					Phone:    input.Profile.Phone,
				},
				Email:     input.Email,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			mockFindByEmailError: nil,
			mockFindByEmailTimes: 1,
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             user.ErrExistUserWithEmail,
		},
		{
			name:                 "Fail on find user on repository",
			input:                input,
			mockFindByEmailValue: user.User{},
			mockFindByEmailError: errors.New("Repository error"),
			mockFindByEmailTimes: 1,
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             errors.New("Repository error"),
		},
		{
			name:                 "Fail in create user on repository",
			input:                input,
			mockFindByEmailValue: user.User{},
			mockFindByEmailError: nil,
			mockFindByEmailTimes: 1,
			mockCreateError:      errors.New("Repository error"),
			mockCreateErrorTimes: 1,
			expected:             errors.New("Repository error"),
		},
	}

	for _, v := range tests {
		if v.mockFindByEmailTimes > 0 {
			s.userRepository.On("FindByEmail", mock.Anything, v.input.Email).
				Return(v.mockFindByEmailValue, v.mockFindByEmailError).
				Times(v.mockFindByEmailTimes)
		}
		if v.mockCreateErrorTimes > 0 {
			s.userRepository.On("Create", mock.Anything, mock.Anything).
				Return(v.mockCreateError).
				Times(v.mockCreateErrorTimes)
		}
		err := s.userModule.Create.Execute(context.Background(), v.input)
		assert.Equal(s.T(), err, v.expected, v.name)
		s.userRepository.AssertExpectations(s.T())
	}
}
