package controllers

import (
	"authentication-go/initializers"
	"authentication-go/models"
	"authentication-go/request"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context){
	// Get email and pass of request body
	var request request.UserRequest

	err := c.ShouldBind(&request)

	if err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Error Input",
		})

		return
	}

	// Look up requested user
	var user models.User

	err = initializers.DB.First(&user,"email = ?",request.Email).Error

	if err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})

		return
	}

	if user.ID == 0{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Email or Password",
		})

		return
	}

	// Compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Email or Password",
		})

		return
	}

	// Generate a JWT token
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "test",
		ID: string(user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// Save token in AuthUser table
	authUser := models.AuthUser{
		UserID: int(user.ID),
		User: user,
		Token: tokenString,
		ExpiresAt: time.Now().Add(24*time.Hour),
	}

	err = initializers.DB.Create(&authUser).Error

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed saved auth session",
		})

		return
	}

	// Send response
	c.JSON(http.StatusOK,gin.H{
		"messages":"Success Login",
		"data": authUser,
	})
}