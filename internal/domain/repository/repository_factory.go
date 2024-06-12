package repository

type Factory interface {
	NewUserRepository() UserRepository
}
