package wallet

// Wallet is a struct that contains the wallet model
type Wallet struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
