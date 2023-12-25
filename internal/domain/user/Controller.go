package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
	"github.com/oprimogus/cardapiogo/internal/api/validator"
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
		return
	}

	validationErr := c.validator.Validate(userParams)
	if len(validationErr) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	userCreated := c.service.CreateUser(ctx, userParams)
	if userCreated != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": userCreated})
		return
	}
	ctx.Status(http.StatusCreated)
}
