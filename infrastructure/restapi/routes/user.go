package routes

import (
	userController "tennet/gethired/infrastructure/restapi/controllers/user"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerAuth := router.Group("/user")

	// user routes without middleware
	{
		routerAuth.POST("", controller.NewUser)
	}

}
