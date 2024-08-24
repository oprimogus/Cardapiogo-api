package integration

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func MakeKeycloak(ctx context.Context) (*Container, error) {
	reqKeycloak := testcontainers.ContainerRequest{
		Image:        "quay.io/keycloak/keycloak:24.0.4",
		ExposedPorts: []string{"8081/tcp"},
		WaitingFor:   wait.ForLog("Listening on: http://0.0.0.0:8080"),
	}
	keycloakTest, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: reqKeycloak,
		Started:          true,
	})
	if err != nil {
		return  nil, fmt.Errorf("could not start Keycloak: %w", err)
	}

	return &Container{name: "Keycloak", instance: keycloakTest}, nil
}