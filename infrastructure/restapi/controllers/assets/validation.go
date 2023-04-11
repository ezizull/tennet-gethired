package assets

import (
	"errors"
	"strings"

	domainError "tennet/gethired/domain/errors"
)

func createValidation(assetBody *NewAssetRequest) (err error) {
	var errorsValidation []string

	if assetBody.WalletID == 0 {
		err = domainError.NewAppError(errors.New("Wallet id not found"), domainError.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainError.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainError.ValidationError)
	}
	return
}

func updateValidation(assetBody *UpdateAssetRequest) (err error) {
	var errorsValidation []string

	if *assetBody.WalletID == 0 {
		err = domainError.NewAppError(errors.New("Wallet id not found"), domainError.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainError.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainError.ValidationError)
	}
	return
}
