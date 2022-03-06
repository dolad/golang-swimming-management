package router

import (
	"net/http"
	errors "swimming-content-management/router/http/errors"
	healthRoutes "swimming-content-management/router/http/health"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler() http.Handler {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(errors.Handler)
	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory()(healthGroup)

	// api := router.Group("/api")
	return router
}
