package book

import (
	"pustaka-api/database"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

func BookResponse(b *database.Book) Book {

	return Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Rating:      b.Rating,
		Price:       b.Price,
	}
}

func BookEntity(b *Book) database.Book {
	return database.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Rating:      b.Rating,
		Price:       b.Price,
	}
}
