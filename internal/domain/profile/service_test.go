package profile_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oprimogus/cardapiogo/internal/domain/profile"
	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	mock_profile "github.com/oprimogus/cardapiogo/internal/infra/mocks/profile"
	mock_user "github.com/oprimogus/cardapiogo/internal/infra/mocks/user"
)

type ServiceSuite struct {
	suite.Suite
	controller          *gomock.Controller
	repository          *mock_profile.MockRepository
	userRepository      *mock_user.MockRepository
	Service             *profile.Service
	user                user.User
	createProfileParams profile.CreateProfileParams
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) ServiceSuiteDown(t *testing.T) {
	defer s.controller.Finish()
}

func (s *ServiceSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.repository = mock_profile.NewMockRepository(s.controller)
	s.userRepository = mock_user.NewMockRepository(s.controller)
	s.Service = profile.NewService(s.repository, s.userRepository)
	s.user = user.User{
		ID:              "13e4c848-ee9d-4673-a2fc-a3620cec94f0",
		ProfileID:       0,
		Email:           "johndoe@example.com",
		Password:        "$2a$14$QH8.9CMBXx3T8pOQVhNsM.Fz2yU1/tNTsQ3Iq1htWX/XqsLSGayQq", // teste123
		Role:            types.UserRoleConsumer,
		AccountProvider: types.AccountProviderLocal,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	s.createProfileParams = profile.CreateProfileParams{
		Name:     "John",
		LastName: "Marston",
		CPF:      "69950517850",
		Phone:    "13997590578",
	}
}

// TestCreateProfileShouldReturnSuccess must create a profile and associate this profile to user
func (s *ServiceSuite) TestCreateProfileShouldReturnSuccess() {
	s.userRepository.EXPECT().GetUserByID(gomock.Any(), s.user.ID).Return(s.user, nil).Times(1)
	s.repository.EXPECT().CreateProfile(gomock.Any(), s.user.ID, s.createProfileParams).Return(nil).Times(1)
	err := s.Service.CreateProfile(context.Background(), s.user.ID, s.createProfileParams)
	assert.Nil(s.T(), err)
}

// TestCreateProfileShouldReturnConflictError must return an error because the user already has an associated profile.
func (s *ServiceSuite) TestCreateProfileShouldReturnConflictError() {
	s.user.ProfileID = 63735689
	s.userRepository.EXPECT().GetUserByID(gomock.Any(), s.user.ID).Return(s.user, nil).Times(1)
	err := s.Service.CreateProfile(context.Background(), s.user.ID, s.createProfileParams)
	assert.EqualError(s.T(), err, profile.ExistRegisteredProfile)
}

// TestCreateProfileShouldReturnUserNotFoundError must return an error because the user already has an associated profile.
func (s *ServiceSuite) TestCreateProfileShouldReturnUserNotFoundError() {
	s.userRepository.EXPECT().GetUserByID(gomock.Any(), s.user.ID).Return(user.User{}, errors.New(http.StatusNotFound, errors.NOT_FOUND_RECORD)).Times(1)
	err := s.Service.CreateProfile(context.Background(), s.user.ID, s.createProfileParams)
	assert.EqualError(s.T(), err, errors.NOT_FOUND_RECORD)
}
