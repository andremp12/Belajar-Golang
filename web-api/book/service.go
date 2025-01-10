package book

import "pustaka-api/database"

type Service interface {
	GetAll() ([]Book, error)
	GetBook(id string) (Book, error)
	CreateBook(request BookRequest) (Book, error)
	UpdateBook(request BookRequest, id string) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Book, error) {

	books, err := s.repository.GetAll()

	return books, err
}

func (s *service) GetBook(id string) (Book, error) {
	book, err := s.repository.GetBook(id)
	bookResponse := BookResponse(&book)

	return bookResponse, err
}

func (s *service) CreateBook(request BookRequest) (Book, error) {
	price, _ := request.Price.Int64()
	rating, _ := request.Rating.Int64()

	b := database.Book{
		Title:       request.Title,
		Price:       int(price),
		Rating:      int(rating),
		Description: request.Description,
	}

	//book.Title = request.Title
	//book.Price = request.Price
	//book.Rating = request.Rating
	//book.Description = request.Description

	book, err := s.repository.CreateBook(b)
	bookResponse := BookResponse(&book)

	return bookResponse, err
}

func (s *service) UpdateBook(request BookRequest, id string) (Book, error) {
	price, _ := request.Price.Int64()
	rating, _ := request.Rating.Int64()

	b, err := s.repository.GetBook(id)
	b.Title = request.Title
	b.Price = int(price)
	b.Rating = int(rating)
	b.Description = request.Description

	book, err := s.repository.UpdateBook(b)
	bookResponse := BookResponse(&book)

	return bookResponse, err
}
