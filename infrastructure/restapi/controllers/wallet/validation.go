package wallet

import (
	"errors"
	"strings"

	domainError "tennet/gethired/domain/errors"
)

func createValidation(assetBody *NewWalletRequest) (err error) {
	var errorsValidation []string

	if errorsValidation != nil {
		err = domainError.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainError.ValidationError)
	}
	return
}

func updateValidation(assetBody *UpdateWalletRequest) (err error) {
	var errorsValidation []string

	if assetBody.Name == nil {
		err = domainError.NewAppError(errors.New("Wallet name cant empty"), domainError.NotFound)
		return
	}

	if errorsValidation != nil {
		err = domainError.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainError.ValidationError)
	}
	return
}
