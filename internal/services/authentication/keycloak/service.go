package keycloak

import (
	"context"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v13"

	"github.com/oprimogus/cardapiogo/internal/config"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	log              = logger.NewLogger("Keycloak")
	keycloakInstance *KeycloakService
)

type KeycloakService struct {
	client       *gocloak.GoCloak
	token        *gocloak.JWT
	realm        string
	clientID     string
	clientSecret string
	mu           sync.Mutex
}

func GetInstance(ctx context.Context) *KeycloakService {
	if keycloakInstance == nil {
		keycloakInstance = newKeycloakService(ctx)
	}
	return keycloakInstance
}

func newKeycloakService(ctx context.Context) *KeycloakService {

	config := config.GetInstance().Keycloak

	client := gocloak.NewClient(config.BaseURL())
	token, err := client.LoginClient(ctx, config.ClientID(), config.ClientSecret(), config.Realm())
	if err != nil {
		log.Errorf("fail on get keycloak client: %s", err)
		panic(err)
	}
	service := &KeycloakService{
		client:       client,
		token:        token,
		realm:        config.Realm(),
		clientID:     config.ClientID(),
		clientSecret: config.ClientSecret(),
	}

	go service.startTokenRenewer(ctx)

	return service
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
