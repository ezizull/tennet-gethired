package assets

import (
	"errors"
	"strings"

	domainErrors "tennet/gethired/domain/errors"
)

func createValidation(assetBody *NewAssetRequest) (err error) {
	var errorsValidation []string

	if assetBody.WalletID == 0 {
		err = domainErrors.NewAppError(errors.New("Wallet id not found"), domainErrors.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainErrors.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainErrors.ValidationError)
	}
	return
}

func updateValidation(assetBody *UpdateAssetRequest) (err error) {
	var errorsValidation []string

	if *assetBody.WalletID == 0 {
		err = domainErrors.NewAppError(errors.New("Wallet id not found"), domainErrors.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainErrors.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainErrors.ValidationError)
	}
	return
}
