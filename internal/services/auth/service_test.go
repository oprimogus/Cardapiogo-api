package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	mock_user "github.com/oprimogus/cardapiogo/internal/infra/mocks/user"
	"github.com/oprimogus/cardapiogo/internal/services/auth"
)

type ServiceSuite struct {
	suite.Suite
	controller     *gomock.Controller
	userRepository *mock_user.MockRepository
	userService    *user.Service
	user           user.User
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) ServiceSuiteDown(t *testing.T) {
	defer s.controller.Finish()
}

func (s *ServiceSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.userRepository = mock_user.NewMockRepository(s.controller)
	s.userService = user.NewService(s.userRepository)
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
}

func (s *ServiceSuite) TestGenerateJWTForValidationShouldReturnValidJWT() {
	jwt, errJWT := auth.GenerateJWTForValidation()
	isValidJWT, err := auth.ValidateStateToken(jwt)
	assert.Nil(s.T(), errJWT)
	assert.Nil(s.T(), err)
	assert.True(s.T(), isValidJWT)
}

func (s *ServiceSuite) TestGenerateJWTWithClaimsShouldReturnValidJWT() {
	jwt, errJWT := auth.GenerateJWTWithClaims(s.user)
	isValidJWT, err := auth.ValidateStateToken(jwt)
	assert.Nil(s.T(), errJWT)
	assert.Nil(s.T(), err)
	assert.True(s.T(), isValidJWT)
}

func (s *ServiceSuite) TestValidateStateTokenShouldValidateJWT() {
	jwt, errJWT := auth.GenerateJWTForValidation()
	isValid, err := auth.ValidateStateToken(jwt)
	assert.Nil(s.T(), errJWT)
	assert.Nil(s.T(), err)
	assert.True(s.T(), isValid)
}

func (s *ServiceSuite) TestLoginShouldReturnJWTOnValidCredentials() {
	s.userRepository.EXPECT().GetUserByEmail(gomock.Any(), s.user.Email).Return(s.user, nil)
	loginParams := &user.Login{Email: s.user.Email, Password: "teste123"}
	jwt, errJWT := auth.Login(context.Background(), s.userService, loginParams)
	isValid, err := auth.ValidateStateToken(jwt)
	assert.Nil(s.T(), errJWT)
	assert.NotEmpty(s.T(), jwt)
	assert.Nil(s.T(), err)
	assert.True(s.T(), isValid)
}

func (s *ServiceSuite) TestLoginShouldReturnErrorWithInvalidCredentials() {
	s.userRepository.EXPECT().GetUserByEmail(gomock.Any(), s.user.Email).Return(s.user, nil)
	loginParams := &user.Login{Email: s.user.Email, Password: "teste12345"}
	jwt, errJWT := auth.Login(context.Background(), s.userService, loginParams)
	assert.EqualError(s.T(), errJWT, user.INVALID_PASSWORD)
	assert.Empty(s.T(), jwt)
}
