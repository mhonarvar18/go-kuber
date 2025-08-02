package controllers

import (
	"auth-service/models"
	"auth-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var users = []models.User{} // در مرحله اول، بدون DB

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	users = append(users, user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var login models.User
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	for _, u := range users {
		if u.Username == login.Username && u.Password == login.Password {
			token, _ := utils.GenerateJWT(u.Username)
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
