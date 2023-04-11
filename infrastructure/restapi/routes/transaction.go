package routes

import (
	transactionController "tennet/gethired/infrastructure/restapi/controllers/transaction"
	"tennet/gethired/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// TransactionRoutes is a function that contains all routes of the transaction
func TransactionRoutes(router *gin.RouterGroup, controller *transactionController.Controller) {
	routerTransaction := router.Group("/transaction")
	routerTransaction.Use(middlewares.AuthJWTMiddleware())
	{
		routerTransaction.POST("", controller.NewTranserAsset)
	}
}
