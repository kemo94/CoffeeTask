package main

import (
	"fmt"
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/swnsonhe/task/routes"
	"github.com/swnsonhe/task/middlewares"
)
 

func main() { 
	// Creates a gin router with default middleware:
	router := gin.Default()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          300 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
  
	router.Use(middlewares.CORSMiddleware())

	// routes 
	routerGroup := router.Group("/api/v1")
	{
		routes.SKURoutes(routerGroup)
	}

	routes.HealthCheckRoutes(router)

	
	fmt.Println("connected") 

	// Listen and server
	port := ":8080"
	log.Println("Port is " + port)
	router.Run(port)


	
}
