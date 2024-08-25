package main

import (
	"github.com/oprimogus/cardapiogo/internal/api/router"
	"github.com/oprimogus/cardapiogo/internal/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/persistence"
	"github.com/subosito/gotenv"
)

//	@title			Cardapiogo API
//	@version		1.0
//	@description	Documentação da API de delivery Cardapiogo.
//	@contact.name	Gustavo Ferreira de Jesus
//	@contact.email	gustavo081900@gmail.com

//	@host		localhost:8080
//	@BasePath	/api
//	@accept		json
//	@produce	json

// @securityDefinitions.apikey	Bearer Token
// @in							header
// @name						Authorization
func main() {
	// env
	_ = gotenv.Load()

	// database
	db := postgres.GetInstance()
	defer db.Close()

	// routes
	factoryRepository := persistence.NewDataBaseRepositoryFactory(db)
	router.Initialize(factoryRepository)
}
