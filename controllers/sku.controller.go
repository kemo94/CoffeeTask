package controllers

import (
	"github.com/gin-gonic/gin" 
	"github.com/swnsonhe/task/services" 
	"github.com/swnsonhe/task/requests" 
	"fmt"
)

type SKUController struct {
	service  *services.SKUService  
} 

func NewSKUController() *SKUController {
	controller := SKUController{
		service: services.NewSKUService(),
	}
	return &controller
}

func (controller SKUController) Get(context *gin.Context) {
	
	var listSkusRequest requests.ListSkusRequest
	
	context.Bind(&listSkusRequest)
	fmt.Println("listSkusRequest" , listSkusRequest);
	response := controller.service.Find(listSkusRequest)
	context.JSON(response.GetStatus(), response.GetData())
}

