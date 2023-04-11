package wallet

// NewWalletRequest is a struct that contains the new wallet request information
type NewWalletRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateWalletRequest is a struct that contains the new wallet request information
type UpdateWalletRequest struct {
	Name *string `json:"name,omitempty" binding:"-"`
}
