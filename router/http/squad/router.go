package squad

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"swimming-content-management/domain/squad"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service squad.SquadDataService) {
	squadRouteFactory := func(service squad.SquadDataService) {
		group.POST("", func(c *gin.Context) {
			swimmerDataRequest, err := Bind(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			response, err := service.CreateSquad(swimmerDataRequest)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			c.JSON(http.StatusOK, response)
			return
		})

		group.GET("/", func(c *gin.Context) {
			result, err := service.GetSquads()
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusOK, result)
			return
		})

		group.GET("/:squadId", func(c *gin.Context) {
			reqParam := c.Param("squadId")
			u, err := strconv.ParseUint(reqParam, 0, 32)
			squadId := uint32(u)
			result, err := service.GetSquad(squadId)
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusOK, result)
			return
		})

		group.POST("/add-coach/:squadId", func(c *gin.Context) {
			reqParam := c.Param("squadId")
			coachId, err := BindAddCoach(c)

			userId, _ := uuid.FromString(coachId)
			u, err := strconv.ParseUint(reqParam, 0, 32)
			squadId := uint32(u)
			result, err := service.AddCoachToSquad(squadId, userId)
			fmt.Println(err)
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusOK, result)
			return
		})

		group.POST("/add-swimmer/:squadId", func(c *gin.Context) {
			reqParam := c.Param("squadId")
			swimmerId, err := BindAddSwimmer(c)
			userId, _ := uuid.FromString(swimmerId)
			u, err := strconv.ParseUint(reqParam, 0, 32)
			squadId := uint32(u)
			result, err := service.AddSwimmerToSquad(squadId, userId)
			fmt.Println(err)
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusOK, result)
			return
		})

	}
	return squadRouteFactory
}
