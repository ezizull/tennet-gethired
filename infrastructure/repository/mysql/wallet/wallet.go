package wallet

import (
	"encoding/json"
	domainError "tennet/gethired/domain/errors"
	domainWallet "tennet/gethired/domain/wallet"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all wallet data
func (r *Repository) GetAll(page int64, limit int64) (*PaginationResultWallet, error) {
	var wallets []Wallet
	var total int64

	err := r.DB.Model(&Wallet{}).Count(&total).Error
	if err != nil {
		return &PaginationResultWallet{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&wallets).Error

	if err != nil {
		return &PaginationResultWallet{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &PaginationResultWallet{
		Data:       arrayToDomainMapper(&wallets),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// Create ... Insert New data
func (r *Repository) Create(newWallet *domainWallet.Wallet) (createdWallet *domainWallet.Wallet, err error) {
	wallet := fromDomainMapper(newWallet)

	tx := r.DB.Create(wallet)

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

	createdWallet = wallet.toDomainMapper()
	return
}

// GetByID ... Fetch only one wallet by Id
func (r *Repository) GetByID(id int) (*domainWallet.Wallet, error) {
	var wallet Wallet
	err := r.DB.Where("id = ?", id).First(&wallet).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainError.NewAppErrorWithType(domainError.NotFound)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
		return &domainWallet.Wallet{}, err
	}

	return wallet.toDomainMapper(), nil
}

// GetWithQuery ... Fetch only one wallet by Id
func (r *Repository) GetWithQuery(id int) (*domainWallet.Wallet, error) {
	var wallet Wallet
	err := r.DB.Where("id = ?", id).First(&wallet).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainError.NewAppErrorWithType(domainError.NotFound)
		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
		}
		return &domainWallet.Wallet{}, err
	}

	return wallet.toDomainMapper(), nil
}

// GetOneByMap ... Fetch only one wallet by Map
func (r *Repository) GetOneByMap(mapWallet map[string]interface{}) (*domainWallet.Wallet, error) {
	var wallet Wallet

	err := r.DB.Where(mapWallet).Limit(1).Find(&wallet).Error
	if err != nil {
		err = domainError.NewAppErrorWithType(domainError.UnknownError)
		return nil, err
	}
	return wallet.toDomainMapper(), err
}

// Update ... Update wallet
func (r *Repository) Update(id int64, updateWallet *domainWallet.Wallet) (*domainWallet.Wallet, error) {
	var wallet Wallet

	wallet.ID = id
	err := r.DB.Model(&wallet).
		Updates(updateWallet).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainError.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainWallet.Wallet{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainError.NewAppErrorWithType(domainError.ResourceAlreadyExists)
			return &domainWallet.Wallet{}, err

		default:
			err = domainError.NewAppErrorWithType(domainError.UnknownError)
			return &domainWallet.Wallet{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&wallet).Error

	return wallet.toDomainMapper(), err
}

// Delete ... Delete wallet
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&Wallet{}, id)
	if tx.Error != nil {
		err = domainError.NewAppErrorWithType(domainError.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = domainError.NewAppErrorWithType(domainError.NotFound)
	}

	return
}
