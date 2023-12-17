package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oprimogus/cardapiogo/config/logger"
	"github.com/oprimogus/cardapiogo/database"
	"github.com/oprimogus/cardapiogo/database/sqlc"
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


	password := pgtype.Text {
		String: "testeteste123",
		Valid: true,
	}

	userParams := sqlc.CreateUserParams {
		Email: "gustavo081900@gmail.com",
		Password: password,
		Role: sqlc.NullCardapioUserRole{CardapioUserRole: sqlc.CardapioUserRoleConsumer, Valid: true},
		AccountProvider: sqlc.CardapioAccountProviderGoogle,
	}

	a := queries.CreateUser(ctx, userParams)
	log.Info(a)

}
