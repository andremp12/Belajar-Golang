package middleware

import (
	"authentication-go/initializers"
	"authentication-go/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/* func Auth(c *gin.Context){
	authHeader := c.GetHeader("Authorization")

	// Check the prefix header as a Bearer Token
	if authHeader == "" || !strings.HasPrefix(authHeader,"Bearer "){
		c.JSON(http.StatusUnauthorized,gin.H{
			"messages":"Authorization token not provided",
		})

		c.Abort()
		return
	}

	// Check token in table auth
	var auth models.AuthUser

	tokenStr := strings.TrimPrefix(authHeader,"Bearer ")
	err := initializers.DB.Where("token = ? AND expires_at > ?",tokenStr,time.Now()).First(&auth).Error
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"messages":"Unauthorization",
		})

		c.Abort()
		return
	}

	c.Set("userID",auth.UserID)
	c.Next()
} */

func Auth() gin.HandlerFunc{
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		// Check the prefix header as a Bearer Token
		if authHeader == "" || !strings.HasPrefix(authHeader,"Bearer "){
			c.JSON(http.StatusUnauthorized,gin.H{
				"messages":"Authorization token not provided",
			})

			c.Abort()
			return
		}

		// Check token in table auth
		var auth models.AuthUser

		tokenStr := strings.TrimPrefix(authHeader,"Bearer ")
		err := initializers.DB.Where("token = ? AND expires_at > ?",tokenStr,time.Now()).First(&auth).Error
		if err != nil{
			c.JSON(http.StatusUnauthorized,gin.H{
				"messages":"Unauthorization",
			})

			c.Abort()
			return
		}

		c.Set("userID",auth.UserID)
		c.Next()
	}
}