package main

import (
	"os"

	"github.com/subosito/gotenv"

	"github.com/oprimogus/cardapiogo/docs"
	"github.com/oprimogus/cardapiogo/internal/api/router"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/factory"
)

// @termsOfService  http://swagger.io/terms/

// @contact.name   Gustavo Ferreira
// @contact.url    http://www.swagger.io/support
// @contact.email  gustavo081900@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey JWT Token
// @in header Cookie: token=$VALUE
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// env
	_ = gotenv.Load()

	configureSwaggerDocs()

	// database
	db := postgres.GetInstance()
	defer db.Close()

	// routes
	factory := factory.NewDataBaseRepositoryFactory(db)
	router.Initialize(factory)
}

// configureSwaggerDocs Configuração da documentação Swagger
func configureSwaggerDocs() {
	docs.SwaggerInfo.Title = "Cardapio-Go"
	docs.SwaggerInfo.Description = "Simple API of Cardapio"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("API_HOST")
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
