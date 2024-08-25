package integration

import (
	"context"

	logger "github.com/oprimogus/cardapiogo/pkg/log"
	"github.com/testcontainers/testcontainers-go"
)

var log = logger.NewLogger("Integration")

type Container struct {
	name     string
	instance testcontainers.Container
}

func (c *Container) Kill(ctx context.Context) {
	if err := c.instance.Terminate(ctx); err != nil {
		log.Errorf("could not stop %s: %s", c.name, err)
	}
}
