package routes

import (
	userController "tennet/gethired/infrastructure/restapi/controllers/user"
	"tennet/gethired/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerAuth := router.Group("/user")

	// user routes without middleware
	{
		routerAuth.POST("", controller.NewUser)
	}

	// user routes with middleware validation
	routerAuth.Use(middlewares.AuthJWTMiddleware())
	{
		routerAuth.GET("", controller.GetUsersByID)
		routerAuth.PATCH("/:id", controller.UpdateUser)
		routerAuth.DELETE("", controller.DeleteUser)
	}
}
