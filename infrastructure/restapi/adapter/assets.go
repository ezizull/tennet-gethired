package adapter

import (
	assetService "tennet/gethired/application/usecases/assets"
	assetRepository "tennet/gethired/infrastructure/repository/mysql/assets"
	assetController "tennet/gethired/infrastructure/restapi/controllers/assets"

	"gorm.io/gorm"
)

// AssetAdapter is a function that returns a asset controller
func AssetAdapter(db *gorm.DB) *assetController.Controller {
	uRepository := assetRepository.Repository{DB: db}
	service := assetService.Service{AssetRepository: uRepository}
	return &assetController.Controller{AssetService: service}
}
