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
	update store.Update
	getByID store.GetByID
}

func NewStoreController(validator *validatorutils.Validator, repository repository.StoreRepository) *StoreController {
	return &StoreController{
		validator: validator,
		create:    store.NewCreate(repository),
		update: store.NewUpdate(repository),
		getByID: store.NewGetByID(repository),
	}
}

// CreateStore godoc
//
//	@Summary		Owner can create stores.
//	@Description	Owner user can create store
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			request	body	store.CreateParams	false	"CreateParams"
//	@Success		201
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		401	{object}	xerrors.ErrorResponse
//	@Failure		403	{object}	xerrors.ErrorResponse
//	@Failure		409	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/store [post]
func (c *StoreController) Create(ctx *gin.Context) {
	var params store.CreateParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	errValidate := c.validator.Validate(params)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	err = c.create.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.Status(http.StatusCreated)
}

// UpdateStore godoc
//
//	@Summary		Owner can update your stores.
//	@Description	Owner can update your stores.
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			request	body	store.UpdateParams	false	"UpdateParams"
//	@Success		200
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		401	{object}	xerrors.ErrorResponse
//	@Failure		403	{object}	xerrors.ErrorResponse
//	@Failure		409	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/store [put]
func (c *StoreController) Update(ctx *gin.Context) {
	var params store.UpdateParams
	err := ctx.BindJSON(&params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	errValidate := c.validator.Validate(params)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	err = c.update.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.Status(http.StatusOK)
}

// GetStoreByID godoc
//
//	@Summary		Any user can view a store.
//	@Description	Any user can view a store.
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Store ID"
//	@Success		200	{object}	store.GetStoreByIdOutput
//	@Failure		404	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/store/{id} [get]
func (c *StoreController) GetStoreByID(ctx *gin.Context) {
	id := ctx.Param("id")
	store, err := c.getByID.Execute(ctx, id)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.JSON(http.StatusOK, store)
}
