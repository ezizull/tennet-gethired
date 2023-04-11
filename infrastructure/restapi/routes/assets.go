package routes

import (
	assetsController "tennet/gethired/infrastructure/restapi/controllers/assets"
	"tennet/gethired/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the assets
func AssetRoutes(router *gin.RouterGroup, controller *assetsController.Controller) {
	routerAsset := router.Group("/assets")
	routerAsset.Use(middlewares.AuthJWTMiddleware())
	{
		routerAsset.GET("", controller.GetAllAsset)
		routerAsset.GET("/:id", controller.GetAssetByID)
		routerAsset.POST("", controller.NewAsset)
		routerAsset.PATCH("/:id", controller.UpdateAsset)
		routerAsset.DELETE("/:id", controller.DeleteAsset)
	}
}
