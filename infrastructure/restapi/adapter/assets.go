package adapter

import (
	assetService "tennet/gethired/application/usecases/assets"
	assetRepository "tennet/gethired/infrastructure/repository/mysql/assets"
	walletRepository "tennet/gethired/infrastructure/repository/mysql/wallet"
	assetController "tennet/gethired/infrastructure/restapi/controllers/assets"

	"gorm.io/gorm"
)

// AssetAdapter is a function that returns a asset controller
func AssetAdapter(db *gorm.DB) *assetController.Controller {
	assetRepository := assetRepository.Repository{DB: db}
	walletRepository := walletRepository.Repository{DB: db}

	service := assetService.Service{AssetRepository: assetRepository, WalletRepository: walletRepository}
	return &assetController.Controller{AssetService: service}
}
