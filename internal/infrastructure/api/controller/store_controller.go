package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oprimogus/cardapiogo/internal/application/store"
	"github.com/oprimogus/cardapiogo/internal/domain/repository"
	validatorutils "github.com/oprimogus/cardapiogo/internal/infrastructure/api/validator"
	xerrors "github.com/oprimogus/cardapiogo/internal/infrastructure/errors"
)

type StoreController struct {
	validator *validatorutils.Validator
	create    store.Create
}

func NewStoreController(validator *validatorutils.Validator, repository repository.StoreRepository) *StoreController {
	return &StoreController{
		validator: validator,
		create:    store.NewCreate(repository),
	}
}

// CreateStore godoc
//
//	@Summary		User can create stores.
//	@Description	Owner user can create store
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			request	body	store.CreateParams	false	"CreateParams"
//	@Success		201
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/store [post]
func (c *StoreController) Create(ctx *gin.Context) {
	var createParams store.CreateParams
	err := ctx.BindJSON(&createParams)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	errValidate := c.validator.Validate(createParams)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	err = c.create.Execute(ctx, createParams)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.Status(http.StatusCreated)
}
