package wallet

import (
	"errors"
	"net/http"
	"strconv"
	useCaseWallet "tennet/gethired/application/usecases/wallet"
	useCaseWallets "tennet/gethired/application/usecases/wallet"
	domainError "tennet/gethired/domain/errors"
	domainWallet "tennet/gethired/domain/wallet"
	"tennet/gethired/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	WalletService useCaseWallets.Service
}

// NewWallet godoc
// @Tags asset
// @Summary Create New Wallet
// @Descriptioniption Create new asset on the system
// @Accept  json
// @Produce  json
// @Param data body NewWalletRequest true "body data"
// @Success 200 {object} domainWallet.Wallet
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset [post]
func (c *Controller) NewWallet(ctx *gin.Context) {
	var request NewWalletRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newWallet := useCaseWallet.NewWallet{
		Name: request.Name,
	}

	domainWallet, err := c.WalletService.Create(&newWallet)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainWallet)
}

// GetAllWallet godoc
// @Tags asset
// @Summary Get all Wallets
// @Description Get all Wallets on the system
// @Param   limit  query   string  true        "limit"
// @Param   page  query   string  true        "page"
// @Success 200 {object} []useCaseWallet.PaginationResultWallet
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset [get]
func (c *Controller) GetAllWallet(ctx *gin.Context) {
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

	wallet, err := c.WalletService.GetAll(page, limit)
	if err != nil {
		appError := domainError.NewAppErrorWithType(domainError.UnknownError)
		_ = ctx.Error(appError)
		return
	}
	ctx.JSON(http.StatusOK, wallet)
}

// GetWalletByID godoc
// @Tags asset
// @Summary Get wallet by ID
// @Descriptioniption Get Wallets by ID on the system
// @Param asset_id path int true "id of asset"
// @Success 200 {object} domainWallet.Wallet
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset/{asset_id} [get]
func (c *Controller) GetWalletByID(ctx *gin.Context) {
	assetID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("asset id is invalid"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainWallet, err := c.WalletService.GetByID(assetID)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainWallet)
}

// UpdateWallet godoc
// @Tags asset
// @Summary Get wallet by ID
// @Descriptioniption Get Wallets by ID on the system
// @Param asset_id path int true "id of asset"
// @Success 200 {object} domainWallet.Wallet
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset/{asset_id} [patch]
func (c *Controller) UpdateWallet(ctx *gin.Context) {

	assetID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request UpdateWalletRequest
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

	var asset *domainWallet.Wallet
	asset, err = c.WalletService.Update(int64(assetID), updateToUsecaseMapper(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, asset)

}

// DeleteWallet godoc
// @Tags asset
// @Summary Get wallet by ID
// @Descriptioniption Get Wallets by ID on the system
// @Param asset_id path int true "id of asset"
// @Success 200 {object} domainWallet.Wallet
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset/{asset_id} [delete]
func (c *Controller) DeleteWallet(ctx *gin.Context) {
	assetID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.WalletService.Delete(assetID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
