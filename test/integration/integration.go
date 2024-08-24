package integration

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type Container struct {
	name     string
	instance testcontainers.Container
}

func (c *Container) Kill(ctx context.Context) {
	if err := c.instance.Terminate(ctx); err != nil {
		fmt.Printf("could not stop %s: %s", c.name, err)
	}
}
