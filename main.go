package main

import (
	"auth-service/config"
	"auth-service/models"
	"auth-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	// ساخت جدول‌ها به صورت خودکار
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
