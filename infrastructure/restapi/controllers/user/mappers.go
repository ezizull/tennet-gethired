// Package user contains the user controller
package user

import (
	userUseCase "tennet/gethired/application/usecases/user"
	userDomain "tennet/gethired/domain/user"
)

func domainToResponseMapper(userDomain *userDomain.User) (createUserResponse *ResponseUser) {
	createUserResponse = &ResponseUser{ID: userDomain.ID, UserName: userDomain.UserName,
		Email: userDomain.Email, FirstName: userDomain.FirstName, LastName: userDomain.LastName,
		CreatedAt: userDomain.CreatedAt, UpdatedAt: userDomain.UpdatedAt}

	return
}

func arrayDomainToResponseMapper(usersDomain *[]userDomain.User) *[]ResponseUser {
	usersResponse := make([]ResponseUser, len(*usersDomain))
	for i, user := range *usersDomain {
		usersResponse[i] = *domainToResponseMapper(&user)
	}
	return &usersResponse
}

func toUsecaseMapper(user *NewUserRequest) userUseCase.NewUser {
	return userUseCase.NewUser{UserName: user.UserName, Password: user.Password, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}

}
