package transaction

import (
	"github.com/shopspring/decimal"
)

// NewAssetTransaction is a struct that contains the new asset request information
type NewAssetTransaction struct {
	SrcWalletID  int64           `json:"src_wallet_id"`
	SrcAssetID   int64           `json:"src_asset_id"`
	DestWalletID int64           `json:"dest_wallet_id"`
	DestAssetID  int64           `json:"dest_asset_id"`
	Amount       decimal.Decimal `json:"amount" gorm:"type:decimal(16,8)"`
	GasFee       decimal.Decimal `json:"gas_fee" gorm:"type:decimal(16,8)"`
	Total        decimal.Decimal `json:"total" gorm:"type:decimal(16,8)"`
}
