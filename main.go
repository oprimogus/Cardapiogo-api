package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oprimogus/cardapiogo/config/logger"
	database "github.com/oprimogus/cardapiogo/database/sqlc"
	"github.com/oprimogus/cardapiogo/infra/database2"
	"github.com/subosito/gotenv"
)

var (
	log = logger.GetLogger("Main")
)

func main() {
	_ = gotenv.Load()
	ctx := context.Background()
	db := database2.GetInstance()
	log.Info(&db)

	queries := database.New(db)

	password := pgtype.Text {
		String: "testeteste123",
		Valid: true,
	}

	userParams := database.CreateUserParams {
		Email: "gustavo081900@gmail.com",
		Password: password,
		Role: database.NullUserRole{UserRole: database.UserRoleConsumer, Valid:  true},
		AccountProvider: database.AccountProviderGoogle,
	}

	queries.CreateUser(ctx, userParams)

}
