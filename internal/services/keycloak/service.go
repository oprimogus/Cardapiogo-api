package keycloak

import (
	"context"
	"fmt"
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
	Client       *gocloak.GoCloak
	Token        *gocloak.JWT
	Realm        string
	ClientID     string
	ClientSecret string
	mu           sync.Mutex
}

func GetInstance(ctx context.Context) (k *KeycloakService, err error) {
	if keycloakInstance == nil {
		keycloakInstance, err = newKeycloakService(ctx)
		if err != nil {
			return keycloakInstance, err
		}
	}
	return keycloakInstance, nil
}

func newKeycloakService(ctx context.Context) (k *KeycloakService, err error) {

	config := config.GetInstance().Keycloak

	client := gocloak.NewClient(config.BaseURL)
	token, err := client.LoginClient(ctx, config.ClientID, config.ClientSecret, config.Realm)
	if err != nil {
		log.Errorf("fail on get keycloak client: %s", err)
		return nil, fmt.Errorf("fail on get keycloak client: %w", err)
	}
	service := &KeycloakService{
		Client:       client,
		Token:        token,
		Realm:        config.Realm,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
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

	accessTokenTicker, accessTokenExpireIn := k.startTicker("accessToken", refreshBefore, k.Token.ExpiresIn)
	defer accessTokenTicker.Stop()

	refreshTokenTicker, refreshTokenExpireIn := k.startTicker("refreshToken", refreshBefore, k.Token.RefreshExpiresIn)
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
			token, err = k.Client.LoginClient(ctx, k.ClientID, k.ClientSecret, k.Realm)
		} else {
			token, err = k.Client.RefreshToken(ctx, k.Token.RefreshToken, k.ClientID, k.ClientSecret, k.Realm)
		}

		if err != nil {
			log.Errorf("Failed to renew token: %v", err)
			return
		}

		k.Token = token
		log.Infof("Token refreshed successfully (isRefreshToken: %v)", isRefreshToken)
	}
}
