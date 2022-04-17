package permissions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swimming-content-management/domain/permission"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service permission.PermissionService) {
	permissionRouteFactory := func(service permission.PermissionService) {
		group.GET("", func(c *gin.Context) {
			permissionList, err := service.FindAll()
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.JSON(http.StatusOK, permissionList)
			return
		})
	}
	return permissionRouteFactory
}
