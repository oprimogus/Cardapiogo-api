package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oprimogus/cardapiogo/internal/application/user"
	mock_user_repository "github.com/oprimogus/cardapiogo/internal/application/user/mocks"
	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

type CreateSuite struct {
	suite.Suite
	controller     *gomock.Controller
	userRepository *mock_user_repository.MockUserRepository
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
	s.userRepository = mock_user_repository.NewMockUserRepository(s.controller)
	s.create = user.NewCreate(s.userRepository)
}

func (s *CreateSuite) TestExecute() {
	input := user.CreateParams{
		Email:    "johndoe@example.com",
		Password: "teste123",
		Profile:  user.CreateProfileParams{Name: "John", LastName: "Doe", Document: "50338097949", Phone: "13981142501"},
	}

	tests := []struct {
		name                 string
		input                user.CreateParams
		mockFindByEmailValue entity.User
		mockFindByEmailError error
		mockCreateError      error
		mockCreateErrorTimes int
		expected             error
	}{
		{
			name:                 "Not exist user with this email and document",
			input:                input,
			mockFindByEmailValue: entity.User{},
			mockFindByEmailError: nil,
			mockCreateError:      nil,
			mockCreateErrorTimes: 1,
			expected:             nil,
		},
		{
			name:  "Exist user with this email",
			input: input,
			mockFindByEmailValue: entity.User{
				ID: "randowm_uuid",
				Profile: entity.Profile{
					Name:     input.Profile.Name,
					LastName: input.Profile.LastName,
					Document: input.Profile.Document,
					Phone:    input.Profile.Phone,
				},
				Email:     input.Email,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			mockFindByEmailError: nil,
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             user.ErrExistUserWithEmail,
		},
		{
			name:  "Exist user with this document",
			input: input,
			mockFindByEmailValue: entity.User{
				ID: "randowm_uuid",
				Profile: entity.Profile{
					Name:     input.Profile.Name,
					LastName: input.Profile.LastName,
					Document: input.Profile.Document,
					Phone:    input.Profile.Phone,
				},
				Email:     "johndoe2@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			mockFindByEmailError: nil,
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             user.ErrExistUserWithDocument,
		},
		{
			name:  "Exist user with this phone",
			input: input,
			mockFindByEmailValue: entity.User{
				ID: "randowm_uuid",
				Profile: entity.Profile{
					Name:     input.Profile.Name,
					LastName: input.Profile.LastName,
					Document: input.Profile.Phone,
					Phone:    input.Profile.Phone,
				},
				Email:     "johndoe2@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			mockFindByEmailError: nil,
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             user.ErrExistUserWithPhone,
		},
		{
			name:                 "Fail on find user on repository",
			input:                input,
			mockFindByEmailValue: entity.User{},
			mockFindByEmailError: errors.New("Repository error"),
			mockCreateError:      nil,
			mockCreateErrorTimes: 0,
			expected:             errors.New("Repository error"),
		},
		{
			name:                 "Fail in create user on repository",
			input:                input,
			mockFindByEmailValue: entity.User{},
			mockFindByEmailError: nil,
			mockCreateError:      errors.New("Repository error"),
			mockCreateErrorTimes: 1,
			expected:             errors.New("Repository error"),
		},
	}

	for _, v := range tests {
		s.userRepository.EXPECT().FindByEmail(gomock.Any(), v.input.Email).Return(v.mockFindByEmailValue, v.mockFindByEmailError)
		s.userRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(v.mockCreateError).Times(v.mockCreateErrorTimes)
		err := s.create.Execute(context.Background(), v.input)
		assert.Equal(s.T(), err, v.expected, v.name)
	}
}
