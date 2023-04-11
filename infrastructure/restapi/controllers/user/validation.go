// Package user contains the user controller
package user

import (
	"errors"
	"regexp"
	"strings"
	domainError "tennet/gethired/domain/errors"
)

func updateValidation(request *UpdateUserRequest) (err error) {
	var errorsValidation []string

	// Username must have minimum length of 4
	if request.UserName != nil {
		if len(*request.UserName) < 4 {
			errorsValidation = append(errorsValidation, "Username must be at least 4 characters long")
		}
	}

	// Email must be a valid email format
	if request.Email != nil {
		if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(*request.Email) {
			errorsValidation = append(errorsValidation, "Invalid email format")
		}
	}

	// Password must have minimum length of 8, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if request.Password != nil {
		if len(*request.Password) < 8 {
			errorsValidation = append(errorsValidation, "password must be at least 8 characters long")
		}
		hasSpecialChar := regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString
		if !hasSpecialChar(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one special character")
		}
		hasCapitalLetter := regexp.MustCompile(`[A-Z]+`).MatchString
		if !hasCapitalLetter(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one capital letter")
		}
		hasLowerCase := regexp.MustCompile(`[a-z]+`).MatchString
		if !hasLowerCase(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one lowercase letter")
		}
		hasNumber := regexp.MustCompile(`[0-9]+`).MatchString
		if !hasNumber(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one number")
		}
	}

	// First name must have minimum length of 2
	if request.FirstName != nil {
		if len(*request.FirstName) < 2 {
			errorsValidation = append(errorsValidation, "First name must be at least 2 characters long")
		}
	}

	// Last name must have minimum length of 2
	if request.LastName != nil {
		if len(*request.LastName) < 2 {
			errorsValidation = append(errorsValidation, "Last name must be at least 2 characters long")
		}
	}

	if errorsValidation != nil {
		err = domainError.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainError.ValidationError)
	}
	return
}

func createValidation(request NewUserRequest) error {
	var errorsValidation []string

	// Username must have minimum length of 4
	if len(request.UserName) < 4 {
		errorsValidation = append(errorsValidation, "Username must be at least 4 characters long")
	}

	// Email must be a valid email format
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(request.Email) {
		errorsValidation = append(errorsValidation, "Invalid email format")
	}

	// Password must have minimum length of 8, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if len(request.Password) < 8 {
		errorsValidation = append(errorsValidation, "Password should be at least 8 characters long")
	}
	if !regexp.MustCompile(`[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]`).MatchString(request.Password) {
		errorsValidation = append(errorsValidation, "Password should contain at least one special character")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(request.Password) {
		errorsValidation = append(errorsValidation, "Password should contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(request.Password) {
		errorsValidation = append(errorsValidation, "Password should contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`\d`).MatchString(request.Password) {
		errorsValidation = append(errorsValidation, "Password should contain at least one number")
	}

	// First name must have minimum length of 2
	if len(request.FirstName) < 2 {
		errorsValidation = append(errorsValidation, "First name must be at least 2 characters long")
	}

	// Last name must have minimum length of 2
	if len(request.LastName) < 2 {
		errorsValidation = append(errorsValidation, "Last name must be at least 2 characters long")
	}

	if errorsValidation != nil {
		return domainError.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), domainError.ValidationError)
	}
	return nil
}
