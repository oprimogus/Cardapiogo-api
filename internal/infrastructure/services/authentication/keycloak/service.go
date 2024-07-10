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

	go service.startTokenRenewer(ctx)

	return service, nil
}

func (k *KeycloakService) shouldRefreshToken(initialTime time.Time, duration time.Duration) bool {
	now := time.Now()
	maxTime := initialTime.Add(duration)
	return now.After(maxTime)
}

func (k *KeycloakService) startTokenRenewer(ctx context.Context) {
	actualTime := time.Now()
	refreshBefore := time.Second * 60

	accessTokenExpireIn := actualTime.Add(time.Duration(k.token.ExpiresIn) * time.Second)
	accessTokenShouldRenewIn := accessTokenExpireIn.Sub(actualTime) - refreshBefore
	tickerAccessToken := time.NewTicker(accessTokenShouldRenewIn)
	defer tickerAccessToken.Stop()

	refreshTokenExpireIn := actualTime.Add(time.Duration(k.token.RefreshExpiresIn) * time.Second)
	refreshTokenShouldRenewIn := refreshTokenExpireIn.Sub(actualTime) - refreshBefore
	tickerRefreshToken := time.NewTicker(refreshTokenShouldRenewIn)
	defer tickerRefreshToken.Stop()

	for {
		select {
		case <-tickerAccessToken.C:
			k.mu.Lock()
			if k.shouldRefreshToken(actualTime, accessTokenShouldRenewIn) {
				token, err := k.client.RefreshToken(ctx, k.token.RefreshToken, clientID, clientSecret, realm)
				if err != nil {
					log.Errorf("Failed in refresh client token: %v", err)
				} else {
					k.token = token
					log.Info("AccessToken refreshed successfully")
					actualTime = time.Now()
					accessTokenExpireIn = actualTime.Add(time.Duration(k.token.ExpiresIn) * time.Second)
					accessTokenShouldRenewIn = accessTokenExpireIn.Sub(actualTime) - refreshBefore
					tickerAccessToken.Reset(accessTokenShouldRenewIn)
				}
			}
			k.mu.Unlock()
		case <-tickerRefreshToken.C:
			k.mu.Lock()
			if k.shouldRefreshToken(actualTime, refreshTokenShouldRenewIn) {
				token, err := k.client.LoginClient(ctx, clientID, clientSecret, realm)
				if err != nil {
					log.Errorf("Failed in authentication: %v", err)
				} else {
					k.token = token
					log.Info("RefreshToken refreshed successfully")

					actualTime = time.Now()
					refreshTokenExpireIn = actualTime.Add(time.Duration(k.token.RefreshExpiresIn) * time.Second)
					refreshTokenShouldRenewIn = refreshTokenExpireIn.Sub(actualTime) - refreshBefore
					tickerRefreshToken.Reset(refreshTokenShouldRenewIn)

					accessTokenExpireIn = actualTime.Add(time.Duration(k.token.ExpiresIn) * time.Second)
					accessTokenShouldRenewIn = accessTokenExpireIn.Sub(actualTime) - refreshBefore
					tickerAccessToken.Reset(accessTokenShouldRenewIn)
				}
			}
			k.mu.Unlock()
		case <-ctx.Done():
			log.Info("AccessToken renewal stopped")
			return
		}
	}
}
