package keycloak

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nerzal/gocloak/v13"

	"github.com/oprimogus/cardapiogo/internal/core/authentication"
	"github.com/oprimogus/cardapiogo/internal/core/user"
)

func entityUserToKeycloakUser(user user.User, userEnabled, userEmailVerified bool) *gocloak.User {
	realmRoles := make([]string, len(user.Roles))
	for i, v := range user.Roles {
		realmRoles[i] = string(v)
	}

	attributes := make(map[string][]string)
	if user.Profile.Document != "" {
		attributes["document"] = []string{user.Profile.Document}
	}

	return &gocloak.User{
		FirstName:     &user.Profile.Name,
		LastName:      &user.Profile.LastName,
		Username:      &user.Email,
		Enabled:       &userEnabled,
		EmailVerified: &userEmailVerified,
		Email:         &user.Email,
		RealmRoles:    &realmRoles,
		Attributes:    &attributes,
	}
}

func keycloakUserToEntityUser(userGocloak *gocloak.User) user.User {
	var document, phone string
	if userGocloak.Attributes != nil {
		if docValues, ok := (*userGocloak.Attributes)["document"]; ok && len(docValues) > 0 {
			document = docValues[0]
		}
		if phoneValues, ok := (*userGocloak.Attributes)["phone"]; ok && len(phoneValues) > 0 {
			phone = phoneValues[0]
		}
	}

	var userRoles []user.Role
	if userGocloak.RealmRoles != nil {
		userRoles = make([]user.Role, len(*userGocloak.RealmRoles))
		for i, v := range *userGocloak.RealmRoles {
			if user.IsValidRole(v) {
				userRoles[i] = user.Role(v)
			}
		}

	}

	return user.User{
		ID:    *userGocloak.ID,
		Email: *userGocloak.Email,
		Profile: user.Profile{
			Name:     *userGocloak.FirstName,
			LastName: *userGocloak.LastName,
			Document: document,
			Phone:    phone,
		},
		Roles: userRoles,
	}
}

func (k *KeycloakService) Create(ctx context.Context, user user.User) error {
	enabled := true
	emailVerified := false
	keycloakUser := entityUserToKeycloakUser(user, enabled, emailVerified)

	id, err := k.client.CreateUser(ctx, k.token.AccessToken, k.realm, *keycloakUser)
	if err != nil {
		return err
	}
	errSetPassword := k.client.SetPassword(ctx, k.token.AccessToken, id, k.realm, user.Password, false)
	if errSetPassword != nil {
		return errSetPassword
	}
	return nil
}

func (k *KeycloakService) FindByID(ctx context.Context, id string) (user.User, error) {
	userGocloak, err := k.client.GetUserByID(ctx, k.token.AccessToken, k.realm, id)
	if err != nil {
		return user.User{}, err
	}
	return keycloakUserToEntityUser(userGocloak), nil
}

func (k *KeycloakService) FindByEmail(ctx context.Context, email string) (user.User, error) {
	maxUsers := 1
	userGocloak, err := k.client.GetUsers(
		ctx,
		k.token.AccessToken,
		k.realm,
		gocloak.GetUsersParams{Email: &email, Max: &maxUsers},
	)
	if err != nil {
		return user.User{}, err
	}
	if len(userGocloak) == 0 {
		return user.User{}, nil
	}
	return keycloakUserToEntityUser(userGocloak[0]), nil
}

func (k *KeycloakService) Update(ctx context.Context, user user.User) error {
	actualUser, err := k.client.GetUserByID(ctx, k.token.AccessToken, k.realm, user.ID)
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
	errUpdateUser := k.client.UpdateUser(ctx, k.token.AccessToken, k.realm, *actualUser)
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

func (k *KeycloakService) SignIn(ctx context.Context, email, password string) (authentication.JWT, error) {
	jwtInstance, err := k.client.Login(ctx, k.clientID, k.clientSecret, k.realm, email, password)
	if err != nil {
		return authentication.JWT{}, err
	}
	return authentication.JWT{
		AccessToken:      jwtInstance.AccessToken,
		IDToken:          jwtInstance.IDToken,
		ExpiresIn:        jwtInstance.ExpiresIn,
		RefreshExpiresIn: jwtInstance.RefreshExpiresIn,
		RefreshToken:     jwtInstance.RefreshToken,
		TokenType:        jwtInstance.TokenType,
		NotBeforePolicy:  jwtInstance.NotBeforePolicy,
		SessionState:     jwtInstance.SessionState,
		Scope:            jwtInstance.Scope,
	}, nil
}

func (k *KeycloakService) RefreshToken(ctx context.Context, refreshToken string) (authentication.JWT, error) {
	jwtInstance, err := k.client.RefreshToken(ctx, refreshToken, k.clientID, k.clientSecret, k.realm)
	if err != nil {
		return authentication.JWT{}, err
	}
	return authentication.JWT{
		AccessToken:      jwtInstance.AccessToken,
		IDToken:          jwtInstance.IDToken,
		ExpiresIn:        jwtInstance.ExpiresIn,
		RefreshExpiresIn: jwtInstance.RefreshExpiresIn,
		RefreshToken:     jwtInstance.RefreshToken,
		TokenType:        jwtInstance.TokenType,
		NotBeforePolicy:  jwtInstance.NotBeforePolicy,
		SessionState:     jwtInstance.SessionState,
		Scope:            jwtInstance.Scope,
	}, nil
}

func (k *KeycloakService) IsValidToken(ctx context.Context, token string) (bool, error) {
	a, err := k.client.RetrospectToken(ctx, token, k.clientID, k.clientSecret, k.realm)
	if err != nil {
		return false, fmt.Errorf("ocurred an error while validate your access token: %w", err)
	}
	if a == nil {
		return false, errors.New("ocurred an error while validate your access token")
	}
	return *a.Active, nil
}

func (k *KeycloakService) DecodeAccessToken(ctx context.Context, accessToken string) (map[string]interface{}, error) {
	token, mapClaims, err := k.client.DecodeAccessToken(ctx, accessToken, k.realm)
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

func (k *KeycloakService) AddRoles(ctx context.Context, userID string, roles []user.Role) error {
	userKeycloakRoles, err := k.client.GetRealmRolesByUserID(ctx, k.token.AccessToken, k.realm, userID)
	if err != nil {
		return fmt.Errorf("fail in set new roles: %w", err)
	}

	allKeycloakRoles, err := k.client.GetRealmRoles(ctx, k.token.AccessToken, k.realm, gocloak.GetRoleParams{})
	if err != nil {
		return fmt.Errorf("fail in get all roles of realm: %w", err)
	}

	userKeycloakRolesMap := make(map[string]bool)
	for _, v := range userKeycloakRoles {
		userKeycloakRolesMap[*v.Name] = true
	}

	allKeycloakRolesMap := make(map[string]*gocloak.Role)
	for _, v := range allKeycloakRoles {
		allKeycloakRolesMap[*v.Name] = v
	}

	var rolesToAdd []gocloak.Role
	for _, role := range roles {
		if !userKeycloakRolesMap[string(role)] {
			if keycloakRole, exists := allKeycloakRolesMap[string(role)]; exists {
				rolesToAdd = append(rolesToAdd, *keycloakRole)
			}
		}
	}

	if len(rolesToAdd) == 0 {
		return nil
	}

	errSetRoles := k.client.AddRealmRoleToUser(ctx, k.token.AccessToken, k.realm, userID, rolesToAdd)
	if errSetRoles != nil {
		return fmt.Errorf("fail in set new roles for user: %w", errSetRoles)
	}
	return nil
}
