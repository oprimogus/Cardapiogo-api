package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

func NewGoogleOauthConf() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	host := os.Getenv("API_HOST")
	redirectPath := os.Getenv("GOOGLE_REDIRECT_PATH")

	redirectUrl := fmt.Sprintf("http://%s://%s", host, redirectPath)

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GetGoogleUserData(ctx context.Context, conf *oauth2.Config, code string) (*GoogleUserInfo, error) {
	// log := logger.GetLogger("OAuth2.0", ctx)

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, errors.New(http.StatusBadGateway, err.Error())
	}

	client := conf.Client(ctx, token)
	url := os.Getenv("GOOGLE_USER_INFO_URL")

	res, err := client.Get(url)
	if err != nil {
		return nil, errors.New(http.StatusBadGateway, err.Error())
	}
	defer res.Body.Close()

	var userinfo GoogleUserInfo
	if err = json.NewDecoder(res.Body).Decode(&userinfo); err != nil {
		return nil, errors.New(http.StatusBadGateway, err.Error())
	}

	return &userinfo, nil

}
