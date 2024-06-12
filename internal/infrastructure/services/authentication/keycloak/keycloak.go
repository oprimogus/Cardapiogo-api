package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"github.com/oprimogus/cardapiogo/internal/domain/entity"
	"github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

var (
	keycloakURL      = os.Getenv("KEYCLOAK_BASE_URL")
	keycloakRealm    = os.Getenv("KEYCLOAK_REALM")
	keycloakUser     = os.Getenv("KEYCLOAK_ADMIN_USER")
	keycloakPassword = os.Getenv("KEYCLOAK_ADMIN_PASSWORD")
	clientID         = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret     = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	issuer           = os.Getenv("KEYCLOAK_BASE_URL") + "/realms/cardapiogo"
	redirectURL      = os.Getenv("KEYCLOAK_REDIRECT_URL")
)

const TimeExpireInHour = 1

type KeycloakService struct {
	provider *oidc.Provider
	Config   oauth2.Config
}

func NewKeycloakService(c context.Context) (*KeycloakService, error) {
	provider, err := oidc.NewProvider(c, issuer)
	if err != nil {
		return nil, err
	}

	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}
	return &KeycloakService{
		provider: provider,
		Config:   oauth2Config,
	}, nil
}

func GenerateJWTForValidation() (string, error) {
	key := os.Getenv("JWT_SECRET")
	expireIn := time.Now().Add(time.Hour * time.Duration(TimeExpireInHour)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": os.Getenv("JWT_EMISSOR"),
		"exp": expireIn,
	})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return s, err
}

func ValidateStateToken(stateToken string) (bool, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(stateToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.BadRequest(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func (k *KeycloakService) SignUp(c context.Context, username string, password string) (*oauth2.Token, error) {
	oauth2Token, err := k.Config.PasswordCredentialsToken(c, username, password)
	if err != nil {
		return nil, err
	}
	return oauth2Token, nil
}

func (k *KeycloakService) getAdminToken(c context.Context) (*oauth2.Token, error) {
	token, err := k.Config.PasswordCredentialsToken(c, keycloakUser, keycloakPassword)
	if err != nil {
		return nil, err
	}
	return token, err
}

func (k *KeycloakService) Create(ctx context.Context, params KeycloakUser) error {
	adminURL := fmt.Sprintf("%s/admin/realms/%s/users", keycloakURL, keycloakRealm)
	adminToken, err := k.getAdminToken(ctx)
	if err != nil {
		return err
	}
	data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(adminToken))
	req, err := http.NewRequest("POST", adminURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create user, status code: %d, body: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (k *KeycloakService) GetByID(ctx context.Context, id string) (entity.User, error) {
	return entity.User{}, nil
}

func (k *KeycloakService) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return entity.User{}, nil
}
