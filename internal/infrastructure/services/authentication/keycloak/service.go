package keycloak

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v13"

	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	log = logger.NewLogger("Keycloak")
)

type KeycloakService struct {
	client       *gocloak.GoCloak
	token        *gocloak.JWT
	realm        string
	clientID     string
	clientSecret string
	mu           sync.Mutex
}

func NewKeycloakService(ctx context.Context) (*KeycloakService, error) {

	keycloakURL := os.Getenv("KEYCLOAK_BASE_URL")
	realm := os.Getenv("KEYCLOAK_REALM")
	clientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")

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

	go service.startTokenRenewer(ctx)

	return service, nil
}

func shouldRefreshToken(expirationTime time.Time) bool {
	return time.Now().After(expirationTime)
}

func (k *KeycloakService) startTokenRenewer(ctx context.Context) {
	log.Info("Starting renew service...")

	refreshBefore := time.Second * 60 // 1 min

	accessTokenTicker, accessTokenExpireIn := k.startTicker("accessToken", refreshBefore, k.token.ExpiresIn)
	defer accessTokenTicker.Stop()

	refreshTokenTicker, refreshTokenExpireIn := k.startTicker("refreshToken", refreshBefore, k.token.RefreshExpiresIn)
	defer refreshTokenTicker.Stop()

	for {
		select {
		case <-accessTokenTicker.C:
			k.handleTokenRenewal(ctx, accessTokenExpireIn, false)
		case <-refreshTokenTicker.C:
			k.handleTokenRenewal(ctx, refreshTokenExpireIn, true)
		case <-ctx.Done():
			log.Info("AccessToken renewal stopped")
			return
		}
	}
}

func (k *KeycloakService) startTicker(name string, renewBefore time.Duration, tokenExpiresIn int) (ticker *time.Ticker, willExpireIn time.Time) {
	actualTime := time.Now()
	tokenExpireIn := time.Second * time.Duration(tokenExpiresIn)
	willExpireIn = actualTime.Add(time.Duration(tokenExpireIn) - renewBefore)
	log.Infof("token %s will expire in: %s", name, willExpireIn)

	return time.NewTicker(time.Duration(tokenExpireIn) - renewBefore), willExpireIn
}

func (k *KeycloakService) handleTokenRenewal(ctx context.Context, tokenExpireIn time.Time, isRefreshToken bool) {
	k.mu.Lock()
	defer k.mu.Unlock()

	if shouldRefreshToken(tokenExpireIn) {
		var token *gocloak.JWT
		var err error

		if isRefreshToken {
			token, err = k.client.LoginClient(ctx, k.clientID, k.clientSecret, k.realm)
		} else {
			token, err = k.client.RefreshToken(ctx, k.token.RefreshToken, k.clientID, k.clientSecret, k.realm)
		}

		if err != nil {
			log.Errorf("Failed to renew token: %v", err)
			return
		}

		k.token = token
		log.Infof("Token refreshed successfully (isRefreshToken: %v)", isRefreshToken)
	}
}
