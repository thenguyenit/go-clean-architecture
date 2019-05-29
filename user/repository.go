package user

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
)

//Repository represent the book repository
type Repository interface {
	Fetch(ctx context.Context, cursor string, number int64) ([]*models.User, string, error)
	// FindAll() ([]*Book, error)
	// Add(book *Book) (int64, error)
	// Update(book *Book) error
	// Delete(id int64) error
}
