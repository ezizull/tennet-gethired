package assets

import (
	"errors"
	"strings"

	domainErrors "tennet/gethired/domain/errors"
)

func updateValidation(assetBody *UpdateAssetRequest) (err error) {
	var errorsValidation []string

	if *assetBody.WalletID == 0 {
		err = domainErrors.NewAppError(errors.New("Wallet ID NotFound"), domainErrors.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainErrors.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainErrors.ValidationError)
	}
	return
}

func createValidation(assetBody *UpdateAssetRequest) (err error) {
	var errorsValidation []string

	if *assetBody.WalletID == 0 {
		err = domainErrors.NewAppError(errors.New("Wallet ID NotFound"), domainErrors.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainErrors.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainErrors.ValidationError)
	}
	return
}
