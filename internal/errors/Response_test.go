package errors

import (
	"log"
	"testing"

	"github.com/go-playground/assert/v2"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
)

func TestInvalidInput(t *testing.T) {

	createUser := user.CreateUserParams{
		Email:           "invalidEmail",
		Password:        "invalidEmail",
		Role:            "invalidEmail",
		AccountProvider: "invalidEmail",
	}

	v, _ := validatorutils.NewValidator("en")

	test := v.Validate(createUser)
	log.Println(test)
	InvalidInput()

	assert.IsEqual(test, test)
}
