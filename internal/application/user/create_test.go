package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oprimogus/cardapiogo/internal/application/user"
	mock_user "github.com/oprimogus/cardapiogo/internal/application/user/mocks"
	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

type CreateSuite struct {
	suite.Suite
	controller     *gomock.Controller
	userRepository *mock_user.UserRepositoryMock
	create         user.Create
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateSuite))
}

func (s *CreateSuite) SuitDown() {
	defer s.controller.Finish()
}

func (s *CreateSuite) SetupSuite() {
	s.controller = gomock.NewController(s.T())
	s.userRepository = &mock_user.UserRepositoryMock{}
	s.create = user.NewCreate(s.userRepository)
}

func (s *CreateSuite) TestExecute() {
	input := user.CreateParams{
		Email:    "johndoe@example.com",
		Password: "teste123",
		Profile:  user.CreateProfileParams{
			Name: "John", 
			LastName: "Doe", 
			Phone: "13981142501"},
	}
	tests := []struct {
		name                 string
		input                user.CreateParams
		mockFindByEmailValue entity.User
		mockFindByEmailError error
		mockFindByEmailTimes int
		mockCreateError      error
		mockCreateErrorTimes int
		expected             error
	}{
		{
			name:  "Exist user with this email",
			input: input,
			mockFindByEmailValue: entity.User{
				ID: "randowm_uuid",
				Profile: entity.Profile{
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
			expected:             entity.ErrExistUserWithEmail,
		},
		{
			name:                 "Fail on find user on repository",
			input:                input,
			mockFindByEmailValue: entity.User{},
			mockFindByEmailError: errors.New("Repository error"),
			mockFindByEmailTimes: 1,
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             errors.New("Repository error"),
		},
		{
			name:                 "Fail in create user on repository",
			input:                input,
			mockFindByEmailValue: entity.User{},
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
		err := s.create.Execute(context.Background(), v.input)
		assert.Equal(s.T(), err, v.expected, v.name)
		s.userRepository.AssertExpectations(s.T())
	}
}
