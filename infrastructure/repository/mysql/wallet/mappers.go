package wallet

import (
	domainWallet "tennet/gethired/domain/wallet"
)

// toDomainMapper function to convert wallet repo to wallet domain
func (wallet *Wallet) toDomainMapper() *domainWallet.Wallet {
	return &domainWallet.Wallet{
		ID:   wallet.ID,
		Name: wallet.Name,
	}
}

// fromDomainMapper function to convert wallet domain to wallet repo
func fromDomainMapper(wallet *domainWallet.Wallet) *Wallet {
	return &Wallet{
		ID:   wallet.ID,
		Name: wallet.Name,
	}
}

// arrayToDomainMapper function to convert list wallet domain to list wallet repo
func arrayToDomainMapper(wallets *[]Wallet) *[]domainWallet.Wallet {
	walletsDomain := make([]domainWallet.Wallet, len(*wallets))
	for i, wallet := range *wallets {
		walletsDomain[i] = *wallet.toDomainMapper()
	}

	return &walletsDomain
}
