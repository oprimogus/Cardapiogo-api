package user_test

import (
	// "context"
	"os"
	"testing"

	// "github.com/oprimogus/cardapiogo/test/integration"
)

func TestMain(m *testing.M) {

	// ctx := context.Background()

	// keycloakContainer, err := integration.MakeKeycloak(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	code := m.Run()

	// keycloakContainer.Kill(ctx)

	os.Exit(code)
}