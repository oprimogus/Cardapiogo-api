package main

import (
	"github.com/subosito/gotenv"

	"github.com/oprimogus/cardapiogo/internal/api/router"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/factory"
)

// @title           Cardapio-Go
// @version         1.0
// @description     Simple API of Cardapio
// @termsOfService  http://swagger.io/terms/

// @contact.name   Gustavo Ferreira
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey JWT Token
// @in header Cookie: token=$VALUE
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// env
	_ = gotenv.Load()

	// database
	db := postgres.GetInstance()
	defer db.Close()

	// routes
	factory := factory.NewDataBaseRepositoryFactory(db)
	router.Initialize(factory)
}
