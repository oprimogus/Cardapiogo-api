package keycloak

type KeycloakUser struct {
	Username    string            `json:"username"`
	Enabled     bool              `json:"enabled"`
	FirstName   string            `json:"firstName,omitempty"`
	LastName    string            `json:"lastName,omitempty"`
	Email       string            `json:"email,omitempty"`
	Credentials []UserCredentials `json:"credentials,omitempty"`
}

type UserCredentials struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}
