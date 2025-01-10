package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pustaka-api/book"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service: service}
}

func (h *bookHandler) GetAll(c *gin.Context) {
	books, err := h.service.GetAll()

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := h.service.GetBook(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var request book.BookRequest

	err := c.Bind(&request)

	if err != nil {
		return
	}

	book, err := h.service.CreateBook(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var request book.BookRequest

	id := c.Param("id")
	err := c.Bind(&request)

	book, err := h.service.UpdateBook(request, id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

/*func Bio(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Andre Mahesa Putra",
		"bio":  "Web Developer",
	})
}

func GetAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Hello World",
		"author":  "Andre Mahesa Putra",
	})
}

func BooksHandlerDetail(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	page := c.Query("page")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"page":  page,
	})
}

func BooksHandlerDetailTitle(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBooksHandler(c *gin.Context) {
	var request request.BookRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on field : %s, condition: %s", e.Field(), e.ActualTag())
			errMessages = append(errMessages, errMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": request.Title,
		"price": request.Price,
	})
}*/
