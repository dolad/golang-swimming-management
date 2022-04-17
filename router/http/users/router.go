package users

import (
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

			newUser, err := service.SignUp(usersRequestPayload)

			if err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, newUser)
		})

		group.POST("/login", func(c *gin.Context) {

			loginRequestPayload, errors := LoginRequestValidator(c)
			if errors != "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errors,
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
			c.JSON(http.StatusOK, authenticateUser)
			return
		})

	}

	return usersRouteFactory
}
