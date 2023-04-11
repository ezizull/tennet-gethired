// Package  contains the  controller
package book

// NewBookRequest is a struct that contains the new  request information
type NewBookRequest struct {
	Title       string `json:"title" example:"Book" gorm:"unique" binding:"required"`
	Description string `json:"description" example:"Something" binding:"required"`
	Author      string `json:"author" example:"Roche" binding:"required"`
}
