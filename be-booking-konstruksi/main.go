package main

import (
	"booking-konstruksi/controller"
	"booking-konstruksi/database"
	"booking-konstruksi/repository"
	"booking-konstruksi/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_booking_konstruksi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	repoTipeKonstruksi := repository.NewRepoTipeKonstruksi(db)
	serviceTipeKonstruksi := service.NewServiceTipeKonstruksi(repoTipeKonstruksi)
	tipeKonstruksiController := controller.NewTipeKonstruksiController(serviceTipeKonstruksi)

	repoSatuan := repository.NewRepoSatuan(db)
	serviceSatuan := service.NewServiceSatuan(repoSatuan)
	satuanController := controller.NewSatuanController(serviceSatuan)

	router := gin.Default()
	v1 := router.Group("/api/v1")

	v1.GET("/tipe-konstruksi", tipeKonstruksiController.GetAllData)
	v1.GET("/tipe-konstruksi/:id", tipeKonstruksiController.GetData)
	v1.POST("/tipe-konstruksi", tipeKonstruksiController.CreateData)
	v1.PUT("/tipe-konstruksi/:id", tipeKonstruksiController.UpdateData)
	v1.DELETE("/tipe-konstruksi/:id", tipeKonstruksiController.Delete)

	v1.GET("/satuan", satuanController.GetAllData)
	v1.GET("/satuan/:id", satuanController.GetData)
	v1.POST("/satuan", satuanController.CreateData)
	v1.PUT("/satuan/:id", satuanController.UpdateData)
	v1.DELETE("/satuan/:id", satuanController.Delete)
	router.Static("/images", "./images")

	router.Run(":8080")
}
