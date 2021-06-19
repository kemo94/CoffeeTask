package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckRoutes(route *gin.Engine) {
	healthCheck := route.Group("/health-check")
	healthCheck.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Working!")
	})
}