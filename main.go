package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/database"
	"github.com/oprimogus/cardapiogo/internal/database/sqlc"
	"github.com/oprimogus/cardapiogo/pkg/log"
	"github.com/subosito/gotenv"
)

var (
	log = logger.GetLogger("Main")
)

func main() {
	_ = gotenv.Load()
	ctx := context.Background()
	db := database.GetInstance()
	queries := sqlc.New(db.GetDB())

	userRepository := user.NewSQLUserRepository(queries)

}
