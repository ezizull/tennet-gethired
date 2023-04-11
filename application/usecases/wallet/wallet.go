package wallet

import (
	walletDomain "tennet/gethired/domain/wallet"
	walletRepository "tennet/gethired/infrastructure/repository/mysql/wallet"
)

// Service is a struct that contains the repository implementation for wallet use case
type Service struct {
	WalletRepository walletRepository.Repository
}

// GetAll is a function that returns all wallet
func (s *Service) GetAll(page int64, limit int64) (*PaginationResultWallet, error) {

	all, err := s.WalletRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &PaginationResultWallet{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a wallet by id
func (s *Service) GetByID(id int) (*walletDomain.Wallet, error) {
	return s.WalletRepository.GetByID(id)
}

// Create is a function that creates a wallet
func (s *Service) Create(wallet *NewWallet) (*walletDomain.Wallet, error) {
	walletModel := wallet.toDomainMapper()
	return s.WalletRepository.Create(walletModel)
}

// GetByMap is a function that returns a wallet by map
func (s *Service) GetByMap(walletMap map[string]interface{}) (*walletDomain.Wallet, error) {
	return s.WalletRepository.GetOneByMap(walletMap)
}

// Delete is a function that deletes a wallet by id
func (s *Service) Delete(id int) error {
	return s.WalletRepository.Delete(id)
}

// Update is a function that updates a wallet by id
func (s *Service) Update(id int64, wallet UpdateWallet) (*walletDomain.Wallet, error) {
	walletModel := wallet.toDomainMapper()
	return s.WalletRepository.Update(id, &walletModel)
}
