package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	permissionDomain "swimming-content-management/domain/permission"
	roleDomain "swimming-content-management/domain/role"
	userdomain "swimming-content-management/domain/userdomain"
	errors "swimming-content-management/router/http/errors"
	healthRoutes "swimming-content-management/router/http/health"
	"swimming-content-management/router/http/middleware"
	permissionRoutes "swimming-content-management/router/http/permissions"
	roleRoutes "swimming-content-management/router/http/role"
	swimmerRoutes "swimming-content-management/router/http/swimmers-data"
	usersRoutes "swimming-content-management/router/http/users"
)

func NewHTTPHandler(userServices userdomain.UserService, permissionService permissionDomain.PermissionService, roleServices roleDomain.RoleServices) http.Handler {
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
	// authRoutes
	authGroup := api.Group("/auth")
	usersRoutes.NewAuthRoutesFactory(authGroup)(userServices)

	swimmerGroup := router.Group("/api/swimmers")
	swimmerGroup.Use(middleware.MiddlewareValidAccessToken)
	swimmerRoutes.NewRoutesFactory()(swimmerGroup)

	permissionGroup := router.Group("/api/permissions")
	permissionGroup.Use(middleware.MiddlewareValidAccessToken)
	permissionRoutes.NewRoutesFactory(permissionGroup)(permissionService)

	roleGroup := router.Group("/api/roles")
	roleGroup.Use(middleware.MiddlewareValidAccessToken)
	roleRoutes.NewRoutesFactory(roleGroup)(roleServices)

	// authGroup.Use(middleware.MiddlewareValidAccessToken)

	// map routers

	// api := router.Group("/api")
	return router
}
