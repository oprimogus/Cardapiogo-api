package persistence_test

import (
	"context"
	"os"
	"testing"

	"github.com/oprimogus/cardapiogo/test/integration"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	postgres, err := integration.MakePostgres(ctx)
	if err != nil {
		panic("fail on initialize postgres with testContainers")
	}
	defer postgres.Kill(ctx)
	
	code := m.Run()
	os.Exit(code)
}