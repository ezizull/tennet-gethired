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

func (n *UpdateAsset) toDomainMapper() domainAsset.Asset {
	domainAsset := domainAsset.Asset{}

	if n.WalletID != nil {
		domainAsset.WalletID = *n.WalletID
	}

	if n.Name != nil {
		domainAsset.Name = *n.Name
	}

	if n.Symbol != nil {
		domainAsset.Symbol = *n.Symbol
	}

	if n.Network != nil {
		domainAsset.Network = *n.Network
	}

	if n.Address != nil {
		domainAsset.Address = *n.Address
	}

	if n.Balance != nil {
		domainAsset.Balance = *n.Balance
	}

	return domainAsset
}
