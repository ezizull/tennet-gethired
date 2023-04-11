package transaction

import "github.com/shopspring/decimal"

// NewTranserAssetRequest is a struct that contains the new asset request information
type NewTranserAssetRequest struct {
	SrcWalletID  int64           `json:"src_wallet_id"`
	SrcAssetID   int64           `json:"src_asset_id" binding:"required"`
	DestWalletID int64           `json:"dest_wallet_id" binding:"required"`
	Amount       decimal.Decimal `json:"amount" gorm:"type:decimal(16,8)" binding:"required"`
	GasFee       decimal.Decimal `json:"gas_fee" gorm:"type:decimal(16,8)" binding:"required"`
	Total        decimal.Decimal `json:"total" gorm:"type:decimal(16,8)" binding:"required"`
}
