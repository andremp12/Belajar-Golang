package book

import (
	"gorm.io/gorm"
	"pustaka-api/database"
)

type Repository interface {
	GetAll() ([]Book, error)
	GetBook(id string) (database.Book, error)
	CreateBook(book database.Book) (database.Book, error)
	UpdateBook(book database.Book) (database.Book, error)
	DeleteBook(id string) error
	//updateBook(id string, request BookRequest) (Book, error)
	//deleteBook(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) GetBook(id string) (database.Book, error) {
	var book database.Book

	err := r.db.First(&book, id).Error

	return book, err
}

func (r *repository) CreateBook(book database.Book) (database.Book, error) {
	err := r.db.Debug().Create(&book).Error

	return book, err
}

func (r *repository) UpdateBook(book database.Book) (database.Book, error) {
	err := r.db.Debug().Save(&book).Error

	return book, err
}

func (r *repository) DeleteBook(id string) error {
	var book database.Book

	err := r.db.Delete(&book{}, id).Error
}
