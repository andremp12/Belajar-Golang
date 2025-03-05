package controllers

import (
	"authentication-go/initializers"
	"authentication-go/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// Check the prefix header as a Bearer Token
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "Authorization token not provided",
		})

		return
	}

	// Check the token in auth table
	var auth models.AuthUser

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	err := initializers.DB.Debug().Preload("User").Where("token = ? AND expires_at > ?", tokenStr, time.Now()).First(&auth).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "Unauthorization",
		})

		return
	}

	// Send back response
	c.JSON(http.StatusOK, gin.H{
		"auth":  auth,
		"user":  auth.User,
		"token": tokenStr,
	})
}
