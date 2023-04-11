// Package errors contains the error handler controller
package errors

import (
	"net/http"

	domainError "tennet/gethired/domain/errors"

	"github.com/gin-gonic/gin"
)

// MessagesResponse is a struct that contains the response body for the message
type MessagesResponse struct {
	Message string `json:"message"`
}

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*domainError.AppError)
		if ok {
			resp := MessagesResponse{Message: err.Error()}
			switch err.Type {
			case domainError.NotFound:
				c.JSON(http.StatusNotFound, resp)
				return
			case domainError.ValidationError:
				c.JSON(http.StatusBadRequest, resp)
				return
			case domainError.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, resp)
				return
			case domainError.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, resp)
				return
			case domainError.NotAuthorized:
				c.JSON(http.StatusForbidden, resp)
				return
			case domainError.RepositoryError:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			default:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			}
		}

		return
	}
}
