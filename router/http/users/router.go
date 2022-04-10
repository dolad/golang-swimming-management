package users

import (
	"fmt"
	"net/http"
	userdomain "swimming-content-management/domain/userdomain"

	"github.com/gin-gonic/gin"
)

func NewAuthRoutesFactory(group *gin.RouterGroup) func(service userdomain.UserService) {
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
			fmt.Println(usersRequestPayload)
			newUser, err := service.SignUp(usersRequestPayload)

			if err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, *toResponseModel(newUser))
		})

		group.POST("/login", func(c *gin.Context) {
			loginRequestPayload, err := LoginRequestValidator(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			authenticateUser, err := service.Login(loginRequestPayload)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, *toAuthResponseModel(authenticateUser))
		})

	}

	return usersRouteFactory
}
