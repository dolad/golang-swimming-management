package permissions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"swimming-content-management/domain/role"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service role.RoleServices) {
	roleRouteFactory := func(service role.RoleServices) {
		group.GET("", func(c *gin.Context) {
			roleList, err := service.FindAll()
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.JSON(http.StatusOK, roleList)
			return
		})

		group.GET("/:name", func(c *gin.Context) {
			roleName := c.Param("name")
			result, err := service.FindByName(roleName)
			if err != nil {
				c.JSON(http.StatusNotFound, err)
			}
			c.JSON(http.StatusOK, result)
			return
		})

		group.GET("/roleId/:roleId", func(c *gin.Context) {
			reqParam := c.Param("roleId")
			roleId, _ := strconv.ParseUint(reqParam, 10, 32)
			roleIdU32 := uint32(roleId)
			result, err := service.FindById(roleIdU32)
			if err != nil {
				c.JSON(http.StatusNotFound, err)
			}
			c.JSON(http.StatusOK, result)
			return
		})
	}
	return roleRouteFactory
}
