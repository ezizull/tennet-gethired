package assets

import domainAsset "tennet/gethired/domain/assets"

func (book *Asset) toDomainMapper() *domainAsset.Asset {
	return &domainAsset.Asset{
		ID:       book.ID,
		WalletID: book.WalletID,
		Name:     book.Name,
		Symbol:   book.Symbol,
		Network:  book.Network,
		Address:  book.Address,
		Balance:  book.Balance,
	}
}

func fromDomainMapper(book *domainAsset.Asset) *Asset {
	return &Asset{
		ID:       book.ID,
		WalletID: book.WalletID,
		Name:     book.Name,
		Symbol:   book.Symbol,
		Network:  book.Network,
		Address:  book.Address,
		Balance:  book.Balance,
	}
}

func arrayToDomainMapper(books *[]Asset) *[]domainAsset.Asset {
	booksDomain := make([]domainAsset.Asset, len(*books))
	for i, book := range *books {
		booksDomain[i] = *book.toDomainMapper()
	}

	return &booksDomain
}
