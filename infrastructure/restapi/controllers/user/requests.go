// Package user contains the user controller
package user

import (
	_ "github.com/go-playground/validator/v10"
)

// NewUserRequest is a struct that contains the request body for the new user
type NewUserRequest struct {
	UserName  string `json:"username" example:"someUser" gorm:"unique" binding:"required"`
	Email     string `json:"email" example:"mail@mail.com" gorm:"unique" binding:"required"`
	FirstName string `json:"firstName" example:"John" binding:"required"`
	LastName  string `json:"lastName" example:"Doe" binding:"required"`
	Password  string `json:"password" example:"Password123" binding:"required"`
}

// UpdateUserRequest is a struct that contains the request body for the new user
type UpdateUserRequest struct {
	UserName  *string `json:"username,omitempty" example:"someUser" gorm:"unique" binding:"-"`
	Email     *string `json:"email,omitempty" example:"mail@mail.com" gorm:"unique" binding:"-"`
	FirstName *string `json:"firstName,omitempty" example:"John" binding:"-"`
	LastName  *string `json:"lastName,omitempty" example:"Doe" binding:"-"`
	Password  *string `json:"password,omitempty" example:"Password123" binding:"-"`
}
