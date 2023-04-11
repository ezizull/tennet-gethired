package wallet

import (
	domainWallet "tennet/gethired/domain/wallet"
	"time"

	"gorm.io/gorm"
)

// Wallet is a struct that contains the wallet model
type Wallet struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *gorm.DeletedAt
}

// TableName overrides the table name used by User to `users`
func (*Wallet) TableName() string {
	return "wallets"
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
