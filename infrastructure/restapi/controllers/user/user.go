// Package user contains the user controller
package user

import (
	"errors"
	"net/http"
	"strconv"

	useCaseUser "tennet/gethired/application/usecases/user"
	domainError "tennet/gethired/domain/errors"
	"tennet/gethired/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the user service
type Controller struct {
	UserService useCaseUser.Service
}

// NewUser godoc
// @Tags user
// @Summary Create New UserName
// @Description Create new user on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body NewUserRequest true "body data"
// @Success 200 {object} ResponseUser
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user [post]
func (c *Controller) NewUser(ctx *gin.Context) {
	var request NewUserRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err := createValidation(request)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	userModel, err := c.UserService.Create(toUsecaseMapper(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	userResponse := domainToResponseMapper(userModel)
	ctx.JSON(http.StatusOK, userResponse)
}

// GetUsersByID godoc
// @Tags user
// @Summary Get users by ID
// @Description Get Users by ID on the system
// @Param user_id path int true "id of user"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseUser
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/{user_id} [get]
func (c *Controller) GetUsersByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("user id is invalid"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	user, err := c.UserService.GetByID(userID)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainToResponseMapper(user))
}

// UpdateUser godoc
// @Tags user
// @Summary Get users by ID
// @Description Get Users by ID on the system
// @Param user_id path int true "id of user"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseUser
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/{user_id} [patch]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request UpdateUserRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = updateValidation(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	user, err := c.UserService.Update(userID, updateToUsecaseMapper(&request))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainToResponseMapper(user))
}

// DeleteUser godoc
// @Tags user
// @Summary Get users by ID
// @Description Get Users by ID on the system
// @Param user_id path int true "id of user"
// @Security ApiKeyAuth
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 401 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /user/{user_id} [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.UserService.Delete(userID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})

}
