package swimmerdata

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutesFactory() func(group *gin.RouterGroup) {
	swimmerRouteFactory := func(group *gin.RouterGroup) {
		group.GET("records", func(c *gin.Context) {
			authUserId, _ := c.Get("UserId")
			c.JSON(http.StatusOK, authUserId)
		})
	}

	return swimmerRouteFactory
}
