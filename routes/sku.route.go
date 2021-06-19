package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swnsonhe/task/controllers"
)


func SKURoutes(route *gin.RouterGroup) {

	SKUController := controllers.NewSKUController()


	route.GET("/sku-list", SKUController.Get) 
}