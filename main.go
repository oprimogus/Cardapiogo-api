package main

import (
	"github.com/oprimogus/cardapiogo/internal/api/router"
	"github.com/oprimogus/cardapiogo/internal/infra/database/postgres"
	"github.com/oprimogus/cardapiogo/internal/infra/factory"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load()
	db := postgres.GetInstance()
	defer db.Close()
	factory := factory.NewDataBaseRepositoryFactory(db)
	router.Initialize(factory)
}
