package adapter

import (
	transactionService "tennet/gethired/application/usecases/transaction"
	assetRepository "tennet/gethired/infrastructure/repository/mysql/assets"
	transactionRepository "tennet/gethired/infrastructure/repository/mysql/transaction"
	walletRepository "tennet/gethired/infrastructure/repository/mysql/wallet"
	transactionController "tennet/gethired/infrastructure/restapi/controllers/transaction"

	"gorm.io/gorm"
)

// TransactionAdapter is a function that returns a transaction controller
func TransactionAdapter(db *gorm.DB) *transactionController.Controller {
	transRepository := transactionRepository.Repository{DB: db}
	assetRepository := assetRepository.Repository{DB: db}
	walletRepository := walletRepository.Repository{DB: db}

	service := transactionService.Service{
		TransactionRepository: transRepository,
		AssetRepository:       assetRepository,
		WalletRepository:      walletRepository,
	}

	return &transactionController.Controller{TransactionService: service}
}
