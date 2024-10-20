package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/api/middleware"
	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/core/user"
	xerrors "github.com/oprimogus/cardapiogo/internal/errors"
)

type UserController struct {
	validator  *validatorutils.Validator
	userModule user.UserModule
}

func NewUserController(validator *validatorutils.Validator, userRepository user.Repository) *UserController {
	return &UserController{
		validator:  validator,
		userModule: user.NewUserModule(userRepository),
	}
}

// CreateUser godoc
//
//	@Summary		Sign-Up with local credentials and data
//	@Description	Sign-Up with local credentials and data
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	user.CreateParams	false	"CreateUserParams"
//	@Success		201
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/auth/sign-up [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	transactionID := ctx.GetString(middleware.TransactionIDLabel)
	fmt.Println(transactionID)
	var params user.CreateParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.HandleError(err, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	errValidate := c.validator.Validate(params, transactionID)
	if errValidate != nil {
		xerror := xerrors.HandleError(errValidate, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	er := c.userModule.Create.Execute(ctx, params)
	if er != nil {
		xerror := xerrors.HandleError(er, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.Status(http.StatusCreated)
}

// UpdateUser godoc
//
//	@Summary		Update user profile
//	@Description	Update user profile
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	user.UpdateProfileParams	false	"UpdateProfileParams"
//	@Success		200
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Security		Bearer Token
//	@Router			/v1/user [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	transactionID := ctx.GetString(middleware.TransactionIDLabel)
	var params user.UpdateProfileParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.HandleError(err, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	errValidate := c.validator.Validate(params, transactionID)
	if errValidate != nil {
		xerror := xerrors.HandleError(errValidate, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	er := c.userModule.Update.Execute(ctx, params)
	if er != nil {
		xerror := xerrors.HandleError(er, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.Status(http.StatusOK)
}

// SetRoleInUser godoc
//
//	@Summary		Add a new role for user
//	@Description	Add a new role for user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	user.AddRolesParams	false	"AddRolesParams"
//	@Success		200
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Security		Bearer Token
//	@Router			/v1/user/roles [post]
func (c *UserController) AddRolesToUser(ctx *gin.Context) {
	transactionID := ctx.GetString(middleware.TransactionIDLabel)
	var params user.AddRolesParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.HandleError(err, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	errValidate := c.validator.Validate(params, transactionID)
	if errValidate != nil {
		xerror := xerrors.HandleError(errValidate, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	er := c.userModule.AddRoles.Execute(ctx, params.Roles)
	if er != nil {
		xerror := xerrors.HandleError(er, transactionID)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.Status(http.StatusOK)
}
