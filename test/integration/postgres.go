package integration

import (
	"context"
	"path/filepath"
	"strings"

	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/oprimogus/cardapiogo/internal/config"
	"github.com/oprimogus/cardapiogo/internal/utils"
)

func MakePostgres(ctx context.Context) (*Container, error) {
	_ = utils.SetWorkingDirToProjectRoot()
	config := config.GetInstance().Database
	config.Host = "localhost"
	config.User = "cardapiogo"
	config.Name = "postgres"
	config.Password = "cardapiogo"
	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithInitScripts(filepath.Join("test", "integration", "testdata", "postgres-init.sh")),
		postgres.WithDatabase(config.Name),
		postgres.WithUsername(config.User),
		postgres.WithPassword(config.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Errorf("failed to start container: %s", err)
		return nil, err
	}

	hostPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Errorf("failed to get mapped port: %s", err)
		return nil, err
	}
	config.Port = strings.Replace(string(hostPort), "/tcp", "", -1)

	return &Container{name: "postgres", instance: postgresContainer}, nil
}
