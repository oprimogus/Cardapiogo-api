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
	validator           *validatorutils.Validator
	create              store.Create
	update              store.Update
	addBusinessHours store.AddBusinessHour
	deleteBusinessHours store.DeleteBusinessHour
	getByID             store.GetByID
}

func NewStoreController(validator *validatorutils.Validator, repository repository.StoreRepository) *StoreController {
	return &StoreController{
		validator:           validator,
		create:              store.NewCreate(repository),
		update:              store.NewUpdate(repository),
		addBusinessHours: store.NewUpdateBusinessHour(repository),
		deleteBusinessHours: store.NewDeleteBusinessHour(repository),
		getByID:             store.NewGetByID(repository),
	}
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
	storeInstance, err := c.getByID.Execute(ctx, id)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.JSON(http.StatusOK, storeInstance)
}

// Create godoc
//
//	@Summary		Owner can create stores.
//	@Description	Owner user can create store
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			Params	body	store.CreateParams	true	"Params to create a store"
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

// Update godoc
//
//	@Summary		Owner can update your stores.
//	@Description	Owner can update your stores.
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			Params	body	store.UpdateParams	true	"Params to update a store"
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

// AddBusinessHours godoc
//
//	@Summary		Owner can update business hours of store.
//	@Description	Owner can update business hours of store.
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			Params	body	store.StoreBusinessHoursParams	true	"Params to update business hours of store"
//	@Success		200
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		401	{object}	xerrors.ErrorResponse
//	@Failure		403	{object}	xerrors.ErrorResponse
//	@Failure		409	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/store/business-hours [put]
func (c *StoreController) AddBusinessHours(ctx *gin.Context) {
	var params store.StoreBusinessHoursParams
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
	err = c.addBusinessHours.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteBusinessHours godoc
//
//	@Summary		Owner can delete business hours of store.
//	@Description	Owner can delete business hours of store.
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			Params	body	store.StoreBusinessHoursParams	true	"Params to delete business hours of store"
//	@Success		200
//	@Failure		400	{object}	xerrors.ErrorResponse
//	@Failure		401	{object}	xerrors.ErrorResponse
//	@Failure		403	{object}	xerrors.ErrorResponse
//	@Failure		409	{object}	xerrors.ErrorResponse
//	@Failure		500	{object}	xerrors.ErrorResponse
//	@Failure		502	{object}	xerrors.ErrorResponse
//	@Router			/v1/store/business-hours [delete]
func (c *StoreController) DeleteBusinessHours(ctx *gin.Context) {
	var params store.StoreBusinessHoursParams
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
	err = c.deleteBusinessHours.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.Status(http.StatusOK)
}
