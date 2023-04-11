package adapter

import (
	userService "tennet/gethired/application/usecases/user"
	userRepository "tennet/gethired/infrastructure/repository/mysql/user"
	userController "tennet/gethired/infrastructure/restapi/controllers/user"

	"gorm.io/gorm"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := userService.Service{UserRepository: uRepository}
	return &userController.Controller{UserService: service}
}
