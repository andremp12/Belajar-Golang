package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title"`
	Price       json.Number `json:"price" binding:"required,numeric"`
	Rating      json.Number `json:"rating" binding:"required,numeric"`
	Description string      `json:"description"`
}
