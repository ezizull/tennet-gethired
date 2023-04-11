package transaction

import domainTransaction "tennet/gethired/domain/transaction"

func (transaction *AssetTransaction) toDomainMapper() *domainTransaction.AssetTransaction {
	return &domainTransaction.AssetTransaction{
		ID:           transaction.ID,
		SrcWalletID:  transaction.SrcWalletID,
		SrcAssetID:   transaction.SrcAssetID,
		DestWalletID: transaction.DestWalletID,
		DestAssetID:  transaction.DestAssetID,
		Amount:       transaction.Amount,
		GasFee:       transaction.GasFee,
		Total:        transaction.Total,
	}
}

func fromDomainMapper(transaction *domainTransaction.AssetTransaction) *AssetTransaction {
	return &AssetTransaction{
		ID:           transaction.ID,
		SrcWalletID:  transaction.SrcWalletID,
		SrcAssetID:   transaction.SrcAssetID,
		DestWalletID: transaction.DestWalletID,
		DestAssetID:  transaction.DestAssetID,
		Amount:       transaction.Amount,
		GasFee:       transaction.GasFee,
		Total:        transaction.Total,
	}
}

func arrayToDomainMapper(assets *[]AssetTransaction) *[]domainTransaction.AssetTransaction {
	assetsDomain := make([]domainTransaction.AssetTransaction, len(*assets))
	for i, transaction := range *assets {
		assetsDomain[i] = *transaction.toDomainMapper()
	}

	return &assetsDomain
}
