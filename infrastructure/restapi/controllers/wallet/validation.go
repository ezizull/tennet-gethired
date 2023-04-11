package wallet

import (
	"errors"
	"strings"

	domainErrors "tennet/gethired/domain/errors"
)

func createValidation(assetBody *NewWalletRequest) (err error) {
	var errorsValidation []string

	if errorsValidation != nil {
		err = domainErrors.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainErrors.ValidationError)
	}
	return
}

func updateValidation(assetBody *UpdateWalletRequest) (err error) {
	var errorsValidation []string

	if assetBody.Name == nil {
		err = domainErrors.NewAppError(errors.New("Wallet name cant empty"), domainErrors.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainErrors.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainErrors.ValidationError)
	}
	return
}
