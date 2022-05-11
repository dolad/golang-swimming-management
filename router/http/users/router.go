package users

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	userdomain "swimming-content-management/domain/userdomain"

	"github.com/gin-gonic/gin"
)

func NewAuthRoutesFactory(group *gin.RouterGroup) func(service userdomain.UserService) {
	usersRouteFactory := func(service userdomain.UserService) {
		// create a new User
		group.POST("signup", func(c *gin.Context) {
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
			c.JSON(http.StatusOK, ToResponseModel(newUser))
			return
		})

		group.POST("login", func(c *gin.Context) {

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
			c.JSON(http.StatusOK, toAuthResponseModel(authenticateUser))
			return
		})

	}

	return usersRouteFactory
}

func NewUserRoutesFactory(group *gin.RouterGroup) func(service userdomain.UserService) {
	usersRouteFactory := func(service userdomain.UserService) {
		// create a new User

		group.GET("", func(c *gin.Context) {

			usersList, err := service.GetUsers()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, usersList)
			return
		})

		group.GET("profile", func(c *gin.Context) {
			authUserId, _ := c.Get("UserId")
			valStr := fmt.Sprint(authUserId)
			userId, _ := uuid.FromString(valStr)
			userDetails, err := service.GetUserById(userId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, ToResponseModel(userDetails))
			return
		})

		group.PUT("update-profile", func(c *gin.Context) {
			authUserId, _ := c.Get("UserId")
			valStr := fmt.Sprint(authUserId)
			userId, _ := uuid.FromString(valStr)
			updateProfileRequest, err := BindUserProfile(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			userDetails, err := service.UpdateUserProfile(updateProfileRequest, userId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, ToResponseModel(userDetails))
			return
		})

	}

	return usersRouteFactory
}
