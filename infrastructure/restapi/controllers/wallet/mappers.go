package wallet

import (
	"fmt"
	walletUseCase "tennet/gethired/application/usecases/wallet"
)

func updateToUsecaseMapper(wallet *UpdateWalletRequest) walletUseCase.UpdateWallet {
	updateUsecase := walletUseCase.UpdateWallet{}

	if wallet.Name != nil {
		updateUsecase.Name = wallet.Name
	}

	fmt.Println("check ", updateUsecase)

	return updateUsecase

}
