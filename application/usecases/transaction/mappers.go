package transaction

import (
	domainAsset "tennet/gethired/domain/assets"
	domainTransaction "tennet/gethired/domain/transaction"
)

func (n *NewAssetTransaction) toDomainMapper() *domainTransaction.AssetTransaction {
	return &domainTransaction.AssetTransaction{
		SrcAssetID:   n.SrcAssetID,
		DestWalletID: n.DestWalletID,
		SrcWalletID:  n.SrcWalletID,
		Amount:       n.Amount,
		GasFee:       n.GasFee,
		Total:        n.Total,
	}
}

func (n *NewAssetTransaction) toDomainAssetMapper() *domainAsset.Asset {
	return &domainAsset.Asset{
		WalletID: n.DestWalletID,
	}
}
