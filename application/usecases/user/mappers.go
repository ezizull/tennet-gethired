// Package user provides the use case for user
package user

import (
	"fmt"
	domainUser "tennet/gethired/domain/user"

	"golang.org/x/crypto/bcrypt"
)

func (n *NewUser) toDomainMapper() *domainUser.User {
	return &domainUser.User{
		UserName:  n.UserName,
		Email:     n.Email,
		FirstName: n.FirstName,
		LastName:  n.LastName,
	}
}

func (n UpdateUser) toDomainMapper() (updateDomain domainUser.User) {
	fmt.Println("check ", n)
	if n.UserName != nil {
		updateDomain.UserName = *n.UserName
	}

	if n.Password != nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(*n.Password), bcrypt.DefaultCost)
		updateDomain.HashPassword = string(hash)
	}

	if n.Email != nil {
		updateDomain.Email = *n.Email
	}

	if n.FirstName != nil {
		updateDomain.FirstName = *n.FirstName
	}

	if n.LastName != nil {
		updateDomain.LastName = *n.LastName
	}

	return updateDomain
}
