package adapter

import (
	walletService "tennet/gethired/application/usecases/wallet"
	walletRepository "tennet/gethired/infrastructure/repository/mysql/wallet"
	walletController "tennet/gethired/infrastructure/restapi/controllers/wallet"

	"gorm.io/gorm"
)

// WalletAdapter is a function that returns a wallet controller
func WalletAdapter(db *gorm.DB) *walletController.Controller {
	uRepository := walletRepository.Repository{DB: db}
	service := walletService.Service{WalletRepository: uRepository}
	return &walletController.Controller{WalletService: service}
}
