package users

import (
	"net/http"
	domainErrors "swimming-content-management/domain"
	userdomain "swimming-content-management/domain/userdomain"

	"github.com/gin-gonic/gin"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service userdomain.UserService) {
	usersRouteFactory := func(service userdomain.UserService) {
		// create a new User
		group.POST("/signup", func(c *gin.Context) {
			usersRequestPayload, err := Bind(c)
			if err != nil {
				appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
				c.Error(appError)
				return
			}
			newUser, err := service.SignUp(usersRequestPayload)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, *toResponseModel(newUser))
		})
	}

	return usersRouteFactory
}
