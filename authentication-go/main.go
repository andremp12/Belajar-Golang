package main

import (
	"authentication-go/controllers"
	"authentication-go/initializers"
	"authentication-go/middleware"
	"authentication-go/migrations"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrations.AutoMigrate()
}

func main(){
	r := gin.Default()

	r.Use(gin.Recovery())

	// Create route without middleware
	r.POST("/register",controllers.Register)

	
	// Create route group and Use the middleware in route group 
	authorized := r.Group("/")

	authorized.Use(middleware.Auth())
	{
		authorized.POST("/login",controllers.Login)
		authorized.POST("/validate-token",controllers.ValidateToken)
	}

	r.Run()
}