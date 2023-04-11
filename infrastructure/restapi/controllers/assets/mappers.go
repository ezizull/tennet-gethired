package assets

import (
	"fmt"
	assetUseCase "tennet/gethired/application/usecases/assets"
)

func updateToUsecaseMapper(asset *UpdateAssetRequest) assetUseCase.UpdateAsset {
	updateUsecase := assetUseCase.UpdateAsset{}

	if asset.WalletID != nil {
		updateUsecase.WalletID = asset.WalletID
	}

	if asset.Name != nil {
		updateUsecase.Name = asset.Name
	}

	if asset.Symbol != nil {
		updateUsecase.Symbol = asset.Symbol
	}

	if asset.Network != nil {
		updateUsecase.Network = asset.Network
	}

	if asset.Address != nil {
		updateUsecase.Address = asset.Address
	}

	if asset.Balance != nil {
		updateUsecase.Balance = asset.Balance
	}

	fmt.Println("check ", updateUsecase)

	return updateUsecase

}
