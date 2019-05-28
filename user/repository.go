package user

import "github.com/thenguyenit/go-clean-architecture/models"

//Repository represent the book repository
type Repository interface {
	Fetch() ([]*models.User, error)
	// FindAll() ([]*Book, error)
	// Add(book *Book) (int64, error)
	// Update(book *Book) error
	// Delete(id int64) error
}
