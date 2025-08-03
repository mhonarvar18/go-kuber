package routes

import (
	"auth-service/controllers"
	"auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	// protected group
	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.GET("/me", controllers.Me) // نمونه route محافظت‌شده
}
