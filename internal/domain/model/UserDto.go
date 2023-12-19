package model

// CreateUserParams DTO
type CreateUserParams struct {
	Email           string          `db:"email" json:"email"`
	Password        string          `db:"password" json:"password"`
	Role            Role            `db:"role" json:"role"`
	AccountProvider AccountProvider `db:"account_provider" json:"account_provider"`
}

// UpdateUserParams DTO
type UpdateUserParams struct {
	ID    string `db:"id" json:"id"`
	Email string      `db:"email" json:"email"`
	Role  Role   `db:"role" json:"role"`
}

// UpdateUserPasswordParams DTO
type UpdateUserPasswordParams struct {
	ID       string `db:"id" json:"id"`
	Password string `db:"password" json:"password"`
}

// UpdateUserProfileParams DTO
type UpdateUserProfileParams struct {
	ID        string `db:"id" json:"id"`
	ProfileID int `db:"profile_id" json:"profile_id"`
}
