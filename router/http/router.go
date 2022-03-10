package router

import (
	"net/http"
	userdomain "swimming-content-management/domain/userdomain"
	errors "swimming-content-management/router/http/errors"
	healthRoutes "swimming-content-management/router/http/health"
	usersRoutes "swimming-content-management/router/http/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(userServices userdomain.UserService) http.Handler {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	// add authorization
	config.AddAllowHeaders("Authorization")

	// add cors config
	router.Use(cors.New(config))

	// handle error hander
	router.Use(errors.Handler)
	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory()(healthGroup)

	// this will map and prefix all endpoint with api
	api := router.Group("/api")

	usersGroup := api.Group("/auth")
	usersRoutes.NewRoutesFactory(usersGroup)(userServices)

	// map routers

	// api := router.Group("/api")
	return router
}
