package assets

import (
	domainAsset "tennet/gethired/domain/assets"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Asset is a struct that contains the asset model
type Asset struct {
	ID        int64           `json:"id" gorm:"primaryKey"`
	WalletID  int64           `json:"wallet_id"`
	Name      string          `json:"name"`
	Symbol    string          `json:"symbol"`
	Network   string          `json:"network"`
	Address   string          `json:"address"`
	Balance   decimal.Decimal `json:"balance" gorm:"type:decimal(16,8)"`
	CreatedAt time.Time       `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time       `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *gorm.DeletedAt
}

// TableName overrides the table name used by User to `users`
func (*Asset) TableName() string {
	return "assets"
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
