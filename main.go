package main

import (
	"github.com/oprimogus/cardapiogo/internal/api/router"
	"github.com/oprimogus/cardapiogo/internal/infra/database"
	"github.com/oprimogus/cardapiogo/internal/infra/database/sqlc"
	"github.com/oprimogus/cardapiogo/internal/infra/factory"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
	"github.com/subosito/gotenv"
)


func main() {
	log := logger.GetLogger("Main")
	// _ = gotenv.Load()
	_ = gotenv.Load()
	db := database.GetInstance()
	factory := factory.NewDataBaseRepositoryFactory(db)
	// factory := factory.NewDataBaseRepositoryFactory(db)
	router.Initialize(factory)

}
