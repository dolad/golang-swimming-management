package users

import (
	"net/http"
	userdomain "swimming-content-management/domain/userdomain"

	"github.com/gin-gonic/gin"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service userdomain.UserService) {
	usersRouteFactory := func(service userdomain.UserService) {
		// create a new User
		group.POST("/signup", func(c *gin.Context) {
			usersRequestPayload, err := Bind(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
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
