package authentication

type AuthenticationModule struct {
	repository Repository
	SignIn useCaseSignIn
	Refresh useCaseRefresh
}

func NewAuthenticationModule(repository Repository) AuthenticationModule {
	return AuthenticationModule{
		repository: repository,
		SignIn: newUseCaseSignIn(repository),
		Refresh: newUseCaseRefresh(repository),
	}
}