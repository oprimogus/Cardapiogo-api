package user_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
	mock_user "github.com/oprimogus/cardapiogo/internal/infra/mocks/user"
)

const (
	ID       = "13e4c848-ee9d-4673-a2fc-a3620cec94f0"
	EMAIL    = "johndoe@example.com"
	PASSWORD = "teste123"
	LOCAL    = types.UserRoleConsumer
)

type ServiceSuite struct {
	suite.Suite
	controller               *gomock.Controller
	Repository               *mock_user.MockRepository
	Service                  *user.Service
	createUserParams         user.CreateUserParams
	updateUserPasswordParams user.UpdateUserPasswordParams
	updateUserParams         user.UpdateUserParams
	user                     user.User
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) ServiceSuiteDown(t *testing.T) {
	defer s.controller.Finish()
}

func (s *ServiceSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.Repository = mock_user.NewMockRepository(s.controller)
	s.Service = user.NewService(s.Repository)
	s.createUserParams = user.CreateUserParams{
		Email:           EMAIL,
		Password:        EMAIL,
		Role:            types.UserRoleConsumer,
		AccountProvider: types.AccountProviderGoogle,
	}
	s.updateUserPasswordParams = user.UpdateUserPasswordParams{
		ID:          ID,
		Password:    PASSWORD,
		NewPassword: PASSWORD,
	}
	s.updateUserParams = user.UpdateUserParams{
		ID:       ID,
		Email:    EMAIL,
		Password: PASSWORD,
		Role:     types.UserRoleConsumer,
	}
	s.user = user.User{
		ID:              ID,
		ProfileID:       0,
		Email:           EMAIL,
		Password:        "$2a$14$QH8.9CMBXx3T8pOQVhNsM.Fz2yU1/tNTsQ3Iq1htWX/XqsLSGayQq", // teste123
		Role:            types.UserRoleConsumer,
		AccountProvider: types.AccountProviderLocal,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// TestCreateUserShouldReturnNil should return success to create a new user
func (s *ServiceSuite) TestCreateUserShouldReturnNil() {
	s.Repository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	err := s.Service.CreateUser(context.Background(), s.createUserParams)
	assert.Nil(s.T(), err)
}

// TestCreateUserShouldReturnError should return an error in hash
func (s *ServiceSuite) TestCreateUserShouldReturnError() {
	s.Repository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(fmt.Errorf("fail in generate hash of password: banana")).Times(1)
	err := s.Service.CreateUser(context.Background(), s.createUserParams)
	assert.NotNil(s.T(), err)
}

// TestUpdateUserPasswordShouldReturnNil should return success to change user password
func (s *ServiceSuite) TestUpdateUserPasswordShouldReturnNil() {
	s.Repository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(s.user, nil).Times(1)
	s.Repository.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	err := s.Service.UpdateUserPassword(context.Background(), s.updateUserPasswordParams)
	assert.Nil(s.T(), err)
}

// TestUpdateUserPasswordShouldReturnErrorInGetUser should return error when not found a exist user
func (s *ServiceSuite) TestUpdateUserPasswordShouldReturnErrorInGetUser() {
	s.Repository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(user.User{}, errors.New(http.StatusNotFound, errors.NOT_FOUND_RECORD)).Times(1)
	err := s.Service.UpdateUserPassword(context.Background(), s.updateUserPasswordParams)
	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, errors.NOT_FOUND_RECORD)
}

// TestUpdateUserPasswordShouldReturnError should return error when input password is not equal to user password
func (s *ServiceSuite) TestUpdateUserPasswordShouldReturnError() {
	s.Repository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(s.user, nil).Times(1)
	s.updateUserPasswordParams.Password = "teste1234"
	err := s.Service.UpdateUserPassword(context.Background(), s.updateUserPasswordParams)
	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, user.INVALID_PASSWORD)
}

// TestUpdateUserShouldReturnNil should return success to change user data
func (s *ServiceSuite) TestUpdateUserShouldReturnNil() {
	s.Repository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(s.user, nil).Times(1)
	s.Repository.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	err := s.Service.UpdateUser(context.Background(), s.updateUserParams)
	assert.Nil(s.T(), err)
}

// TestUpdateUserShouldReturnErrorInGetUser should return error when not found a exist user
func (s *ServiceSuite) TestUpdateUserShouldReturnErrorInGetUser() {
	s.Repository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(user.User{}, errors.New(http.StatusNotFound, errors.NOT_FOUND_RECORD)).Times(1)
	err := s.Service.UpdateUser(context.Background(), s.updateUserParams)
	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, errors.NOT_FOUND_RECORD)
}

// TestUpdateUserShouldReturnError should return error when input password is not equal to user password
func (s *ServiceSuite) TestUpdateUserShouldReturnError() {
	s.Repository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(s.user, nil).Times(1)
	s.updateUserParams.Password = "teste1234"
	err := s.Service.UpdateUser(context.Background(), s.updateUserParams)
	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, user.INVALID_PASSWORD)
}

func (s *ServiceSuite) TestIsValidPasswordShouldReturnTrue() {
	isValidPassword := s.Service.IsValidPassword("teste123", s.user.Password)
	assert.True(s.T(), isValidPassword)
}

func (s *ServiceSuite) TestIsValidPasswordShouldReturnFalse() {
	isValidPassword := s.Service.IsValidPassword("teste1234", s.user.Password)
	assert.False(s.T(), isValidPassword)
}

// TestHashPassword should validate the creation of hashs
func (s *ServiceSuite) TestHashPassword() {
	pass := "teste123"
	anotherPass := "teste124"
	hashPassword, err := s.Service.HashPassword(pass)

	assert.Nil(s.T(), err)
	assert.True(s.T(), s.Service.IsValidPassword(pass, hashPassword))
	assert.False(s.T(), s.Service.IsValidPassword(anotherPass, hashPassword))
}
