package user

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

// Service struct
type Service struct {
	repository Repository
}

// NewService Service constructor
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

// CreateUser create a user in database
func (u *Service) CreateUser(ctx context.Context, newUser CreateUserParams) error {
	hashPassword, err := u.HashPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("fail in generate hash of password: %s", err)
	}
	newUser.Password = hashPassword
	return u.repository.CreateUser(ctx, newUser)
}

// GetUser return a user from database
func (u *Service) GetUser(ctx context.Context, id string) (User, error) {
	return u.repository.GetUserByID(ctx, id)
}

// GetUsersList return a user from database
func (u *Service) GetUsersList(ctx context.Context, items int, page int) ([]User, error) {
	return u.repository.GetUsersList(ctx, items, page)
}

// UpdateUserPassword change the password of user
func (u *Service) UpdateUserPassword(ctx context.Context, params UpdateUserPasswordParams) error {
	user, err := u.GetUser(ctx, params.ID)
	if err != nil {
		return err
	}

	if !u.IsValidPassword(params.Password, user.Password) {
		return errors.NewErrorResponse(http.StatusBadRequest, "Invalid Password")
	}

	params.NewPassword, err = u.HashPassword(params.NewPassword)
	if err != nil {
		return err
	}
	return u.repository.UpdateUserPassword(ctx, params)
}

// UpdateUser change the user data
func (u *Service) UpdateUser(ctx context.Context, params UpdateUserParams) error {
	user, err := u.GetUser(ctx, params.ID)
	if err != nil {
		return err
	}

	if !u.IsValidPassword(params.Password, user.Password) {
		return errors.NewErrorResponse(http.StatusBadRequest, "Invalid Password")
	}

	return u.repository.UpdateUser(ctx, params)
}

// HashPassword generate a hash of password for save in database
func (u *Service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash verify if password hash is valid and 
func (u *Service) IsValidPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
