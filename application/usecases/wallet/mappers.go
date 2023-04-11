package wallet

import domainWallet "tennet/gethired/domain/wallet"

func (n *NewWallet) toDomainMapper() *domainWallet.Wallet {
	return &domainWallet.Wallet{
		Name: n.Name,
	}
}

func (n *UpdateWallet) toDomainMapper() domainWallet.Wallet {
	domainWallet := domainWallet.Wallet{}

	if n.Name != nil {
		domainWallet.Name = *n.Name
	}

	return domainWallet
}
