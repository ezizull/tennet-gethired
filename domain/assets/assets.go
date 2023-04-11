package assets

import "github.com/shopspring/decimal"

type Asset struct {
	ID       int64           `json:"id" gorm:"primaryKey"`
	WalletID int64           `json:"wallet_id"`
	Name     string          `json:"name"`
	Symbol   string          `json:"symbol"`
	Network  string          `json:"network"`
	Address  string          `json:"address"`
	Balance  decimal.Decimal `json:"balance" gorm:"type:decimal(16,8)"`
}
