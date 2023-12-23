package main

import (
	"github.com/oprimogus/cardapiogo/internal/api/router"
	"github.com/oprimogus/cardapiogo/internal/infra/database"
	"github.com/oprimogus/cardapiogo/internal/infra/factory"
	"github.com/subosito/gotenv"
)

func main() {
	// _ = gotenv.Load()
	_ = gotenv.Load()
	db := database.GetInstance()
	defer db.Close()
	factory := factory.NewDataBaseRepositoryFactory(db)
	router.Initialize(factory)

}
