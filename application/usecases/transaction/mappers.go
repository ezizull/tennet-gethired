package transaction

import (
	domainAsset "tennet/gethired/domain/assets"
	domainTransaction "tennet/gethired/domain/transaction"
)

func (n *NewAssetTransaction) toDomainMapper() *domainTransaction.AssetTransaction {
	return &domainTransaction.AssetTransaction{
		SrcWalletID:  n.SrcWalletID,
		SrcAssetID:   n.SrcAssetID,
		DestWalletID: n.DestWalletID,
		DestAssetID:  n.DestAssetID,
		Amount:       n.Amount,
		GasFee:       n.GasFee,
		Total:        n.Total,
	}
}

func (n *NewAssetTransaction) toDomainAssetMapper(asset *domainAsset.Asset) *domainAsset.Asset {
	return &domainAsset.Asset{
		WalletID: n.DestWalletID,
		Name:     asset.Name,
		Symbol:   asset.Symbol,
		Network:  asset.Network,
		Address:  asset.Address,
		Balance:  asset.Balance,
	}
}
