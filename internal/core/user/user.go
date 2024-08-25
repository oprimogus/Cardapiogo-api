package user

type UserModule struct {
	repository  Repository
	Create      useCaseCreate
	Update      useCaseUpdate
	Delete      useCaseDelete
	FindByID    useCaseFindByID
	FindByEmail useCaseFindByEmail
	AddRoles    useCaseAddRoles
}

func NewUserModule(userRepository Repository) UserModule {
	return UserModule{
		repository:  userRepository,
		Create:      newUseCaseCreate(userRepository),
		Update:      newUseCaseUpdate(userRepository),
		Delete:      newUseCaseDelete(userRepository),
		FindByID:    newUseCaseFindByID(userRepository),
		FindByEmail: newUseCaseFindByEmail(userRepository),
		AddRoles:    newUseCaseAddRoles(userRepository),
	}
}
