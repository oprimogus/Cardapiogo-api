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

// @BasePath /api/v1

// CreateUserHandler godoc
// @Summary Adiciona um novo usuário
// @Description Cria um novo usuário através de login email/senha
// @Tags User
// @Accept  json
// @Produce  json
// @Param   request body user.CreateUserParams true "CreateUserParams"
// @Success 201
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /user [post]
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
	validateErrorResponse(ctx, err)
	ctx.Status(http.StatusCreated)
}

// GetUserHandler godoc
// @Summary Retorna um usuário
// @Description Retorna um usuário através do ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param   id path string true "ID do usuário (UUID)"
// @Success 200  {object} user.User
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /user/{id} [get]
func (c *UserController) GetUserHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	getUser, err := c.service.GetUser(ctx, id)
	validateErrorResponse(ctx, err)
	ctx.JSON(http.StatusOK, getUser)
}

// GetUsersListHandler godoc
// @Summary Retorna uma lista de usuários
// @Description Retorna uma lista de usuários com paginação
// @Tags User
// @Accept  json
// @Produce  json
// @Param   items query number true "Número de itens por página"
// @Param   page query number true "Página"
// @Success 200 {array} user.User
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /user [get]
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
	validateErrorResponse(ctx, err)
	ctx.JSON(http.StatusOK, listUsers)
}

// UpdateUserPasswordHandler godoc
// @Summary Atualiza a senha do usuário
// @Description Atualiza a senha do usuário
// @Tags User
// @Accept  json
// @Produce  json
// @Param   request body user.UpdateUserPasswordParams true "UpdateUserPasswordParams"
// @Success 200
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /user/change-password [put]
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

	validateIsSameUser(ctx, updateUserPasswordParams.ID)

	err = c.service.UpdateUserPassword(ctx, updateUserPasswordParams)
	validateErrorResponse(ctx, err)
	ctx.Status(http.StatusOK)
}

// UpdateUserHandler godoc
// @Summary Atualiza os dados do usuário
// @Description Atualiza os dados do usuário
// @Tags User
// @Accept  json
// @Produce  json
// @Param   request body user.UpdateUserParams true "UpdateUserParams"
// @Success 200
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /user [put]
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

	validateIsSameUser(ctx, updateUserParams.ID)

	err = c.service.UpdateUser(ctx, updateUserParams)
	validateErrorResponse(ctx, err)

	ctx.Status(http.StatusOK)
}
