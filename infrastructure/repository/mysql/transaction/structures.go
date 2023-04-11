package transaction

import (
	domainTransaction "tennet/gethired/domain/transaction"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// AssetTransaction is a struct that contains the asset model
type AssetTransaction struct {
	ID           int64           `json:"id" gorm:"primaryKey"`
	SrcWalletID  int64           `json:"src_wallet_id"`
	SrcAssetID   int64           `json:"src_asset_id"`
	DestWalletID int64           `json:"dest_wallet_id"`
	DestAssetID  int64           `json:"dest_asset_id"`
	Amount       decimal.Decimal `json:"amount" gorm:"type:decimal(16,8)"`
	GasFee       decimal.Decimal `json:"gas_fee" gorm:"type:decimal(16,8)"`
	Total        decimal.Decimal `json:"total" gorm:"type:decimal(16,8)"`
	CreatedAt    time.Time       `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time       `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt    *gorm.DeletedAt
}

// TableName overrides the table name used by User to `users`
func (*AssetTransaction) TableName() string {
	return "asset_transactions"
}

// PaginationResultAssetTransaction is a struct that contains the pagination result for asset
type PaginationResultAssetTransaction struct {
	Data       *[]domainTransaction.AssetTransaction
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
