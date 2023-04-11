package assets

import "github.com/shopspring/decimal"

// NewAssetRequest is a struct that contains the new asset request information
type NewAssetRequest struct {
	WalletID int64           `json:"wallet_id" binding:"required"`
	Name     string          `json:"name" binding:"required"`
	Symbol   string          `json:"symbol" binding:"required"`
	Network  string          `json:"network" binding:"required"`
	Address  string          `json:"address" binding:"required"`
	Balance  decimal.Decimal `json:"balance" gorm:"type:decimal(16,8)" binding:"required"`
}

// UpdateAssetRequest is a struct that contains the new asset request information
type UpdateAssetRequest struct {
	WalletID *int64           `json:"wallet_id,omitempty" binding:"-"`
	Name     *string          `json:"name,omitempty" binding:"-"`
	Symbol   *string          `json:"symbol,omitempty" binding:"-"`
	Network  *string          `json:"network,omitempty" binding:"-"`
	Address  *string          `json:"address,omitempty" binding:"-"`
	Balance  *decimal.Decimal `json:"balance,omitempty" gorm:"type:decimal(16,8)" binding:"-"`
}
