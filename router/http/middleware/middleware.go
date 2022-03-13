package middleware

import (
	"errors"
	"net/http"
	"strings"
	authdomain "swimming-content-management/domain/authdomain"

	"github.com/gin-gonic/gin"
)

func MiddlewareValidAccessToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	token, err := extractToken(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
	}

	userId, errr := authdomain.ValidateAccessToken(token)
	if errr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "invalid token",
		})
	}
	c.Set("UserId", userId)
	c.Next()

}

func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", errors.New("Token not provided or malformed")
	}

	return authHeaderContent[1], nil
}
