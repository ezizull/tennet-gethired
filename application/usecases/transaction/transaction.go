package transaction

import (
	"errors"
	domainError "tennet/gethired/domain/errors"
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
	assetMap := map[string]interface{}{"id": newTrans.SrcAssetID, "wallet_id": newTrans.SrcWalletID}
	_, err := s.AssetRepository.GetOneByMap(assetMap)
	if err != nil {
		return nil, domainError.NewAppError(errors.New("Source asset and wallet not found"), domainError.NotFound)
	}

	_, err = s.WalletRepository.GetByID((int(newTrans.DestWalletID)))
	if err != nil {
		err = domainError.NewAppError(errors.New("Destination wallet id not found"), domainError.NotFound)
		return nil, err
	}

	domainAsset := newTrans.toDomainAssetMapper()
	updateAsset, err := s.AssetRepository.Update(newTrans.SrcAssetID, domainAsset)
	if err != nil {
		return nil, err
	}

	domainTrans := newTrans.toDomainMapper()
	domainTrans.DestAssetID = updateAsset.ID

	return s.TransactionRepository.Create(domainTrans)
}
