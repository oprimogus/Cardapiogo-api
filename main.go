package main

import (
	"github.com/oprimogus/cardapiogo/config/logger"
	"github.com/oprimogus/cardapiogo/infra/database"
	"github.com/subosito/gotenv"
)

var (
	log = logger.GetLogger("Main")
)

func main() {
	_ = gotenv.Load()
	db := database.GetInstance()
	log.Info(&db)

}
