package keycloak

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nerzal/gocloak/v13"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
)

func entityUserToKeycloakUser(user entity.User, userEnabled, userEmailVerified bool) *gocloak.User {
	realmRoles := make([]string, len(user.Roles))
	for i, v := range user.Roles {
		realmRoles[i] = string(v)
	}

	return &gocloak.User{
		FirstName:     &user.Profile.Name,
		LastName:      &user.Profile.LastName,
		Username:      &user.Email,
		Enabled:       &userEnabled,
		EmailVerified: &userEmailVerified,
		Email:         &user.Email,
		RealmRoles:    &realmRoles,
		Attributes: &map[string][]string{
			"document": {user.Profile.Document},
		},
	}
}

func keycloakUserToEntityUser(user *gocloak.User) entity.User {
	var document, phone string
	if user.Attributes != nil {
		if docValues, ok := (*user.Attributes)["document"]; ok && len(docValues) > 0 {
			document = docValues[0]
		}
		if phoneValues, ok := (*user.Attributes)["phone"]; ok && len(phoneValues) > 0 {
			phone = phoneValues[0]
		}
	}

	userRoles := make([]entity.UserRole, len(*user.RealmRoles))
	for i, v := range *user.RealmRoles {
		if entity.IsValidUserRole(v) {
			userRoles[i] = entity.UserRole(v)
		}
	}

	return entity.User{
		ID: *user.ID,
		Email:      *user.Email,
		Profile: entity.Profile{
			Name:     *user.FirstName,
			LastName: *user.LastName,
			Document: document,
			Phone:    phone,
		},
		Roles: userRoles,
	}
}

func (k *KeycloakService) Create(ctx context.Context, user entity.User) error {
	enabled := true
	emailVerified := false
	keycloakUser := entityUserToKeycloakUser(user, enabled, emailVerified)

	id, err := k.client.CreateUser(ctx, k.token.AccessToken, realm, *keycloakUser)
	if err != nil {
		return err
	}
	errSetPassword := k.client.SetPassword(ctx, k.token.AccessToken, id, realm, user.Password, false)
	if errSetPassword != nil {
		return errSetPassword
	}
	return nil
}

func (k *KeycloakService) FindByID(ctx context.Context, id string) (entity.User, error) {
	user, err := k.client.GetUserByID(ctx, k.token.AccessToken, realm, id)
	if err != nil {
		return entity.User{}, err
	}
	return keycloakUserToEntityUser(user), nil
}

func (k *KeycloakService) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	maxUsers := 1
	user, err := k.client.GetUsers(
		ctx,
		k.token.AccessToken,
		realm,
		gocloak.GetUsersParams{Email: &email, Max: &maxUsers},
	)
	if err != nil {
		return entity.User{}, err
	}
	if len(user) == 0 {
		return entity.User{}, nil
	}
	return keycloakUserToEntityUser(user[0]), nil
}

func (k *KeycloakService) Update(ctx context.Context, user entity.User) error {
	actualUser, err := k.client.GetUserByID(ctx, k.token.AccessToken, realm, user.ID)
	if err != nil {
		return fmt.Errorf("fail in found user: %w", err)
	}
	actualUser.FirstName = &user.Profile.Name
	actualUser.LastName = &user.Profile.LastName
	if actualUser.Attributes != nil {
		if phoneValues, ok := (*actualUser.Attributes)["phone"]; ok && len(phoneValues) > 0 {
			phoneValues[0] = user.Profile.Phone
		}
	}
	errUpdateUser := k.client.UpdateUser(ctx, k.token.AccessToken, realm, *actualUser) 
	if errUpdateUser != nil {
		return fmt.Errorf("fail in update user data: %w", errUpdateUser)
	}
	return nil
}

func (k *KeycloakService) Delete(ctx context.Context, id string) error {
	return k.client.DeleteUser(ctx, k.token.AccessToken, k.realm, id)
}

func (k *KeycloakService) ResetPasswordByEmail(ctx context.Context, id string) error {
	lifespan := 30 * 60
	actions := []string{"UPDATE_PASSWORD"}
	return k.client.ExecuteActionsEmail(
		ctx,
		k.token.AccessToken,
		k.realm,
		gocloak.ExecuteActionsEmail{
			Lifespan: &lifespan,
			ClientID: &k.clientID,
			Actions:  &actions,
		},
	)
}

func (k *KeycloakService) SignIn(ctx context.Context, email, password string) (entity.JWT, error) {
	jwt, err := k.client.Login(ctx, clientID, clientSecret, realm, email, password)
	if err != nil {
		return entity.JWT{}, err
	}
	return entity.JWT{
		AccessToken:      jwt.AccessToken,
		IDToken:          jwt.IDToken,
		ExpiresIn:        jwt.ExpiresIn,
		RefreshExpiresIn: jwt.RefreshExpiresIn,
		RefreshToken:     jwt.RefreshToken,
		TokenType:        jwt.TokenType,
		NotBeforePolicy:  jwt.NotBeforePolicy,
		SessionState:     jwt.SessionState,
		Scope:            jwt.Scope,
	}, nil
}

func (k *KeycloakService) IsValidToken(ctx context.Context, token string) (bool, error) {
	a, err := k.client.RetrospectToken(ctx, token, clientID, clientSecret, realm)
	if err != nil {
		return false, fmt.Errorf("ocurred an error while validate your access token: %w", err)
	}
	if a == nil {
		return false, errors.New("ocurred an error while validate your access token")
	}
	return *a.Active, nil
}

func (k *KeycloakService) DecodeAccessToken(ctx context.Context, accessToken string) (map[string]interface{}, error) {
	token, mapClaims, err := k.client.DecodeAccessToken(ctx, accessToken, realm)
	if err != nil {
		return nil, fmt.Errorf("unable to decode access token: %w", err)
	}
	if mapClaims == nil {
		return nil, fmt.Errorf("unable to decode access token and get claims")
	}
	if token == nil {
		return nil, fmt.Errorf("unable to decode access token and get metadata")
	}

	return *mapClaims, nil
}
