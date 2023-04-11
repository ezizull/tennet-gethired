package wallet

import (
	domainWallet "tennet/gethired/domain/wallet"
)

// Wallet is a struct that contains the wallet model
type Wallet struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// NewWalletRequest is a struct that contains the new wallet request information
type NewWallet struct {
	Name string `json:"name"`
}

// UpdateWallet is a struct that contains the new wallet request information
type UpdateWallet struct {
	Name *string `json:"name,omitempty" binding:"-"`
}

// PaginationResultWallet is a struct that contains the pagination result for wallet
type PaginationResultWallet struct {
	Data       *[]domainWallet.Wallet
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
