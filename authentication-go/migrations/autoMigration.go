package migrations

import (
	"authentication-go/initializers"
	"authentication-go/models"
)

func AutoMigrate(){
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.AuthUser{})
}