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
	Host     string
	Port     string
	Name     string
	User     string
	Password string
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
	BaseURL      string
	Realm        string
	ClientID     string
	ClientSecret string
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
	Keycloak *keycloakConfig
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
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Api: apiConfig{
			basePath:  os.Getenv("API_BASE_PATH"),
			port:      os.Getenv("API_PORT"),
			ginMode:   os.Getenv("GIN_MODE"),
			sqlcDebug: os.Getenv("SQLCDEBUG"),
		},
		Keycloak: &keycloakConfig{
			BaseURL:      os.Getenv("KEYCLOAK_BASE_URL"),
			Realm:        os.Getenv("KEYCLOAK_REALM"),
			ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
			ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
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
