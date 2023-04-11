package assets

import (
	"errors"
	"net/http"
	"strconv"
	useCaseAsset "tennet/gethired/application/usecases/assets"
	useCaseAssets "tennet/gethired/application/usecases/assets"
	domainAsset "tennet/gethired/domain/assets"
	domainError "tennet/gethired/domain/errors"
	"tennet/gethired/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	AssetService useCaseAssets.Service
}

// NewAsset godoc
// @Tags asset
// @Summary Create New Asset
// @Descriptioniption Create new asset on the system
// @Accept  json
// @Produce  json
// @Param data body NewAssetRequest true "body data"
// @Success 200 {object} domainAsset.Asset
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset [post]
func (c *Controller) NewAsset(ctx *gin.Context) {
	var request NewAssetRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newAsset := useCaseAsset.NewAsset{
		WalletID: request.WalletID,
		Name:     request.Name,
		Symbol:   request.Symbol,
		Network:  request.Network,
		Address:  request.Address,
		Balance:  request.Balance,
	}

	domainAsset, err := c.AssetService.Create(&newAsset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainAsset)
}

// GetAllAsset godoc
// @Tags asset
// @Summary Get all Assets
// @Description Get all Assets on the system
// @Param   limit  query   string  true        "limit"
// @Param   page  query   string  true        "page"
// @Success 200 {object} []useCaseAsset.PaginationResultAsset
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset [get]
func (c *Controller) GetAllAsset(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param page is necessary to be an integer"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param limit is necessary to be an integer"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	assets, err := c.AssetService.GetAll(page, limit)
	if err != nil {
		appError := domainError.NewAppErrorWithType(domainError.UnknownError)
		_ = ctx.Error(appError)
		return
	}
	ctx.JSON(http.StatusOK, assets)
}

// GetAssetByID godoc
// @Tags asset
// @Summary Get assets by ID
// @Descriptioniption Get Assets by ID on the system
// @Param asset_id path int true "id of asset"
// @Success 200 {object} domainAsset.Asset
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset/{asset_id} [get]
func (c *Controller) GetAssetByID(ctx *gin.Context) {
	assetID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("asset id is invalid"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainAsset, err := c.AssetService.GetByID(assetID)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainAsset)
}

// UpdateAsset godoc
// @Tags asset
// @Summary Get assets by ID
// @Descriptioniption Get Assets by ID on the system
// @Param asset_id path int true "id of asset"
// @Success 200 {object} domainAsset.Asset
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset/{asset_id} [patch]
func (c *Controller) UpdateAsset(ctx *gin.Context) {

	assetID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request UpdateAssetRequest
	err = controllers.BindJSON(ctx, &request)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = updateValidation(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var asset *domainAsset.Asset
	asset, err = c.AssetService.Update(int64(assetID), updateToUsecaseMapper(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, asset)

}

// DeleteAsset godoc
// @Tags asset
// @Summary Get assets by ID
// @Descriptioniption Get Assets by ID on the system
// @Param asset_id path int true "id of asset"
// @Success 200 {object} domainAsset.Asset
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset/{asset_id} [delete]
func (c *Controller) DeleteAsset(ctx *gin.Context) {
	assetID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.AssetService.Delete(assetID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
