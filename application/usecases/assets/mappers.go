package assets

import domainAsset "tennet/gethired/domain/assets"

func (n *NewAsset) toDomainMapper() *domainAsset.Asset {
	return &domainAsset.Asset{
		WalletID: n.WalletID,
		Name:     n.Name,
		Symbol:   n.Symbol,
		Network:  n.Network,
		Address:  n.Address,
		Balance:  n.Balance,
	}
}

func (n *UpdateAsset) toDomainMapper() *domainAsset.Asset {
	return &domainAsset.Asset{
		WalletID: *n.WalletID,
		Name:     *n.Name,
		Symbol:   *n.Symbol,
		Network:  *n.Network,
		Address:  *n.Address,
		Balance:  *n.Balance,
	}
}
