package authentication

type SignInParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshParams struct {
	RefreshToken    string `json:"refreshToken" validate:"required"`
}
