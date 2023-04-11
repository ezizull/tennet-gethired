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

func updateToUsecaseMapper(user *UpdateUserRequest) userUseCase.UpdateUser {
	updateUsecase := userUseCase.UpdateUser{}

	if user.UserName != nil {
		updateUsecase.UserName = user.UserName
	}

	if user.Password != nil {
		updateUsecase.Password = user.Password
	}

	if user.Email != nil {
		updateUsecase.Email = user.Email
	}

	if user.FirstName != nil {
		updateUsecase.FirstName = user.FirstName
	}

	if user.LastName != nil {
		updateUsecase.LastName = user.LastName
	}

	return updateUsecase

}
