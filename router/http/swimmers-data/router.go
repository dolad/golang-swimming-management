package swimmerdata

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"swimming-content-management/domain/swimmingdata"
)

//func NewRoutesFactory() func(group *gin.RouterGroup) {
//	swimmerRouteFactory := func(group *gin.RouterGroup) {
//		group.GET("", func(c *gin.Context) {
//
//		})
//	}
//
//	return swimmerRouteFactory
//}

func NewRoutesFactory(group *gin.RouterGroup) func(service swimmingdata.SwimmingDataService) {
	roleRouteFactory := func(service swimmingdata.SwimmingDataService) {
		group.POST("", func(c *gin.Context) {
			swimmerDataRequest, err := Bind(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			response, err := service.AddSwimmingDataToUser(swimmerDataRequest)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.JSON(http.StatusOK, response)
			return
		})

		group.GET("/:userId", func(c *gin.Context) {
			reqParam := c.Param("userId")
			userId, _ := uuid.FromString(reqParam)
			result, err := service.GetUsersSwimmingData(userId)
			if err != nil {
				c.JSON(http.StatusNotFound, err)
			}
			c.JSON(http.StatusOK, result)
			return
		})

	}
	return roleRouteFactory
}
