package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service struct
type Service struct {
	repository Repository
}

// NewService Service constructor
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// CreateUser create a user in database
func (u Service) CreateUser(ctx context.Context, newUser CreateUserParams) error {
	hashPassword, err := u.HashPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("Fail in generate hash hash of password: %s", err)
	}
	newUser.Password = hashPassword
	return u.repository.CreateUser(ctx, newUser)
}

// GetUser return a user from database
func (u Service) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	return u.repository.GetUserByID(ctx, id)
}

// GetUsersList return a user from database
func (u Service) GetUsersList(ctx context.Context, items int, page int) ([]User, error) {
	return u.repository.GetUsersList(ctx, items, page)
}

// HashPassword generate a hash of password for save in database
func (u Service) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash verify if password hash is valid
func (u Service) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
