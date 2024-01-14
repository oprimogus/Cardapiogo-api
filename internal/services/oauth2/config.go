package oauth2

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

func NewGoogleOauthConf() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectUrl := os.Getenv("GOOGLE_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
    		TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
}

func InitializeOauth2(ctx context.Context) {
	log := logger.GetLoggerDefault("Oauth2.0")

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "",
		Scopes:       []string{"scope1", "scope2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "",
			TokenURL: "",
		},
	}

	verifier := oauth2.GenerateVerifier()
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Error(err)
	}
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Error(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")
}
