package assets

import domainAsset "tennet/gethired/domain/assets"

func (asset *Asset) toDomainMapper() *domainAsset.Asset {
	return &domainAsset.Asset{
		ID:       asset.ID,
		WalletID: asset.WalletID,
		Name:     asset.Name,
		Symbol:   asset.Symbol,
		Network:  asset.Network,
		Address:  asset.Address,
		Balance:  asset.Balance,
	}
}

func fromDomainMapper(asset *domainAsset.Asset) *Asset {
	return &Asset{
		ID:       asset.ID,
		WalletID: asset.WalletID,
		Name:     asset.Name,
		Symbol:   asset.Symbol,
		Network:  asset.Network,
		Address:  asset.Address,
		Balance:  asset.Balance,
	}
}

func arrayToDomainMapper(assets *[]Asset) *[]domainAsset.Asset {
	assetsDomain := make([]domainAsset.Asset, len(*assets))
	for i, asset := range *assets {
		assetsDomain[i] = *asset.toDomainMapper()
	}

	return &assetsDomain
}
