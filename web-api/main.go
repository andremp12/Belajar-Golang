package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/database"
	"pustaka-api/handler"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&database.Book{})

	bookRepository := book.NewBookRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/book", bookHandler.GetAll)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.POST("/book", bookHandler.CreateBook)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id")
	/*	v1.GET("/say-hello", handler.SayHello)
		v1.GET("/books/:id", handler.BooksHandlerDetail)
		v1.GET("/books/:id/:title", handler.BooksHandlerDetailTitle)
		v1.GET("/query", handler.QueryHandler)
		v1.POST("/books", handler.PostBooksHandler)*/

	router.Run()
}

//struktur directory
//- handler
//- service
//- repository
//- DB

/**/
