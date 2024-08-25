package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/core/store"
	xerrors "github.com/oprimogus/cardapiogo/internal/errors"
)

type StoreController struct {
	validator   *validatorutils.Validator
	storeModule store.StoreModule
}

func NewStoreController(validator *validatorutils.Validator, repository store.Repository) *StoreController {
	return &StoreController{
		validator:   validator,
		storeModule: store.NewStoreModule(repository),
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
	storeInstance, err := c.storeModule.GetByID.Execute(ctx, id)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.JSON(http.StatusOK, storeInstance)
}

// GetStoreByFilter godoc
//
//	@Summary		Any user can view filtered stores.
//	@Description	Any user can view filtered stores.
//	@Tags			Store
//	@Accept			json
//	@Produce		json
//	@Param			range		query		int		false	"Specify max range"
//	@Param			score		query		int		false	"Specify in score"
//	@Param			name		query		string	false	"Specify name like"
//	@Param			city		query		string	false	"Specify city"
//	@Param			latitude	query		string	false	"latitude of address selected"
//	@Param			longitude	query		string	false	"longitude of address selected"
//	@Param			type		query		string	false	"Specify store type"
//	@Success		200			{object}	store.GetStoreByIdOutput
//	@Failure		404			{object}	xerrors.ErrorResponse
//	@Failure		500			{object}	xerrors.ErrorResponse
//	@Failure		502			{object}	xerrors.ErrorResponse
//	@Router			/v1/store [get]
func (c *StoreController) GetStoreByFilter(ctx *gin.Context) {
	var params store.StoreFilter
	params.Name = ctx.Query("name")
	params.City = ctx.Query("city")
	params.Latitude = ctx.Query("latitude")
	params.Longitude = ctx.Query("longitude")
	params.Type = store.ShopType(ctx.Query("type"))

	queryRange := ctx.Query("range")
	if queryRange != "" {
		rangeValue, err := strconv.Atoi(queryRange)
		if err != nil {
			xerror := xerrors.Map(err)
			ctx.JSON(xerror.Status, xerror)
			return
		}
		params.Range = rangeValue
	}

	queryScore := ctx.Query("score")
	if queryScore != "" {
		scoreValue, err := strconv.Atoi(queryScore)
		if err != nil {
			xerror := xerrors.Map(err)
			ctx.JSON(xerror.Status, xerror)
			return
		}
		params.Score = scoreValue
	}

	errValidate := c.validator.Validate(params)
	if errValidate != nil {
		xerror := xerrors.Map(errValidate)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	storeList, err := c.storeModule.GetByFilter.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}
	ctx.JSON(http.StatusOK, storeList)
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
	storeID, err := c.storeModule.Create.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": storeID})
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
	err = c.storeModule.Update.Execute(ctx, params)
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
	err = c.storeModule.AddBusinessHour.Execute(ctx, params)
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
	err = c.storeModule.DeleteBusinessHour.Execute(ctx, params)
	if err != nil {
		xerror := xerrors.Map(err)
		ctx.JSON(xerror.Status, xerror)
		return
	}

	ctx.Status(http.StatusOK)
}
