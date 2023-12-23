package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
	"github.com/oprimogus/cardapiogo/pkg/validator"
)

var (
	log      = logger.GetLogger("UserController")
	validate = validator.New(validator.WithRequiredStructEnabled())
)

type userController struct {
	service   Service
	validator *validatorutils.Validator
}

func NewUserController(repository Repository, validator *validatorutils.Validator) *userController {
	return &userController{service: NewService(repository), validator: validator}
}

func (c *userController) CreateUserHandler(ctx *gin.Context) {
	var userParams CreateUserParams
	err := ctx.BindJSON(&userParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	// Validação
	validationErr := c.validator.Validate(userParams)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	userCreated := c.service.CreateUser(ctx, userParams)
	if userCreated != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": userCreated})
	}
	ctx.Status(http.StatusCreated)

	// Resto do seu código para criar o usuário...
}
