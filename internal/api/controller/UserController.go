package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	errordatabase "github.com/oprimogus/cardapiogo/internal/errors/database"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	log      = logger.GetLogger("UserController")
	validate = validator.New(validator.WithRequiredStructEnabled())
)

type userController struct {
	service   user.Service
	validator *validatorutils.Validator
}

func NewUserController(repository user.Repository, validator *validatorutils.Validator) *userController {
	return &userController{service: user.NewService(repository), validator: validator}
}

func (c *userController) CreateUserHandler(ctx *gin.Context) {
	var userParams user.CreateUserParams
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

	userCreatedErr := c.service.CreateUser(ctx, userParams)
	if userCreatedErr != nil {
		dbErr, ok := userCreatedErr.(*errordatabase.DBError)
		if ok {
			ctx.JSON(dbErr.HttpStatus, gin.H{"error": dbErr.Message})
			return
		} 
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *userController) GetUserHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	_, err := uuid.Parse(id)
	if (err != nil) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
}
