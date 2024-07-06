package keycloak

import (
	"context"
	"os"
	"sync"

	"github.com/Nerzal/gocloak/v13"

	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	log          = logger.NewLogger("Keycloak")
	keycloakURL  string
	realm        string
	clientID     string
	clientSecret string
)

func init() {
	keycloakURL = os.Getenv("KEYCLOAK_BASE_URL")
	realm = os.Getenv("KEYCLOAK_REALM")
	clientID = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
}

type KeycloakService struct {
	client       *gocloak.GoCloak
	token        *gocloak.JWT
	realm        string
	clientID     string
	clientSecret string
	mu           sync.Mutex
}

func NewKeycloakService(ctx context.Context) (*KeycloakService, error) {
	client := gocloak.NewClient(keycloakURL)
	token, err := client.LoginClient(ctx, clientID, clientSecret, realm)
	if err != nil {
		return nil, err
	}
	service := &KeycloakService{
		client:       client,
		token:        token,
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientSecret,
	}

	go service.startTokenRenewer(ctx, token.ExpiresIn)

	return service, nil
}

func (k *KeycloakService) startTokenRenewer(ctx context.Context, expiresIn int) {

}
