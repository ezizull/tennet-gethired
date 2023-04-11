package transaction

import (
	domainTransaction "tennet/gethired/domain/transaction"
	assetRepository "tennet/gethired/infrastructure/repository/mysql/assets"
	transactionRepository "tennet/gethired/infrastructure/repository/mysql/transaction"
	walletRepository "tennet/gethired/infrastructure/repository/mysql/wallet"
)

// Service is a struct that contains the repository implementation for transaction use case
type Service struct {
	TransactionRepository transactionRepository.Repository
	AssetRepository       assetRepository.Repository
	WalletRepository      walletRepository.Repository
}

// NewTransferAsset is a function that transfer between asset using wallet id
func (s *Service) NewTransferAsset(newTrans *NewAssetTransaction) (*domainTransaction.AssetTransaction, error) {
	transMap := map[string]interface{}{"id": newTrans.SrcAssetID, "wallet_id": newTrans.SrcWalletID}
	getAsset, err := s.AssetRepository.GetOneByMap(transMap)
	if err != nil {
		return nil, err
	}

	_, err = s.WalletRepository.GetByID((int(newTrans.DestWalletID)))
	if err != nil {
		return nil, err
	}

	domainAsset := newTrans.toDomainAssetMapper(getAsset)
	_, err = s.AssetRepository.Create(domainAsset)
	if err != nil {
		return nil, err
	}

	domainTrans := newTrans.toDomainMapper()
	transaction, err := s.TransactionRepository.Create(domainTrans)
	if err != nil {
		err = s.AssetRepository.Delete(int(newTrans.DestAssetID))
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return transaction, nil
}
