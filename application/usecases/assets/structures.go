package assets

import (
	domainAsset "tennet/gethired/domain/assets"

	"github.com/shopspring/decimal"
)

// NewAsset is a struct that contains the new asset request information
type NewAsset struct {
	WalletID int64           `json:"wallet_id"`
	Name     string          `json:"name"`
	Symbol   string          `json:"symbol"`
	Network  string          `json:"network"`
	Address  string          `json:"address"`
	Balance  decimal.Decimal `json:"balance" gorm:"type:decimal(16,8)"`
}

// UpdateAsset is a struct that contains the new asset request information
type UpdateAsset struct {
	WalletID *int64           `json:"wallet_id,omitempty" binding:"-"`
	Name     *string          `json:"name,omitempty" binding:"-"`
	Symbol   *string          `json:"symbol,omitempty" binding:"-"`
	Network  *string          `json:"network,omitempty" binding:"-"`
	Address  *string          `json:"address,omitempty" binding:"-"`
	Balance  *decimal.Decimal `json:"balance,omitempty" gorm:"type:decimal(16,8)" binding:"-"`
}

// PaginationResultAsset is a struct that contains the pagination result for asset
type PaginationResultAsset struct {
	Data       *[]domainAsset.Asset
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
