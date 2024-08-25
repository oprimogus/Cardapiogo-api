package integration

import (
	"context"
	"strings"

	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/oprimogus/cardapiogo/internal/config"
)

func MakePostgres(ctx context.Context) (*Container, error) {
	config := config.GetInstance().Database
	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(config.Name()),
		postgres.WithUsername(config.User()),
		postgres.WithPassword(config.Password()),
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
	config.SetPort(strings.Replace(string(hostPort), "/tcp", "", -1))

	return &Container{name: "postgres", instance: postgresContainer}, nil
}
