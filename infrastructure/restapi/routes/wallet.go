package routes

import (
	walletController "tennet/gethired/infrastructure/restapi/controllers/wallet"
	"tennet/gethired/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// WalletRoutes is a function that contains all routes of the wallet
func WalletRoutes(router *gin.RouterGroup, controller *walletController.Controller) {
	routerWallet := router.Group("/wallet")
	routerWallet.Use(middlewares.AuthJWTMiddleware())
	{
		routerWallet.GET("", controller.GetAllWallet)
		routerWallet.GET("/:id", controller.GetWalletByID)
		routerWallet.POST("", controller.NewWallet)
		routerWallet.PATCH("/:id", controller.UpdateWallet)
		routerWallet.DELETE("/:id", controller.DeleteWallet)
	}
}
