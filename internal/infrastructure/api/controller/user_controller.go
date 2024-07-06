package controller

import (
	// "errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/application/user"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	xerrors "github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
	// xerrors "github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

type UserController struct {
	validator *validatorutils.Validator
	create    user.Create
	update    user.Update
	delete    user.Delete
}

func NewUserController(validator *validatorutils.Validator, userRepository repository.UserRepository) *UserController {
	return &UserController{
		validator: validator,
		create:    user.NewCreate(userRepository),
		update:    user.NewUpdate(userRepository),
		delete:    user.NewDelete(userRepository),
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
	var signUpParams user.CreateParams
	err := ctx.BindJSON(&signUpParams)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	errValidate := c.validator.Validate(signUpParams)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	er := c.create.Execute(ctx, signUpParams)
	if er != nil {
		xerror := xerrors.Map(er)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.Status(http.StatusCreated)
}
