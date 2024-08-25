package config

import (
	"os"

	"github.com/subosito/gotenv"

	"github.com/oprimogus/cardapiogo/internal/utils"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	conf *config
	log  = logger.NewLogger("Config")
)

type dbConfig struct {
	host     string
	port     string
	name     string
	user     string
	password string
}

func (d *dbConfig) Host() string {
	return d.host
}

func (d *dbConfig) Port() string {
	return d.port
}

func (d *dbConfig) SetPort(port string) {
	d.port = port
}

func (d *dbConfig) Name() string {
	return d.name
}

func (d *dbConfig) User() string {
	return d.user
}

func (d *dbConfig) Password() string {
	return d.password
}

type apiConfig struct {
	basePath  string
	port      string
	ginMode   string
	sqlcDebug string
}

func (a apiConfig) BasePath() string {
	return a.basePath
}

func (a apiConfig) Port() string {
	return a.port
}

func (a apiConfig) GinMode() string {
	return a.ginMode
}

func (a apiConfig) SQLCDebug() string {
	return a.sqlcDebug
}

type keycloakConfig struct {
	baseURL      string
	realm        string
	clientID     string
	clientSecret string
}

func (k keycloakConfig) BaseURL() string {
	return k.baseURL
}

func (k keycloakConfig) Realm() string {
	return k.realm
}

func (k keycloakConfig) ClientID() string {
	return k.clientID
}

func (k keycloakConfig) ClientSecret() string {
	return k.clientSecret
}

type resendConfig struct {
	apiKey string
}

func (r resendConfig) APIKey() string {
	return r.apiKey
}

type config struct {
	Database *dbConfig
	Api      apiConfig
	Keycloak keycloakConfig
	Resend   resendConfig
}

func newConfig() *config {
	err := utils.SetWorkingDirToProjectRoot()
	if err != nil {
		panic("fail on set project root as workdir")
	}
	err = gotenv.Load(".env")
	if err != nil {
		log.Errorf("fail on load env vars: %s", err)
		panic("fail on load env vars")
	}
	return &config{
		Database: &dbConfig{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			name:     os.Getenv("DB_NAME"),
			user:     os.Getenv("DB_USERNAME"),
			password: os.Getenv("DB_PASSWORD"),
		},
		Api: apiConfig{
			basePath:  os.Getenv("API_BASE_PATH"),
			port:      os.Getenv("API_PORT"),
			ginMode:   os.Getenv("GIN_MODE"),
			sqlcDebug: os.Getenv("SQLCDEBUG"),
		},
		Keycloak: keycloakConfig{
			baseURL:      os.Getenv("KEYCLOAK_BASE_URL"),
			realm:        os.Getenv("KEYCLOAK_REALM"),
			clientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
			clientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		},
		Resend: resendConfig{
			apiKey: os.Getenv("RESEND_API_KEY"),
		},
	}
}

func GetInstance() *config {
	if conf == nil {
		conf = newConfig()
	}
	return conf
}
