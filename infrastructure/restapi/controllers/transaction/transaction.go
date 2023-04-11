package transaction

import (
	"net/http"
	useCaseTransaction "tennet/gethired/application/usecases/transaction"
	domainError "tennet/gethired/domain/errors"
	domainTransaction "tennet/gethired/domain/transaction"
	"tennet/gethired/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TransactionService useCaseTransaction.Service
}

// NewTranserAsset godoc
// @Tags asset
// @Summary Create New Asset
// @Descriptioniption Create new asset on the system
// @Accept  json
// @Produce  json
// @Param data body NewTranserAssetRequest true "body data"
// @Success 201 {object} domainTransaction.AssetTransaction
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /asset [post]
func (c *Controller) NewTranserAsset(ctx *gin.Context) {
	var request NewTranserAssetRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	newTransaction := useCaseTransaction.NewAssetTransaction{
		SrcWalletID:  request.SrcWalletID,
		SrcAssetID:   request.SrcAssetID,
		DestWalletID: request.DestWalletID,
		DestAssetID:  request.DestAssetID,
		Amount:       request.Amount,
		GasFee:       request.GasFee,
		Total:        request.Total,
	}

	var (
		domainTransaction *domainTransaction.AssetTransaction
		err               error
	)
	domainTransaction, err = c.TransactionService.NewTransferAsset(&newTransaction)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, domainTransaction)
}
