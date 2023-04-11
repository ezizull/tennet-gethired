package transaction

import (
	"encoding/json"
	domainError "tennet/gethired/domain/errors"
	domainTransaction "tennet/gethired/domain/transaction"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// Create ... Insert New data
func (r *Repository) Create(newTransaction *domainTransaction.AssetTransaction) (createdTransaction *domainTransaction.AssetTransaction, err error) {
	transaction := fromDomainMapper(newTransaction)

	tx := r.DB.Create(transaction)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError domainError.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = domainError.NewAppErrorWithType(domainError.ResourceAlreadyExists)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
		return
	}

	createdTransaction = transaction.toDomainMapper()
	return
}

// Delete ... Delete transaction
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Unscoped().Delete(&AssetTransaction{}, id)
	if tx.Error != nil {
		err = domainError.NewAppErrorWithType(domainError.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = domainError.NewAppErrorWithType(domainError.NotFound)
	}

	return
}
