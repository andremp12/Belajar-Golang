package controllers

import (
	"authentication-go/initializers"
	"authentication-go/models"
	"authentication-go/request"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	// Get email/pass of request
	var request request.UserRequest

	err := c.ShouldBind(&request)

	if err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Error Input",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password),10)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed hash password",
		})

		return
	}

	// Create User
	user := models.User{Email: request.Email,Password: string(hash)}
	err = initializers.DB.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed create user",
		})

		return
	}

	// Response
	c.JSON(http.StatusOK,gin.H{
		"messages":"Success Create user",
		"Data":user,
	})
}