package main

import (
	"auth-service/config"
	"auth-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	config.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
