package controller

import (
	"net/http"
	"strconv"

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

// UserController struct
type UserController struct {
	service   user.Service
	validator *validatorutils.Validator
}

// NewUserController return a new instance of userController
func NewUserController(repository user.Repository, validator *validatorutils.Validator) *UserController {
	return &UserController{service: user.NewService(repository), validator: validator}
}

// CreateUserHandler create a user in database
func (c *UserController) CreateUserHandler(ctx *gin.Context) {
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
		if ok && dbErr != nil {
			ctx.JSON(dbErr.HttpStatus, gin.H{"error": dbErr.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.Status(http.StatusCreated)
}

// GetUserHandler return a user by id
func (c *UserController) GetUserHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	uuidValue, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	getUser, err := c.service.GetUser(ctx, uuidValue)
	if err != nil {
		dbErr, ok := err.(*errordatabase.DBError)
		if ok {
			ctx.JSON(dbErr.HttpStatus, gin.H{"error": dbErr.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, getUser)
}

// GetUsersListHandler return a paginated list of users
func (c *UserController) GetUsersListHandler(ctx *gin.Context) {
	items, err := strconv.Atoi(ctx.Query("items"))
	log.Info(items)
	if err != nil || items <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for items"})
		return
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	log.Info(page)
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for page"})
		return
	}

	listUsers, err := c.service.GetUsersList(ctx, items, page)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error listing users"})
		return
	}

	ctx.JSON(http.StatusOK, listUsers)
	return
}
