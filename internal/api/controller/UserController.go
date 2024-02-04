package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

// UserController struct
type UserController struct {
	service   *user.Service
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
		ctx.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	validationErr := c.validator.Validate(userParams)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, validationErr)
		return
	}

	err = c.service.CreateUser(ctx, userParams)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if ok && dbErr != nil {
			ctx.JSON(dbErr.Status, dbErr)
			return
		}
		ctx.JSON(http.StatusInternalServerError, dbErr)
		return
	}
	ctx.Status(http.StatusCreated)
}

// GetUserHandler return a user by id
func (c *UserController) GetUserHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	getUser, err := c.service.GetUser(ctx, id)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if ok && dbErr != nil {
			ctx.JSON(dbErr.Status, dbErr.ErrorMessage)
			return
		}
		ctx.JSON(http.StatusInternalServerError, dbErr)
		return
	}
	ctx.JSON(http.StatusOK, getUser)
}

// GetUsersListHandler return a paginated list of users
func (c *UserController) GetUsersListHandler(ctx *gin.Context) {
	items, err := strconv.Atoi(ctx.Query("items"))
	if err != nil || items <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for items"})
		return
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for page"})
		return
	}

	listUsers, err := c.service.GetUsersList(ctx, items, page)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if ok && dbErr != nil {
			ctx.JSON(dbErr.Status, dbErr.ErrorMessage)
			return
		}
		ctx.JSON(http.StatusInternalServerError, dbErr)
		return
	}
	ctx.JSON(http.StatusOK, listUsers)
}

// UpdateUserPasswordHandler update user password
func (c *UserController) UpdateUserPasswordHandler(ctx *gin.Context) {

	var updateUserPasswordParams user.UpdateUserPasswordParams
	err := ctx.BindJSON(&updateUserPasswordParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	validationErr := c.validator.Validate(updateUserPasswordParams)
	if validationErr != nil {
		ctx.JSON(validationErr.Status, validationErr)
		return
	}

	err = validatorutils.IsSameUser(ctx, updateUserPasswordParams.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	err = c.service.UpdateUserPassword(ctx, updateUserPasswordParams)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if ok && dbErr != nil {
			ctx.JSON(dbErr.Status, dbErr)
			return
		}
		ctx.JSON(dbErr.Status, dbErr)
		return
	}
	ctx.Status(http.StatusOK)
}

// UpdateUserHandler update user password
func (c *UserController) UpdateUserHandler(ctx *gin.Context) {

	var updateUserParams user.UpdateUserParams
	err := ctx.BindJSON(&updateUserParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	validationErr := c.validator.Validate(updateUserParams)
	if validationErr != nil {
		ctx.JSON(validationErr.Status, validationErr)
		return
	}

	err = validatorutils.IsSameUser(ctx, updateUserParams.ID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	err = c.service.UpdateUser(ctx, updateUserParams)
	if err != nil {
		dbErr, ok := err.(*errors.ErrorResponse)
		if ok && dbErr != nil {
			ctx.JSON(dbErr.Status, dbErr)
			return
		}
		ctx.JSON(dbErr.Status, dbErr)
		return
	}

	ctx.Status(http.StatusOK)
}
